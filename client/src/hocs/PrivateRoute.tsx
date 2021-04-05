import React, { useContext } from "react";
import { Redirect, Route } from "react-router-dom";
import { useAuthContext } from "src/contexts/Auth";

function PrivateRoute({ component: Component, ...rest }: any) {
  const auth = useAuthContext();

  return (
    <Route
      {...rest}
      render={(props) => {
        console.log(auth.isAuthenticated);
        if (auth.isAuthenticated) {
          return <Component {...props} />;
        }
        return <Redirect to="/login" />;
      }}
    />
  );
}

export default PrivateRoute;
