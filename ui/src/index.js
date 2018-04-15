console.log("starting!")

var websocket = new WebSocket("ws://localhost:8081/ws");

websocket.onmessage = function (event) {
  console.log(event.data);
  document.getElementById("request").innerHTML = event.data;
}

console.log("started");
