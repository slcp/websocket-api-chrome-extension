chrome.commands.onCommand.addListener(function (command) {
  alert(`${JSON.stringify(command)}`);
});

function saveConnectionId(id) {
  chrome.storage.local.set({ lambdaTaskCrxConnectionId: id });
}

function loadConnectionId(callback) {
  chrome.storage.local.get(
    "lambdaTaskCrxConnectionId",
    (items) => callback(items[0].lambdaTaskCrxConnectionId)
  );
}

const endpoint = "wss://mudkgw5pwb.execute-api.eu-west-2.amazonaws.com/dev";
var ws;
try {
  ws = new WebSocket(endpoint);
} catch (e) {
  console.log("helloe");
  console.log(e);
}

ws.onopen = () => {
  ws.send(
    JSON.stringify({
      action: "customaction",
      data: {
        hello: "world",
      },
    })
  );
  console.log("connected");
};
ws.onmessage = function (data) {
  console.log(`From server: ${JSON.stringify(data.data)}`);
  // alert connectionId when received from the server
  alert(`You've recieved a message: ${JSON.stringify(data.data)}`);
};
ws.onclose = () => {
  console.log("disconnected");
  process.exit();
};
