import logo from './logo.svg';
import logogo from './logo-go.png'
import './App.css';

function App() {
  return (
    <div className="App">
      <header className="App-header">
        <img src={logogo} className="App-logo" alt="logo" />
        <h1>
          Famous Personalities
        </h1>
        <img src={logo} className="App-logo" alt="logo" />
      </header>
    </div>
  );
}

export default App;
