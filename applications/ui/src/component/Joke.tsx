import React, {useEffect, useState} from 'react';
import './Joke.css';

function Joke() {
    const [joke, setJoke] = useState('')

    useEffect(() => {
        fetch('http://127.0.0.1.nip.io:8081/api/v1/jokes/random')
            .then((response) => response.json())
            .then((data) => {
                setJoke(data['joke'])
            })
            .catch((err) => {
                console.log(err.message);
            });
    }, []);

    return (
        <div className="joke">{joke}</div>
    );
}

export default Joke;
