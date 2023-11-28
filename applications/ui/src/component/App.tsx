import React, {useEffect} from 'react';
import logo from '../logo.svg';
import './App.css';

function App() {
    useEffect(() => {
        fetch('http://localhost:8080/api/v1/jokes/random')
            .then((response) => response.json())
            .then((data) => {
                console.log(data)
            })
    }, []);
  return (
    <div className="App">
      <header className="App-header">

      </header>
    </div>
  );
}

export default App;
