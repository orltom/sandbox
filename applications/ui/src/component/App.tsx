import React, {useEffect, useState} from 'react';
import './App.css';
import Joke from "./Joke";

function App() {
    return (
        <div className="container">
            <div className="header"></div>
            <div className="content"><Joke /></div>
            <div className="navigation"></div>
            <div className="footer"></div>
        </div>
    );
}

export default App;
