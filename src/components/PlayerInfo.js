import "../styles/playerInfo.css";
import im from "../assets/player1.jfif";

function PlayerInfo({ name }) {
  return (
    <div className="playerInfoDiv">
      <img src={im} className="playerImg" alt="palyer" />
      <p className="playerName">{name}</p>
    </div>
  );
}

export default PlayerInfo;
