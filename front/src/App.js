import InfoPanel from "./components/InfoPanel";
import Table from "./components/Table";
import "./styles/app.css";
import { useEffect, useState } from "react";
import { getBeginingData } from "./fetchingData/begin";

function App() {
  const [users, setUser] = useState({
    player: { name: "me", deck: ["h2"], id: 1 },
    bots: [
      { name: "bot1", deck: ["h2"] },
      { name: "bot2", deck: ["h2"] },
      { name: "bot3", deck: ["h2"] },
    ],
  });

  const [loaded, setLoaded] = useState(false);

  useEffect(() => {
    async function fetchData() {
      let group = await getBeginingData("me");
      setUser(group);
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
      />
    );
  }

  return (
    <div className="appBody">
      <div className="tablePan">{divToShow}</div>
      <div className="infoPan">
        <InfoPanel />
      </div>
    </div>
  );
}

export default App;
