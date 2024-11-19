function login() {
    const username = document.getElementById('username').value.trim();
    const password = document.getElementById('password').value;
    
    if (!username || !password) {
        alert('Username and password are required');
        return;
    }

    fetch('/api/v1/login', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({ username, password })
    })
    .then(response => {
        if (!response.ok) {
            throw new Error('Network response was not ok');
        }
        return response.json()
    })
    .then(resp => {
        if (resp.success) {
            alert('Login successful');
            sessionStorage.setItem('token', resp.data.token);
            localStorage.setItem('username', resp.data.username);
            localStorage.setItem('user_id', resp.data.userID);
            localStorage.setItem('logged', true);
            window.location.href = '/static';
        } else {
            localStorage.setItem('logged', false);
            alert('Login failed: ' + resp.message);
        }
    })
    .catch(error => {
        console.error('Error:', error)
        alert('Connection error. Please try again later.');
        localStorage.setItem('logged', false);
    });
}