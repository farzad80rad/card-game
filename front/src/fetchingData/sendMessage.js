import axios from "axios";

async function sendMessage(userId, userName, message) {
  await axios.post(
    "http://localhost:8080/hokm/sendMessage",
    {
      groupId: userId,
      sender: userName,
      message: message,
    },
    {
      headers: {
        "Content-Type": "application/json",
      },
    }
  );
}

export { sendMessage };
