import Card from "./Card";
import "../styles/deck.css";
import React, { useState, useEffect } from "react";

const horizontal = "h";
const vertical = "v";

function Deck({ cardsId, direction }) {
  const [deck, setDeck] = useState([...cardsId]);
  let cardMaker = makeCard();

  function makeCard() {
    let dis = -35;
    return (id) => {
      dis += 35;
      if (direction === horizontal)
        return (
          <li
            onDoubleClick={() => remove(id)}
            key={id}
            className="deckCard"
            style={{ left: "" + dis + "px" }}
          >
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

  function remove(id) {
    const index = deck.indexOf(id);
    let array = [...deck];
    if (index > -1) {
      array.splice(index, 1);
    }
    setDeck(array);
  }

  return <ul className="deck"> {deck.map((id) => cardMaker(id))}</ul>;
}

export default Deck;
export { horizontal, vertical };
