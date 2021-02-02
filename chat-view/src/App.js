import './App.css';

import {
  BrowserRouter as Router,
  Switch,
  Route,
} from "react-router-dom";

import Home from './pages/Home'
import Setting from './pages/Setting'


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
      </Switch>
    </Router>
  );
}

export default App;
