import Card from "./Card";
import "../styles/deck.css";
import React, { useState, useEffect } from "react";

const top = "t";
const left = "l";
const bot = "b";
const right = "r";

function Deck({ cardsId, direction }) {
  const [deck, setDeck] = useState([...cardsId]);
  let cardMaker = makeCard();

  function playCard(el, id) {
    const tableRect = document.getElementById("table").getBoundingClientRect();
    const selfDeck = document.getElementById(bot + "Deck");

    el.classList.add("playedCard" + direction);
    const rect = el.getBoundingClientRect();
    el.style.left = rect.left + "px";
    el.style.top = rect.y + "px";
    selfDeck.appendChild(el);
    setTimeout(() => {
      el.style.transform = `translate(${tableRect.width / 2 - rect.left}px , ${
        tableRect.height / 2 - rect.top
      }px)`;
    }, 50);
    remove(id);
  }

  function makeCard() {
    let dis = -35;
    return (id) => {
      let newCard;
      dis += 35;

      if (direction === top || direction == bot)
        newCard = (
          <li
            onClick={(el) => {
              playCard(el.target, id);
            }}
            key={id}
            className="deckCard"
            style={{ left: "" + dis + "px" }}
          >
            <Card id={id} />
          </li>
        );
      else
        newCard = (
          <li
            key={id}
            className="deckCard rotate90"
            style={{ top: "" + dis * (3 / 4) + "px" }}
          >
            <Card id={id} />
          </li>
        );

      return newCard;
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

  return (
    <ul className="deck" id={direction + "Deck"}>
      {deck.map((id) => cardMaker(id))}
    </ul>
  );
}

export default Deck;
export { bot, top, right, left };
