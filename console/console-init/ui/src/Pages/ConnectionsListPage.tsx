import * as React from "react";
import { Link } from "react-router-dom";
import { PageSection, PageSectionVariants } from "@patternfly/react-core";

export default function ConnectionsListPage() {
  return (
    <PageSection variant={PageSectionVariants.light}>
      {console.log("vconn")}
      <h1>Connections List Page</h1>
      <Link to="/address-spaces/123/connections/123">
        Connection Detail Page
      </Link>
    </PageSection>
  );
}
