export default function Setting() {
  var token = process.env.REACT_APP_SERVER_PUBLIC
  // eslint-disable-next-line
  var enableNotif = subscribeNotificationCheck
  return <>
    <div className="header">
      <h4>Setting</h4>
    </div>
    <div>
      <div className="setting">
        <p>Enable Push Notification</p>
        <button onClick={() => enableNotif(token)} className="js-push-btn">Enable</button>
      </div>
    </div>
  </>
}