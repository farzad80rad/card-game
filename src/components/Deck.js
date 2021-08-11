import Card from "./Card";
import "../styles/deck.css";

function makeCard() {
  let dis = -35;
  return (id) => {
    dis += 35;
    return (
      <li key={id} className="deckCard" style={{ left: "" + dis + "px" }}>
        <Card id={id} />
      </li>
    );
  };
}

function Deck({ cardsId }) {
  let cardMaker = makeCard();

  return <ul className="deck"> {cardsId.map((id) => cardMaker(id))}</ul>;
}

export default Deck;
