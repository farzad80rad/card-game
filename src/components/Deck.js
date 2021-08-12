import Card from "./Card";
import "../styles/deck.css";
import React, { useState, useEffect } from "react";

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

  useEffect(() => {
    var cards = document.querySelectorAll(".selfDeck .card");
    cards.forEach((card) => {
      card.addEventListener("dblclick", () => {
        card.classList.add("is-flipped");
        const index = deck.indexOf(card.getAttribute("cardid"));
        let array = [...deck];
        if (index > -1) {
          array.splice(index, 1);
        }
        setDeck(array);
      });
    });
  }, [deck]);

  return <ul className="deck"> {deck.map((id) => cardMaker(id))}</ul>;
}

export default Deck;
export { horizontal, vertical };
