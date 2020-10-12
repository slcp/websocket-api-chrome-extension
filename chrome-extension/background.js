chrome.commands.onCommand.addListener(function (command) {
  alert(`${JSON.stringify(command)}`);
});

const endpoint = "wss://mudkgw5pwb.execute-api.eu-west-2.amazonaws.com/dev";
var ws;
try {
  ws = new WebSocket(endpoint);
} catch (e) {
  console.log(`Websocket connection failed with error: ${e}`);
}

ws.onopen = () => {
  console.log("Websocket successfully opened");
};
ws.onmessage = function (event) {
  console.log(`Message recieve: ${JSON.stringify(event)}`);
  const message = JSON.parse(event.data)
  if (message.connectionId) {
    alert(`Your connection ID is:\n${message.connectionId}`);
    return
  }
  alert(`You've recieved a message: ${JSON.parse(event.data).message}`);  
};
ws.onclose = () => {
  console.log("Websocket connection was closed");
  process.exit();
};
