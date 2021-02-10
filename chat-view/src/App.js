import './App.css';

import {
  BrowserRouter as Router,
  Switch,
  Route,
} from "react-router-dom";

import Home from './pages/Home'
import Setting from './pages/Setting'
import Login from './pages/Login'
import Register from './pages/Register'
import Counter from './features/counter/Counter'


function App() {
  return (
    <Router>
      <Switch>
        <Route path="/setting">
          <Setting/>
        </Route>
        <Route path="/chat">
          <Home/>
        </Route>
        <Route path="/login">
          <Login/>
        </Route>
        <Route path="/register">
          <Register/>
        </Route>
        <Route path="/">
          <Counter/>
        </Route>
      </Switch>
    </Router>
  );
}

export default App;
