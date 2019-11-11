import * as React from "react";
import { SwitchWith404, LazyRoute } from "use-patternfly";
import { Redirect } from "react-router";

const getAddressSpaceListPage = () => import("./Pages/AddressSpaceListPage");
const getAddressSpaceDetail = () => import("./Pages/AddressSpaceDetailPage");
const getAddressDetail = () => import("./Pages/AddressDetailPage");
const getConnectionDetail = () => import("./Pages/ConnectionDetailPage");

export const AppRoutes = () => (
  <SwitchWith404>
    <Redirect path="/" to="/address-spaces" exact={true} />
    <LazyRoute
      path="/address-spaces"
      exact={true}
      getComponent={getAddressSpaceListPage}
    />
    <LazyRoute
      path="/address-spaces/:id/addresses"
      exact={true}
      getComponent={getAddressSpaceDetail}
    />
    <LazyRoute
      path="/address-spaces/:id/connections"
      exact={true}
      getComponent={getAddressSpaceDetail}
    />
    <LazyRoute
      path="/address-spaces/:id/addresses/:id"
      exact={true}
      getComponent={getAddressDetail}
    />
    <LazyRoute
      path="/address-spaces/:id/connections/:id"
      getComponent={getConnectionDetail}
    />
  </SwitchWith404>
);
