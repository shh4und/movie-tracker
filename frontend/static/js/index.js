function register() {
    window.location.href = '/register';
}

function login() {
    window.location.href = '/login';
}

function searchTitle() {
    const title = document.getElementById('searchInput').value;
    // Hide comments section when starting new search
    document.getElementById('commentsSection').style.display = 'none';

    fetch(`/api/v1/titles/search?title=${title}`)
        .then(response => response.json())
        .then(resp => {
            if (resp.success) {
                displayMovieInfo(resp.data);
                // Show comments section only on successful search
                document.getElementById('commentsSection').style.display = 'block';
            } else {
                alert("Title not found");
            }
        })
        .catch(error => {
            console.error('Error:', error);
            document.getElementById('commentsSection').style.display = 'none';

        });
}

function displayMovieInfo(data) {
    const movieInfo = document.getElementById('movieInfo');
    movieInfo.innerHTML = `
        <img src="${data.Poster}" alt="${data.Title}">
        <div><strong>Title:</strong> ${data.Title}</div>
        <div><strong>Year:</strong> ${data.Year}</div>
        <div><strong>Released:</strong> ${data.Released}</div>
        <div><strong>Runtime:</strong> ${data.Runtime}</div>
        <div><strong>Genre:</strong> ${data.Genre}</div>
        <div><strong>Director:</strong> ${data.Director}</div>
    `;

    const actionButtons = document.getElementById('actionButtons');
    actionButtons.innerHTML = `
        <button onclick="addRating('${data.imdbID}', event)">Rating</button>
        <button onclick="addFavorite('${data.imdbID}')">Favorites</button>
        <button>Watch Later</button>
        <button>Watched</button>
    `;

    const commentsSection = document.getElementById('commentsSection');
    commentsSection.style.display = 'block';

    document.getElementById('commentsList').innerHTML = '';
    document.getElementById('commentInput').value = '';

    // Adicionar classes para animação
    movieInfo.classList.add('fade-in');
    actionButtons.classList.add('fade-in');
}

function addRating(imdbID, event) {
    // Get button from passed event
    const button = event.target;
    button.classList.add('loading');

    // Validate rating input
    const rating = prompt("Enter your rating (1-10):");
    const ratingNum = parseInt(rating);

    if (!rating || isNaN(ratingNum) || ratingNum < 1 || ratingNum > 10) {
        alert("Please enter a valid rating between 1 and 10");
        button.classList.remove('loading');
        return;
    }

    fetch('/api/v1/rate', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
            'Authorization': 'Bearer ' + sessionStorage.getItem('token')
        },
        body: JSON.stringify({ title_imdbID: imdbID, rating: ratingNum })
    })
        .then(response => response.json())
        .then(resp => {
            if (resp.success) {
                alert('Rating added successfully');
            } else {
                alert('Failed to add rating: ' + resp.message);
            }
        })
        .catch(error => {
            console.error('Error:', error);
            alert('Failed to add rating');
        })
        .finally(() => {
            button.classList.remove('loading');
        });
}

function addFavorite(titleIMDbID) {
    const userID = localStorage.getItem('user_id')
    if (userID) {
        fetch('/api/v1/favorite', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
                'Authorization': 'Bearer ' + sessionStorage.getItem('token')
            },
            body: JSON.stringify({ title_imdbID: titleIMDbID })
        })
            .then(response => response.json())
            .then(resp => {
                if (resp.success) {
                    alert('Favorite added successfully');
                    console.log(resp.data)
                } else {
                    alert('Failed to add rating: ' + resp.message);
                }
            })
            .catch(error => console.error('Error:', error));
    }
}

function addComment() {
    const comment = document.getElementById('commentInput').value;
    if (comment) {
        const commentsList = document.getElementById('commentsList');
        const li = document.createElement('li');
        li.textContent = comment;
        commentsList.appendChild(li);
        document.getElementById('commentInput').value = '';
    }
}

function checkLogged() {
    checkTokenExpiration();
    const logged = localStorage.getItem('logged');
    const username = localStorage.getItem('username');
    if (logged && username) {
        document.getElementById('usernameDisplay').textContent = `Logged in as: ${username}`;
        document.getElementById('registerButton').style.display = 'none';
        document.getElementById('loginButton').style.display = 'none';
        document.getElementById('logOutButton').style.display = 'inline-block';
    } else {
        document.getElementById('usernameDisplay').textContent = '';
        document.getElementById('registerButton').style.display = '';
        document.getElementById('loginButton').style.display = '';
        document.getElementById('logOutButton').style.display = 'none';
    }
}
setInterval(checkTokenExpiration, 60000);


window.onload = function () {
    checkLogged();
    checkTokenExpiration();
};