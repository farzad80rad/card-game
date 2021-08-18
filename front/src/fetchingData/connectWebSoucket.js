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
    console.log(message.data.type);
    let jsonRes = JSON.parse(message.data);
    switch (jsonRes.type) {
      case "put card":
        console.log("Card" + jsonRes.CardToPut);
        let el = document.getElementById("Card" + jsonRes.CardToPut);
        el.setAttribute("selfdefind_ok", true);
        el.click();
        break;
      case "clean table":
        console.log("clean Table");
        let cards = document.getElementsByClassName("playableCard ");
        for (var i = 0; i < cards.length; i++) {
          if (cards[i].classList.contains("gatterCards")) {
            console.log("skip");
            continue;
          }
          cards[i].classList.remove("is-flipped");
          cards[i].classList.add("gatterCards");
        }
        break;
    }

    // PlayCard(group.bots[0].deck[0], group.bots[0].deck);
  };

  return socket;
}

export { getSocket };
