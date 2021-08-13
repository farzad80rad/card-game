import Deck, { top, right, left, bot } from "./Deck";
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
    <section id="table" className="table">
      <div className="topLine">
        <div className="topDeck">
          <div>top player</div>
          <Deck cardsId={decks[1]} direction={top} />
        </div>
      </div>

      <div className="midLine">
        <div className="leftDeck ">
          <div>left player</div>
          <Deck cardsId={decks[1]} direction={left} />
        </div>
        <div className="rightDeck">
          <div>right player</div>

          <Deck cardsId={decks[1]} direction={right} />
        </div>
      </div>

      <div className="botLine">
        <div className="selfDeck">
          <Deck cardsId={decks[0]} direction={bot} />
          <div>self player</div>
        </div>
      </div>
    </section>
  );
}

export default Table;
