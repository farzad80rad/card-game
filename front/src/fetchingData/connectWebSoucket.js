import ReactDOM from "react-dom";
import MessageBobble from "../components/MessageBobble";

function getSocket(userId) {
  let socket = new WebSocket("ws://localhost:8081/hokm/websocketBot");
  let selfWonCards = 0;
  let opponentWonCards = 0;
  socket.onopen = () => {
    console.log("Opened");
    socket.send(JSON.stringify({ id: userId }));
    console.log(userId);
  };
  socket.onclose = () => {
    console.log("closed");
  };

  let messagesBox = document.getElementsByClassName("messagesBox")[0];
  let maxPerviosScrolledHeigt = 0;
  socket.onmessage = (message) => {
    let jsonRes = JSON.parse(message.data);
    switch (jsonRes.type) {
      case "put card":
        console.log("Card" + jsonRes.CardToPut);
        let el = document.getElementById("Card" + jsonRes.CardToPut);
        el.setAttribute("selfdefind_ok", true);
        el.click();
        break;
      case "message":
        console.log(jsonRes.message);
        let newMessageDiv = (
          <MessageBobble
            sender={jsonRes.message.sender}
            messageText={jsonRes.message.message}
          />
        );
        let temp = document.createElement("div");
        ReactDOM.render(newMessageDiv, temp);
        console.log(messagesBox.scrollHeight);
        console.log(messagesBox.scrollTop);

        messagesBox.appendChild(temp);
        if (maxPerviosScrolledHeigt <= messagesBox.scrollTop) {
          messagesBox.scrollTop = messagesBox.scrollHeight;
          maxPerviosScrolledHeigt = messagesBox.scrollTop;
        }

        break;
      case "clean table":
        console.log("clean Table");
        let cards = document.getElementsByClassName("playableCard ");
        for (var i = 0; i < cards.length; i++) {
          if (cards[i].classList.contains("gatterCards")) {
            continue;
          }
          cards[i].classList.remove("is-flipped");
          cards[i].classList.add("gatterCards");
          if (jsonRes.SelfTeamWin === true) {
            cards[i].style.left = "" + (40 * selfWonCards + 10) + "px";
            if (selfWonCards % 2 === 1) {
              cards[i].classList.add("cardGatterdRotated");
              cards[i].style.left = "" + (40 * selfWonCards + 30) + "px";
            }
          } else {
            cards[i].style.left = "auto";
            cards[i].style.right = "" + (40 * opponentWonCards + 10) + "px";
            if (opponentWonCards % 2 === 1) {
              cards[i].classList.add("cardGatterdRotated");
              cards[i].style.right = "" + (40 * opponentWonCards - 10) + "px";
            }
          }
        }
        if (jsonRes.SelfTeamWin === true) {
          selfWonCards += 1;
        } else {
          opponentWonCards += 1;
        }

        break;
      default:
        console.log("somting wrong!");
        break;
    }
  };

  return socket;
}

export { getSocket };
