function getSocket(userId) {
  let socket = new WebSocket("ws://localhost:8081/hokm/websocketBot");
  socket.onopen = () => {
    console.log("Opened");
    socket.send(JSON.stringify({ id: userId }));
  };
  socket.onclose = () => {
    console.log("closed");
  };
  socket.onmessage = (message) => {
    console.log(message);
  };
  return socket;
}

export { getSocket };
