import im from "../assets/player1.jfif";
import "../styles/messageBobble.css";

function MessageBobble({ sender, messageText }) {
  return (
    <div className="messageBobble">
      <img src={im} alt="user"></img>
      <div>
        <div className="triangle"></div>
        <div className="messageBobble-bobble">
          <p className="messageBobble-sender">{sender}</p>
          <p className="messageBobble-text">{messageText}</p>
        </div>
      </div>
    </div>
  );
}

export default MessageBobble;
