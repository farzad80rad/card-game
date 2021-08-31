import "../styles/chatPan.css";
import { useEffect } from "react";
import MessageBobble from "./MessageBobble";

function ChatPan() {
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

  return (
    <div className="chatBody">
      <div className="chatBox">
        <div className="messagesBox">
          <MessageBobble sender="tester" messageText="hiiiii " />
        </div>
      </div>

      <div className="sendBox">
        <textarea id="messageTextArea" placeholder="message to send"></textarea>
        <button type="button"> send</button>
      </div>
    </div>
  );
}
export default ChatPan;
