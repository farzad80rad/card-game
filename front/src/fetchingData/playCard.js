import axios from "axios";

async function putCard(userId, cardName) {
  console.log(cardName);
  let res = await axios.post(
    "http://localhost:8080/hokm/putCardBot",
    {
      id: userId,
      cardToPut: cardName,
    },
    {
      headers: {
        "Content-Type": "application/json",
      },
    }
  );
  console.log(res.data.status);
  return res.data.status;
}

export { putCard };
