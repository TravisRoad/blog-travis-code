import _ from "lodash";
import Graph from "./components/Graph";

function App() {
  const data = _.range(0, 101);
  return (
    <div className="flex flex-col justify-center items-stretch mx-auto w-full py-10 max-w-screen-md px-2">
      <Graph data={data} />
    </div>
  );
}

export default App;
