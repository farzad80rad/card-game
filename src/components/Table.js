import Deck from "./Deck";
import "../styles/table.css";

function Table({ decks }) {
  console.log(decks);
  console.log(decks[0]);
  return (
    <section className="table">
      <Deck cardsId={decks[0]} />
    </section>
  );
}

export default Table;
