import "../styles/chatPan.css";
import { useEffect, useState } from "react";
import MessageBobble from "./MessageBobble";
import { sendMessage } from "../fetchingData/sendMessage";

function ChatPan({ userId, userName }) {
  useEffect(() => {
    var text = document.getElementById("messageTextArea");
    var sendBox = document.getElementsByClassName("sendBox")[0];
    function resize() {
      sendBox.style.height =
        text.scrollHeight + 5 > 120 ? "120px" : text.scrollHeight + 5 + "px";
    }
    /* 0-timeout to get the already changed text */
    function delayedResize() {
      window.setTimeout(resize, 0);
    }
    text.addEventListener("change", resize);
    text.addEventListener("cut", delayedResize);
    text.addEventListener("paste", delayedResize);
    text.addEventListener("drop", delayedResize);
    text.addEventListener("keydown", delayedResize);
    resize();
  }, []);

  const [message, setMessage] = useState("");

  return (
    <div className="chatBody">
      <div className="chatBox">
        <div className="messagesBox">
          <MessageBobble sender="tester" messageText="hiiiii " />
        </div>
      </div>

      <div className="sendBox">
        <textarea
          id="messageTextArea"
          placeholder="message to send"
          value={message}
          onChange={(el) => setMessage(el.target.value)}
        ></textarea>
        <button
          type="button"
          onClick={() => {
            sendMessage(userId, userName, message);
            setMessage("");
            var sendBox = document.getElementsByClassName("sendBox")[0];
            sendBox.style.height = "initial";
          }}
        >
          send
        </button>
      </div>
    </div>
  );
}
export default ChatPan;
