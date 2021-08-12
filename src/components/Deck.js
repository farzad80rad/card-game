import Card from "./Card";
import "../styles/deck.css";
import React, { useState } from "react";

const horizontal = "h";
const vertical = "v";

function makeCard(direct) {
  let dis = -35;
  return (id) => {
    dis += 35;
    if (direct === horizontal)
      return (
        <li key={id} className="deckCard" style={{ left: "" + dis + "px" }}>
          <Card id={id} />
        </li>
      );
    else
      return (
        <li
          key={id}
          className="deckCard"
          style={{ top: "" + dis * (3 / 4) + "px" }}
        >
          <Card style={{ transform: `rotate( 90deg) ` }} id={id} />
        </li>
      );
  };
}

function Deck({ cardsId, direction }) {
  const [deck, setDeck] = useState([...cardsId]);

  let cardMaker = makeCard(direction);

  return <ul className="deck"> {deck.map((id) => cardMaker(id))}</ul>;
}

export default Deck;
export { horizontal, vertical };
