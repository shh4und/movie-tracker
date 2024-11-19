function register() {
    const username = document.getElementById('username').value;
    const email = document.getElementById('email').value;
    const password = document.getElementById('password').value;

    fetch('/api/v1/register', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({ username, email, password })
    })
    .then(response => response.json())
    .then(resp => {
        if (resp.success) {
            alert('Registration successful');
            window.location.href = '/login';
        } else {
            alert('Registration failed: ' + resp.message);
        }
    })
    .catch(error => console.error('Error:', error));
}