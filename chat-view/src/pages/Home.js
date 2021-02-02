import React from "react";
import {
  BrowserRouter as Router,
  Switch,
  Route,
  Link,
  useParams,
  useRouteMatch
} from "react-router-dom";

// Since routes are regular React components, they
// may be rendered anywhere in the app, including in
// child elements.
//
// This helps when it's time to code-split your app
// into multiple bundles because code-splitting a
// React Router app is the same as code-splitting
// any other React app.

export default function NestingExample() {
  return (
    <Router>
      <div>
        <ul>
          <li>
            <Link to="/">Home</Link>
          </li>
          <li>
            <Link to="/1">Detail Chat</Link>
          </li>
        </ul>

        <hr />

        <Switch>
          <Route exact path="/">
            <ListGroupChat />
          </Route>
          <Route path="/:id">
            <DetailChat />
          </Route>
        </Switch>
      </div>
    </Router>
  );
}

function ListGroupChat() {
  return (
    <div>
      <h2>Home</h2>
    </div>
  );
}

function DetailChat() {
  let { id } = useParams();
  return (
    <div>
      <h2>Detail Chats</h2>
    </div>
  );
}
