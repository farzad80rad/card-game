import axios from "axios";

async function getBeginingData(userName) {
  let res = await axios.post(
    "http://localhost:8080/hokm/playWithBot",
    {
      username: userName,
    },
    {
      headers: {
        "Content-Type": "application/json",
      },
    }
  );
  return res.data;
}

export { getBeginingData };
