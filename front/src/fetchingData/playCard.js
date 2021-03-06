import axios from "axios";

async function putCard(userId, cardName) {
  console.log("in");
  console.log(cardName);
  console.log(userId);
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
  return res.data.status;
}

export { putCard };
