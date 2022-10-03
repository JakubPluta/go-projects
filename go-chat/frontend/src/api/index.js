/** @format */

var socket = new WebSocket("ws://localhost:8000/ws");

let connect = (callback) => {
  console.log("connecting");

  socket.onopen = () => {
    console.log("Successfully Connected");
  };

  socket.onmessage = (msg) => {
    console.log(msg);
    callback(msg);
  };

  socket.onclose = (event) => {
    console.log("Socket Closed Connection: ", event);
  };

  socket.onerror = (error) => {
    console.log("Socket Error: ", error);
  };
};

let sendMessage = (msg) => {
  console.log("sending msg: ", msg);
  socket.send(msg);
};

export { connect, sendMessage };
