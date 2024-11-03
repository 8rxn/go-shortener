import "./App.css";
import Links from "./components/links";
import NewLink from "./components/new";

function App() {
  return (
    <div className="flex flex-col items-center justify-center w-fit mx-auto h-screen max-w-fit mb-4">
      <div className="">
        <h1 className="text-5xl font-bold">Hey sh.aryn.wtf manager</h1>
        <p>Here are all the shortened urls:</p>
      </div>

      <Links></Links>


    </div>
  );
}

export default App;
