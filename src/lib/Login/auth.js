export async function auth(password) {
    const response = await fetch('/api/auth', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({
            password: password
        })
    });
    const json = await response.json();
    console.log(json);
    return json;
}