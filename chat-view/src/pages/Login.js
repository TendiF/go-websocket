import {
  Link,
} from "react-router-dom";
export default function Login() {
  return <>
    <div className="header">
      <h4>Login</h4>
    </div>
    <div>
      <div className="login">
        <div>
          <p>Phone Number</p>
          <div className="input" style={{alignItems:'center'}} >
            <span style={{marginLeft:'5px'}}>+62</span>
            <input placeholder="Phone Number" style={{height:'30px', marginLeft:'10px'}}/>
          </div>
        </div>
        <div>
          <p>Password</p>
          <div className="input" style={{alignItems:'center'}} >
            <input type="password"  placeholder="Password" style={{height:'30px', marginLeft:'10px'}}/>
          </div>
        </div>
      <button style={{marginTop: '15px'}} className="button-primary">Login</button>
      <Link style={{marginTop:'10px'}} to="register">Don't have account ? <span style={{color:'blue'}}>go to register</span> </Link>
      </div>
    </div>
  </>
}