import Card from "./Card";
import "../styles/deck.css";
import React, { useState, useEffect } from "react";
import { putCard } from "../fetchingData/playCard";

const top = "t";
const left = "l";
const bot = "b";
const right = "r";

function Deck({ cardsId, direction, userId }) {
  const [deck, setDeck] = useState([...cardsId]);
  let cardMaker = makeCard();
  let userIdin = userId;

  async function playCard(id) {
    console.log("put card event");

    if (direction == bot) await putCard(userIdin, id);
    //if (!(await putCard(userId, id))) {
    //  return;
    //}
    let el = document.getElementById("Card" + id);

    const selfDeck = document.getElementById(direction + "Deck");
    const rect = el.getBoundingClientRect();
    let newCard = el;
    console.log(direction);
    newCard.classList.add("playableCard");
    newCard.style.left = rect.left + "px";
    newCard.style.top = rect.y + "px";
    selfDeck.appendChild(newCard);
    newCard.classList.add("playableCard");

    console.log("thiiiiiiiiiiiiiiiiiiiiiiiis:");
    setTimeout(() => {
      switch (direction) {
        case bot:
          newCard.style.left = "38%";
          newCard.style.top = "55%";
          break;

        case left:
          newCard.style.left = "31%";
          newCard.style.top = "42%";
          newCard.classList.toggle("is-flipped");
          break;

        case top:
          newCard.style.left = "38%";
          newCard.style.top = "30%";
          newCard.classList.toggle("is-flipped");
          break;

        case right:
          newCard.style.left = "45%";
          newCard.style.top = "42%";
          newCard.classList.toggle("is-flipped");
      }
      console.log(newCard);
      remove(id);
    }, 50);
  }

  function remove(id) {
    const index = deck.indexOf(id);
    let array = [...deck];
    if (index > -1) {
      array.splice(index, 1);
    }
    setDeck(array);
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
              playCard(id);
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
            onClick={(el) => {
              playCard(id);
            }}
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

  return (
    <ul
      className={"deck" + (direction === top || direction === bot ? "h" : "v")}
      id={direction + "Deck"}
    >
      {deck.map((id) => cardMaker(id))}
    </ul>
  );
}

export default Deck;
export { bot, top, right, left };
