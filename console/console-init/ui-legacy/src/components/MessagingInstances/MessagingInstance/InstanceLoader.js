import React from 'react';

import axios from 'axios';

import {loadMessagingInstance, loadMessagingInstances, deleteMessagingInstances} from './Enmasse/EnmasseAddressSpaces';

class InstanceLoader {

  translateNamespaces = namespaces => {
    return namespaces.items.map(namespace => namespace.metadata.name);
  };

  loadNamespaces() {

    if (window.env.OPENSHIFT_AVAILABLE) {
      return axios.get('apis/project.openshift.io/v1/projects')
        .then(response => {
          return this.translateNamespaces(response.data);
        })
        .catch(error => {
            console.log('FAILED to load namespaces', error);
            throw(error);
          }
        );
    } else {
      console.log('Can not load namespaces in Kubernetes');
      return new Promise((resolve, reject) => resolve([]));
    }
  }

  loadInstances() {
    if (window.env.OPENSHIFT_AVAILABLE) {
      return this.loadNamespaces().then(namespaces => {
        let promises = [];
        namespaces.forEach(namespace => {
          promises.push(loadMessagingInstance(namespace));
        });
        return Promise.all(promises)
          .then(values => {
            return [].concat.apply([], values);
          })
          .catch(error => {
            console.log('failed: ', error);
            return error;
          });
      });
    } else {
      return loadMessagingInstances();
    }
  }
};

export default (new InstanceLoader);
