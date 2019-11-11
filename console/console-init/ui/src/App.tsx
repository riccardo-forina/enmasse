import React, { useMemo } from "react";
import "./App.css";
import "@patternfly/react-core/dist/styles/base.css";
import avatarImg from "./logo.svg";
import brandImg from "./brand_logo.svg";
import { NavToolBar } from "./Components/NavToolBar/NavToolBar";
import { AppLayout, SwitchWith404 } from "use-patternfly";
import { useHistory, BrowserRouter as Router } from "react-router-dom";
import { Avatar, Brand, Text, TextVariants } from "@patternfly/react-core";
import { AppRoutes } from "./AppRoutes";

const avatar = (
  <React.Fragment>
    <Text component={TextVariants.p}>Ramakrishna Pattnaik</Text>
    <Avatar src={avatarImg} alt="avatar" />
  </React.Fragment>
);
const logo = <Brand src={brandImg} alt="Console Logo" />;

const App: React.FC = () => {
  const history = useHistory();
  const logoProps = useMemo(
    () => ({
      onClick: () => history.push("/")
    }),
    [history]
  );

  return (
    <SwitchWith404>
      <AppLayout
        logoProps={logoProps}
        logo={logo}
        avatar={avatar}
        toolbar={<NavToolBar />}
      >
        <AppRoutes />
      </AppLayout>
    </SwitchWith404>
  );
};

export default App;
