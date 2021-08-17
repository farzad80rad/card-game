import { PlayCard } from "../components/Deck";

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
    let el = document.getElementById("Cardh6");
    el.click();
    //  let elms = document.getElementsByTagName("div");
    //   console.log(elms);
    PlayCard("h6");
  };
  return socket;
}

export { getSocket };
