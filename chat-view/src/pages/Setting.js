export default function Setting() {
  // eslint-disable-next-line
  var enableNotif = subscribeNotificationCheck
  return <>
    <div className="header">
      <h4>Setting</h4>
    </div>
    <div>
      <div className="setting">
        <p>Enable Push Notification</p>
        <button onClick={() => enableNotif()} className="js-push-btn">Enable</button>
      </div>
    </div>
  </>
}