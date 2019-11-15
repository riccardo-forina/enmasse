import * as React from "react";
import { render } from "@testing-library/react";
import {
  IAddressDetailHeaderProps,
  AddressDetailHeader
} from "src/Components/AddressDetail/AddressDetailHeader";

describe("Address Detail Header", () => {
  test("it renders address space headers at top", () => {
    const props: IAddressDetailHeaderProps = {
      name: "newqueue",
      type: "Queue",
      plan: "Small",
      shards: 2,
      onDelete: () => {},
      onEdit: () => {}
    };

    const { getByText } = render(<AddressDetailHeader {...props} />);

    getByText(props.name);
    getByText(props.plan);
    getByText(String(props.shards));
  });
});
