function getSocket(userId) {
  let socket = new WebSocket("ws://localhost:8081/hokm/websocketBot");
  socket.onopen = () => {
    console.log("Opened");
    socket.send(JSON.stringify({ id: userId }));
    console.log(userId);
  };
  socket.onclose = () => {
    console.log("closed");
  };

  let g = true;
  socket.onmessage = (message) => {
    let el = document.getElementById("Cardh6");
    if (g) {
      el.click();
      g = false;
    }
    // PlayCard(group.bots[0].deck[0], group.bots[0].deck);
  };

  return socket;
}

export { getSocket };
