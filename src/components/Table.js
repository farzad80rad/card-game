import Deck from "./Deck";
import "../styles/table.css";
import React, { useEffect } from "react";

function Table({ decks }) {
  useEffect(() => {
    var cards = document.querySelectorAll(".selfDeck .card");
    cards.forEach((card) => {
      card.classList.add("is-flipped");
      card.addEventListener("dblclick", function () {
        card.classList.toggle("is-flipped");
      });
    });
  }, []);

  return (
    <section className="table">
      <div className="selfDeck">
        <Deck cardsId={decks[0]} />
      </div>
      <div className="rightDeck">
        <Deck cardsId={decks[0]} />
      </div>
      <div className="topDeck">
        <Deck cardsId={decks[0]} />
      </div>
      <div className="leftDeck">
        <Deck cardsId={decks[0]} />
      </div>
    </section>
  );
}

export default Table;
