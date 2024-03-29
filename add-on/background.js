/*
On startup, connect to the "ping_pong" app.
*/
let port = browser.runtime.connectNative("ping_pong");
console.log(`port -> `, port);

/*
Listen for messages from the app.
*/
port.onMessage.addListener((response) => {
  console.log("Received: ");
  console.log(response);

});

/*
On a click on the browser action, send the app a message.
*/
browser.browserAction.onClicked.addListener(() => {
  console.log("Sending:  ping");
  port.postMessage("ping");
});
