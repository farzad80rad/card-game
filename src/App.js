import Table from "./components/Table";

function App() {
  let decks = [["h2", "h4", "c4", "s6", "sq"]];
  return (
    <div>
      <header>
        <Table decks={decks} />
      </header>
    </div>
  );
}

export default App;
