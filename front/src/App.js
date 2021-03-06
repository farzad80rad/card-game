import ChatPan from "./components/ChatPan";
import Table from "./components/Table";
import "./styles/app.css";
import { useEffect, useState } from "react";
import { getBeginingData } from "./fetchingData/begin";
import { getSocket } from "./fetchingData/connectWebSoucket";

function App() {
  const [users, setUser] = useState({
    player: { username: "me", deck: ["h2"], id: 0 },
    bots: [
      { username: "bot1", deck: ["h2"], id: 1 },
      { username: "bot2", deck: ["h2"], id: 2 },
      { username: "bot3", deck: ["h2"], id: 3 },
    ],
  });

  const [loaded, setLoaded] = useState(false);

  useEffect(() => {
    async function fetchData() {
      let group = await getBeginingData("me");
      setUser(group);
      getSocket(group.player.id);

      console.log(document.getElementsByTagName("div"));
      setLoaded(true);
    }
    fetchData();
  }, []);

  let divToShow = <p>loadin</p>;
  if (loaded) {
    divToShow = (
      <Table
        users={[
          { name: users.player.username, deck: users.player.deck },
          { name: users.bots[0].username, deck: users.bots[0].deck },
          { name: users.bots[1].username, deck: users.bots[1].deck },
          { name: users.bots[2].username, deck: users.bots[2].deck },
        ]}
        userId={users.player.id}
      />
    );
  }

  return (
    <div className="appBody">
      <div className="tablePan">{divToShow}</div>
      <div className="infoPan">
        <ChatPan userId={users.player.id} userName={users.player.username} />
      </div>
    </div>
  );
}

export default App;
