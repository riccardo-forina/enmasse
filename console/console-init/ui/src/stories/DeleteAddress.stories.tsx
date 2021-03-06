/*
 * Copyright 2020, EnMasse authors.
 * License: Apache License 2.0 (see the file LICENSE or http://apache.org/licenses/LICENSE-2.0.html).
 */

import React, { createElement } from "react";
import { MemoryRouter } from "react-router";
import { Button } from "@patternfly/react-core";
import { DeletePrompt } from "../Components/Common/DeletePrompt";
import { text } from "@storybook/addon-knobs";

export default {
  title: "Address"
};

export const deleteAddressPrompt = () => {
  return createElement(() => {
    const [isOpen, setIsOpen] = React.useState(false);
    const handleCancel = () => setIsOpen(!isOpen);
    const handleDelete = () => setIsOpen(!isOpen);
    return (
      <MemoryRouter>
        <Button
          onClick={() => {
            setIsOpen(!isOpen);
          }}
        >
          Open Modal On Delete
        </Button>
        <DeletePrompt
          header={text("Header", "Delete the Address ?")}
          name={text("Name at top of detials", "leo_b")}
          detail={text(
            "Details",
            "There are some description that telling users what would happenafter deleting this address."
          )}
          handleConfirmDelete={handleDelete}
          handleCancelDelete={handleCancel}
        />
      </MemoryRouter>
    );
  });
};
