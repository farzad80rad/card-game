import Deck, { top, right, left, bot } from "./Deck";
import "../styles/table.css";
import React, { useEffect } from "react";
import PlayerInfo from "./PlayerInfo";

function Table({ users, userId }) {
  useEffect(() => {
    var cards = document.querySelectorAll(".selfDeck .card");
    cards.forEach((card) => {
      card.classList.add("is-flipped");
    });
  }, []);

  return (
    <section id="table" className="table">
      <div className="topLine">
        <div className="toplinePlayer">
          <PlayerInfo name={users[2].name} />
        </div>
        <div className="topDeck">
          <Deck cardsId={users[2].deck} direction={top} />
        </div>
      </div>

      <div className="midLine">
        <div className="leftDeck ">
          <PlayerInfo name={users[3].name} />
          <Deck cardsId={users[3].deck} direction={left} />
        </div>
        <div className="rightDeckPlayer">
          <div className="rightDeck">
            <Deck cardsId={users[1].deck} direction={right} />
          </div>
          <div className="rightlinePlayer">
            <PlayerInfo name={users[1].name} />
          </div>
        </div>
      </div>

      <div className="botLine">
        <div className="selfDeck">
          <Deck userId={userId} cardsId={users[0].deck} direction={bot} />
        </div>
        <div className="selfDeckPlayer">
          <PlayerInfo name={users[0].name} />
        </div>
      </div>

      <div className="selfTeamWoneCards">Self Team Score:</div>
      <div className="opponentTeamWoneCards">Opponent Team Score:</div>
    </section>
  );
}

export default Table;
