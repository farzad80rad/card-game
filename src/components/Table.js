import Deck, { horizontal, vertical } from "./Deck";
import "../styles/table.css";
import React, { useEffect } from "react";

function Table({ decks }) {
  useEffect(() => {
    var cards = document.querySelectorAll(".selfDeck .card");
    cards.forEach((card) => {
      card.classList.add("is-flipped");
    });
  }, []);

  return (
    <section className="table">
      <div className="topLine">
        <div className="topDeck">
          <Deck cardsId={decks[0]} direction={horizontal} />
        </div>
      </div>

      <div className="midLine">
        <div className="leftDeck ">
          <Deck cardsId={decks[0]} direction={vertical} />
        </div>
        <div className="rightDeck">
          <Deck cardsId={decks[0]} direction={vertical} />
        </div>
      </div>

      <div></div>
      <div></div>

      <div className="botLine">
        <div className="selfDeck">
          <Deck cardsId={decks[0]} direction={horizontal} />
        </div>
      </div>

      <div></div>
    </section>
  );
}

export default Table;
