import React, {useEffect, useState} from 'react';
import './App.css';
import Joke from "./Joke";
import Button from "./Button";

function App() {
    return (
        <div className="container">
            <div className="header"></div>
            <div className="content"><Joke /></div>
            <div className="navigation">
                <Button text="&#8249;" float="left"/>
                <Button text="&#8250;" float="right"/>
            </div>
            <div className="footer"></div>
        </div>
    );
}

export default App;
