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

  socket.onmessage = (message) => {
    console.log(message);
    console.log(message.data);
    console.log("Card" + message.data.CardToPut);
    let el = document.getElementById(
      "Card" + JSON.parse(message.data).CardToPut
    );

    el.click();
    // PlayCard(group.bots[0].deck[0], group.bots[0].deck);
  };

  return socket;
}

export { getSocket };
