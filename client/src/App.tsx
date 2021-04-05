import axios from "axios";
import React from "react";
import { QueryClient, QueryClientProvider } from "react-query";
import { ReactQueryDevtools } from "react-query/devtools";
import { BrowserRouter as Router, Route, Switch } from "react-router-dom";
import AuthProvider from "./contexts/Auth";
import PrivateRoute from "./hocs/PrivateRoute";
import UnPrivateRoute from "./hocs/UnPrivateRoute";
import CreatePostPage from "./pages/CreatePostPage";
import HomePage from "./pages/HomePage";
import LoginPage from "./pages/LoginPage";
import PageNotFound from "./pages/PageNotFound";
import PostPage from "./pages/PostPage";
import RegisterPage from "./pages/RegisterPage";

const queryClient = new QueryClient({
  defaultOptions: {
    queries: {
      refetchOnWindowFocus: false,
      retry: false,
    },
  },
});

axios.defaults.baseURL = "http://localhost:8080/api";
axios.defaults.withCredentials = true;
axios.defaults.headers.common = {
  "Content-Type": "application/json",
};

function App() {
  return (
    <QueryClientProvider client={queryClient}>
      <AuthProvider>
        <Router>
          <Switch>
            <Route exact path="/" component={HomePage} />
            <Route exact path="/post/:id" component={PostPage} />
            <UnPrivateRoute exact path="/login" component={LoginPage} />
            <UnPrivateRoute exact path="/register" component={RegisterPage} />
            <PrivateRoute
              exact
              path="/create/post"
              component={CreatePostPage}
            />
            <Route component={PageNotFound} />
          </Switch>
        </Router>
      </AuthProvider>
      <ReactQueryDevtools initialIsOpen={false} />
    </QueryClientProvider>
  );
}

export default App;
