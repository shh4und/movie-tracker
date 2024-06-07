import React from 'react';
import { useParams } from 'react-router-dom';

const MoviePage: React.FC = () => {
    const { id } = useParams<{ id: string }>();

    return (
        <div>
            <h1>Movie Page</h1>
            <p>Movie ID: {id}</p>
        </div>
    );
}

export default MoviePage;