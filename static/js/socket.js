// Create WebSocket connection.
const socket = new WebSocket(`ws://${document.location.host}/ws`);

socket.onopen = () => {
  console.log("Successfully Connected");
};

socket.onmessage = msg => {
  var flash = [{ cssclass: "alert-info", message: `New Note added!, ${msg.data}` }];
  showFlash(flash);
};

socket.onclose = event => {
  console.log("Socket Closed Connection: ", event);
};

socket.onerror = error => {
  console.log("Socket Error: ", error);
};
