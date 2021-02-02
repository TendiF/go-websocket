import React from "react";
import search from '../assets/search.svg'
import papperclip from '../assets/paperclip.svg'
import leftArrow from '../assets/leftArrow.svg'

import {
  BrowserRouter as Router,
  Switch,
  Route,
  Link,
  useParams,
  useRouteMatch
} from "react-router-dom";

export default function NestingExample() {
  return (
    <Router>
      <div>
        <Switch>
          <Route exact path="/chat/">
            <ListGroupChat />
          </Route>
          <Route path="/chat/:id">
            <DetailChat />
          </Route>
        </Switch>
      </div>
    </Router>
  );
}

function DetailChat() {
  return (<>
    <div className="header-detail">
      <div className="title-detail">
        <Link to="/">
          <img style={{height: '25px'}} src={leftArrow}/>
        </Link>
        <h3>Group Badminton</h3>
      </div>
      <div></div>
    </div>
    <div className="chats">
      {[1,2,3,4,5,6,7,8,9].map(v => {
        return <>
          <div className="chat-right">
            <div className="message">
              kumaha damang
              <div style={{textAlign:'right', fontSize: '11px'}}>20:00</div>
            </div>
          </div>
          <div className="chat-left">
            <div className="message">
              <div style={{fontWeight: 'bold', marginBottom: '3px'}}>Tendi F</div>
              Lorep Ipsum dolor sit amet 
              Lorep Ipsum dolor sit amet 
              Lorep Ipsum dolor sit amet 
              <div style={{fontSize: '11px'}}>20:00</div>
            </div>
          </div>
        </>
      })}
    </div>
    <div>
      <div className="footer">
        <div>
          <img style={{height: '18px'}} src={papperclip}/>
          <textarea placeholder="Type a message"/>
        </div>
      </div>
    </div>
  </>);
}

function ListGroupChat() {
  let { id } = useParams();
  return (<>
    <div className="header">
      <div>
        <img src={search}/>
        <input placeholder="Search"/>
      </div>
    </div>
    <div className="groups">
      {[1,2,3,4,5,6,7,8].map(v => {
        return <Link to={'/setting'}>
        <div className="group">
          <div className="title">
            <h4>Group Badminton</h4>
            <span>20/01/2020</span>
          </div>
          <span>kumaha damang kang sadayana ...</span>
        </div>
      </Link>
      })}
    </div>
  </>);
}
