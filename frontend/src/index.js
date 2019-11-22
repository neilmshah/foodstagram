import React from "react";
import ReactDOM from "react-dom";
import { createBrowserHistory } from "history";
import { Router, Route, Switch, Redirect } from "react-router-dom";

import "assets/scss/material-kit-react.scss?v=1.8.0";

// pages for this product

import ProfilePage from "views/ProfilePage/ProfilePage.js";
import LoginPage from "views/LoginPage/LoginPage.js";
import SignupPage from "views/Signup/SignupPage.js";
import TimelinePage from "views/TimelinePage/TimelinePage.js";
import ImagePostPage from "views/ImagePostPage/ImagePostPage.js";

var hist = createBrowserHistory();
function ProtectedRoute({ component: Component, authed, ...rest }) {
  return (
    <Route
      {...rest}
      render={props =>
        authed === true ? (
          <Component {...props} />
        ) : (
          <Redirect to={{ pathname: "/login" }} />
        )
      }
    />
  );
}

function loggedIn() {
  if (
    !localStorage.getItem("user_id") ||
    localStorage.getItem("user_id") === ""
  ) {
    return false;
  } else return true;
}

ReactDOM.render(
  <Router history={hist}>
    <Switch>
      <Route path="/home" component={TimelinePage} />
      <Route path="/profile" component={ProfilePage} />
      <Route path="/login" component={LoginPage} />
      <ProtectedRoute
        authed={loggedIn()}
        path="/image"
        component={ImagePostPage}
      />
      <Route path="/signup" component={SignupPage} />
      <Route path="/" component={TimelinePage} />
    </Switch>
  </Router>,
  document.getElementById("root")
);
