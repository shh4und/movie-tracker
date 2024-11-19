function checkTokenExpiration() {
    const token = sessionStorage.getItem('token');
    if (token) {
        const payload = JSON.parse(atob(token.split('.')[1]));
        console.log(`const payload.exp: ${payload.exp}\ndate.now(): ${Date.now()}`)
        if (payload.exp * 1000 < Date.now()) {
            logOut();
        }
    }
}

function logOut() {
    sessionStorage.clear();
    localStorage.clear();
    window.location.href = '/login';
}