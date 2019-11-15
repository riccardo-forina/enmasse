import * as React from "react";
import {
  IConnectionHeaderDetailProps,
  ConnectionDetailHeader
} from "src/Components/ConnectionDetail/ConnectionDetailHeader";
import { render, fireEvent } from "@testing-library/react";

describe("Connection Detail Header with all connection details", () => {
  test("it renders ConnectionDetailHeader component with all props details at top of the page", () => {
    const props: IConnectionHeaderDetailProps = {
      hostname: "myapp1",
      containerId: "1.219.2.1.33904",
      protocol: "AMQP",
      product: "QpidJMS",
      version: "0.31.0",
      platform: "0.8.0_152.25.125.b16, Oracle Corporation",
      os: "Mac OS X 10.13.6,x86_64",
      messagesIn: 0,
      messagesOut: 1
    };

    const { getByText } = render(<ConnectionDetailHeader {...props} />);
    getByText(props.hostname);
    getByText(props.containerId);
    getByText(props.protocol);
    const seeMoreNode = getByText("see more details");
    fireEvent.click(seeMoreNode);

    getByText("hide details");
    getByText(props.product);
    getByText(props.version + " SNAPSHOT");
    getByText(props.platform);
    getByText(props.os);
    getByText(props.messagesIn + " Message in");
    getByText(props.messagesOut + " Message out");
  });
});
