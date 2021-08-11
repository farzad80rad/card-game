import Deck from "./Deck";
import "../styles/table.css";

function Table({ decks }) {
  console.log(decks);
  console.log(decks[0]);
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
