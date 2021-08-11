import Deck from "./components/Deck";
import "./App.css";

function App() {
  return (
    <div className="App">
      <header className="App-header">
        <Deck cardsId={["h2", "h4", "c4", "s6", "sq"]} />
      </header>
    </div>
  );
}

export default App;
