import Table from "./components/Table";

function App() {
  let decks = [
    [
      "h2",
      "h4",
      "c4",
      "s6",
      "sq",
      "d4",
      "h6",
      "s5",
      "h10",
      "d10",
      "s10",
      "hq",
      "da",
    ],
  ];
  return (
    <div>
      <header>
        <Table decks={decks} />
      </header>
    </div>
  );
}

export default App;
