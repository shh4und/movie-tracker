import React, { useEffect, useState } from 'react';
import axios from 'axios';

const HomePage: React.FC = () => {
    const [movies, setMovies] = useState([]);

    useEffect(() => {
        axios.get('http://localhost:8080/api/movies')
            .then(response => {
                setMovies(response.data);
            })
            .catch(error => {
                console.error("There was an error fetching the movies!", error);
            });
    }, []);

    return (
        <div>
            <h1>Movies</h1>
            <ul>
                {movies.map((movie: any) => (
                    <li key={movie.id}>{movie.title}</li>
                ))}
            </ul>
        </div>
    );
}

export default HomePage;