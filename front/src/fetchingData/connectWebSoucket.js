function getSocket(userId) {
  let socket = new WebSocket("ws://localhost:8081/hokm/websocketBot");
  let selfWonCards = 0;
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
          if (cards[i].classList.contains("gatterCardsSelf")) {
            console.log("skip");
            continue;
          }
          cards[i].classList.remove("is-flipped");
          cards[i].classList.add("gatterCardsSelf");
          cards[i].style.left = "auto";
          cards[i].style.right = "" + (40 * selfWonCards + 10) + "px";
          if (selfWonCards % 2 === 1) {
            cards[i].classList.add("cardGatterdRotated");
            cards[i].style.right = "" + (40 * selfWonCards - 10) + "px";

          }
        }
        selfWonCards += 1;

        break;
      default:
        console.log("somting wrong!");
        break;
    }
  };

  return socket;
}

export { getSocket };
