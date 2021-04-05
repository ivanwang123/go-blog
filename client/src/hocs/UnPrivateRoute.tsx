import React, { useContext } from "react";
import { Redirect, Route } from "react-router-dom";
import { useAuthContext } from "src/contexts/Auth";

function UnPrivateRoute({ component: Component, ...rest }: any) {
  const auth = useAuthContext();

  return (
    <Route
      {...rest}
      render={(props) => {
        if (!auth.isAuthenticated) {
          return <Component {...props} />;
        }
        return <Redirect to="/" />;
      }}
    />
  );
  <div></div>;
}

export default UnPrivateRoute;
