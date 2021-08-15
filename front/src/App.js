import InfoPanel from "./components/InfoPanel";
import Table from "./components/Table";
import "./styles/app.css";

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
    <div className="appBody">
      <div className="tablePan">
        <Table decks={decks} />
      </div>
      <div className="infoPan">
        <InfoPanel />
      </div>
    </div>
  );
}

export default App;
