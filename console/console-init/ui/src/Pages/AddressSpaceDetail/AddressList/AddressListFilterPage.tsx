/*
 * Copyright 2020, EnMasse authors.
 * License: Apache License 2.0 (see the file LICENSE or http://apache.org/licenses/LICENSE-2.0.html).
 */

import * as React from "react";

import {
  DataToolbar,
  DataToolbarItem,
  DataToolbarContent
} from "@patternfly/react-core/dist/js/experimental";
import { CreateAddressPage } from "../../CreateAddress/CreateAddressPage";
import { useParams } from "react-router";
import { useApolloClient } from "@apollo/react-hooks";
import { RETURN_ADDRESS_SPACE_DETAIL } from "src/Queries/Queries";
import { IAddressSpacesResponse } from "src/Types/ResponseTypes";
import {
  AddressListFilter,
  AddressListKebab
} from "src/Pages/AddressSpaceDetail/AddressList/AddressListFilter";
import useWindowDimensions from "src/Components/Common/WindowDimension";
import { SortForMobileView } from "src/Components/Common/SortForMobileView";
import { ISortBy } from "@patternfly/react-table";
interface AddressListFilterProps {
  filterValue: string | null;
  setFilterValue: (value: string | null) => void;
  filterNames: string[];
  setFilterNames: (value: Array<string>) => void;
  typeValue: string | null;
  setTypeValue: (value: string | null) => void;
  statusValue: string | null;
  setStatusValue: (value: string | null) => void;
  sortValue?: ISortBy;
  setSortValue: (value: ISortBy) => void;
  totalAddresses: number;
}
export const AddressListFilterPage: React.FunctionComponent<AddressListFilterProps> = ({
  filterValue,
  setFilterValue,
  filterNames,
  setFilterNames,
  typeValue,
  setTypeValue,
  statusValue,
  setStatusValue,
  sortValue,
  setSortValue,
  totalAddresses
}) => {
  const { name, namespace, type } = useParams();
  const [isCreateWizardOpen, setIsCreateWizardOpen] = React.useState(false);
  const [addressSpacePlan, setAddressSpacePlan] = React.useState();
  const client = useApolloClient();
  const { width } = useWindowDimensions();

  const onDeleteAll = () => {
    setFilterValue("Name");
    setTypeValue(null);
    setStatusValue(null);
    setFilterNames([]);
  };
  const sortMenuItems = [
    { key: "name", value: "Name", index: 1 },
    { key: "messageIn", value: "Messages In", index: 3 },
    { key: "messageOut", value: "Messages Out", index: 4 },
    { key: "storedMessage", value: "Stored Messages", index: 5 }
  ];

  const createAddressOnClick = async () => {
    setIsCreateWizardOpen(!isCreateWizardOpen);
    if (name && namespace) {
      const addressSpace = await client.query<IAddressSpacesResponse>({
        query: RETURN_ADDRESS_SPACE_DETAIL(name, namespace)
      });
      if (
        addressSpace.data &&
        addressSpace.data.addressSpaces &&
        addressSpace.data.addressSpaces.AddressSpaces.length > 0
      ) {
        const plan =
          addressSpace.data.addressSpaces.AddressSpaces[0].Spec.Plan.ObjectMeta
            .Name;
        if (plan) {
          setAddressSpacePlan(plan);
        }
      }
    }
  };

  const toolbarItems = (
    <>
      <AddressListFilter
        filterValue={filterValue}
        setFilterValue={setFilterValue}
        filterNames={filterNames}
        setFilterNames={setFilterNames}
        typeValue={typeValue}
        setTypeValue={setTypeValue}
        statusValue={statusValue}
        setStatusValue={setStatusValue}
        totalAddresses={totalAddresses}
        addressspaceName={name}
        namespace={namespace}
      />
      {width < 769 && (
        <SortForMobileView
          sortMenu={sortMenuItems}
          sortValue={sortValue}
          setSortValue={setSortValue}
        />
      )}
      <DataToolbarItem>
        {isCreateWizardOpen && (
          <CreateAddressPage
            name={name || ""}
            namespace={namespace || ""}
            addressSpace={name || ""}
            addressSpacePlan={addressSpacePlan || ""}
            addressSpaceType={type || ""}
            isCreateWizardOpen={isCreateWizardOpen}
            setIsCreateWizardOpen={setIsCreateWizardOpen}
          />
        )}
      </DataToolbarItem>
      <DataToolbarItem>
        <AddressListKebab createAddressOnClick={createAddressOnClick} />
      </DataToolbarItem>
    </>
  );
  return (
    <DataToolbar
      id="data-toolbar-with-filter"
      className="pf-m-toggle-group-container"
      collapseListedFiltersBreakpoint="xl"
      clearAllFilters={onDeleteAll}>
      <DataToolbarContent>{toolbarItems}</DataToolbarContent>
    </DataToolbar>
  );
};
