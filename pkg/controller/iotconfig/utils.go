/*
 * Copyright 2019, EnMasse authors.
 * License: Apache License 2.0 (see the file LICENSE or http://apache.org/licenses/LICENSE-2.0.html).
 */

package iotconfig

import (
	"context"

	"github.com/enmasseproject/enmasse/pkg/util/cchange"

	"sigs.k8s.io/controller-runtime/pkg/client"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"

	iotv1alpha1 "github.com/enmasseproject/enmasse/pkg/apis/iot/v1alpha1"
	"github.com/enmasseproject/enmasse/pkg/util/install"
	routev1 "github.com/openshift/api/route/v1"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
)

// This sets the default Hono probes
func SetHonoProbes(container *corev1.Container) {

	container.ReadinessProbe = install.ApplyHttpProbe(container.ReadinessProbe, 10, "/readiness", 8088)
	container.LivenessProbe = install.ApplyHttpProbe(container.LivenessProbe, 10, "/liveness", 8088)

}

func FullHostNameForEnvVar(serviceName string) string {
	return serviceName + ".$(KUBERNETES_NAMESPACE).svc"
}

// Append a string to the value of an env-var. If the env-var doesn't exist, it will be created with the provided value.
// A whitespace is added between the existing value and the new value.
func AppendEnvVarValue(container *corev1.Container, name string, value string) {
	if container.Env == nil {
		container.Env = make([]corev1.EnvVar, 0)
	}

	opts := ""

	for _, env := range container.Env {
		if env.Name == name {
			opts = env.Value
		}
	}

	if len(opts) > 0 {
		opts += " "
	}

	opts += value

	container.Env = append(container.Env, corev1.EnvVar{
		Name:  name,
		Value: opts,
	})
}

func AppendStandardHonoJavaOptions(container *corev1.Container) {
	AppendEnvVarValue(
		container,
		"JAVA_APP_OPTS",
		"-Djava.net.preferIPv4Stack=true -Dvertx.logger-delegate-factory-class-name=io.vertx.core.logging.SLF4JLogDelegateFactory",
	)
}

func applyDefaultDeploymentConfig(deployment *appsv1.Deployment, serviceConfig iotv1alpha1.ServiceConfig, configCtx *cchange.ConfigChangeRecorder) {
	deployment.Spec.Replicas = serviceConfig.Replicas
	deployment.Spec.Strategy.Type = appsv1.RollingUpdateDeploymentStrategyType
	if configCtx != nil {
		deployment.Spec.Template.Annotations["iot.enmasse.io/config-hash"] = configCtx.HashString()
	} else {
		delete(deployment.Spec.Template.Annotations, "iot.enmasse.io/config-hash")
	}
}

func applyContainerConfig(container *corev1.Container, config *iotv1alpha1.ContainerConfig) {

	if config == nil {
		return
	}

	if config.Resources != nil {
		container.Resources = *config.Resources
	}

}

func (r *ReconcileIoTConfig) cleanupSecrets(ctx context.Context, config *iotv1alpha1.IoTConfig, adapterName string) error {

	// we need to use an unstructured list, as "SecretList" doesn't work
	// due to kubernetes-sigs/controller-runtime#362

	ul := unstructured.UnstructuredList{}
	ul.SetKind("SecretList")
	ul.SetAPIVersion("")

	ls, err := install.LabelSelectorFromMap(install.CreateDefaultLabels(nil, "iot", adapterName+"-tls"))
	if err != nil {
		return err
	}

	n, err := install.BulkRemoveOwner(ctx, r.client, config, true, &ul, client.ListOptions{
		Namespace:     config.GetNamespace(),
		LabelSelector: ls,
	})

	if err == nil {
		log.Info("cleaned up adapter secrets", "adapter", adapterName, "secretsDeleted", n)
	}

	return err
}

func deviceRegistryImplementation(config *iotv1alpha1.IoTConfig) DeviceRegistryImplementation {

	var file = config.Spec.ServicesConfig.DeviceRegistry.File
	var infinispan = config.Spec.ServicesConfig.DeviceRegistry.Infinispan

	if infinispan != nil && file == nil {
		return DeviceRegistryInfinispan
	} else if infinispan == nil && file != nil {
		return DeviceRegistryFileBased
	} else {
		return DeviceRegistryIllegal
	}
}

func updateEndpointStatus(protocol string, forcePort bool, service *routev1.Route, status *iotv1alpha1.EndpointStatus) {

	status.URI = ""

	if service.Spec.Host == "" {
		return
	}

	status.URI = protocol + "://" + service.Spec.Host

	if forcePort {
		status.URI += ":443"
	}

}
