import * as React from "react";
import { AddressDetailHeader } from "../Components/AddressDetail/AddressDetailHeader";
import { AddressSpaceNavigation } from "../Components/AddressSpace/AddressSpaceNavigation";
import {
  useA11yRouteChange,
  useDocumentTitle,
  SwitchWith404,
  LazyRoute
} from "use-patternfly";
import { PageSection } from "@patternfly/react-core";
import { Redirect, BrowserRouter } from "react-router-dom";

const getConnectionsList = () => import("./ConnectionsListPage");
const getAddressesList = () => import("./AddressesListPage");

export default function AddressSpaceDetailPage() {
  const [activeNavItem, setActiveNavItem] = React.useState("addresses");
  useA11yRouteChange();
  useDocumentTitle("Address Space Detail");

  const onNavSelect = () => {
    if (activeNavItem === "addresses") setActiveNavItem("connections");
    if (activeNavItem === "connections") setActiveNavItem("addresses");
  };

  return (
    <React.Fragment>
      <h1>Address Space Detail Page</h1>
      <AddressSpaceNavigation
        activeItem={activeNavItem}
        onSelect={onNavSelect}
      ></AddressSpaceNavigation>
      <SwitchWith404>
        <Redirect path="/" to="/address-spaces" exact={true} />
        <LazyRoute
          path="/address-spaces/:id/addresses"
          getComponent={getAddressesList}
          exact={true}
        />
        <LazyRoute
          path="/address-spaces/:id/connections"
          getComponent={getConnectionsList}
          exact={true}
        />
      </SwitchWith404>
    </React.Fragment>
  );
}
