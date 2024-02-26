var token = "";

async function getData(endpoint) {
    const response = await fetch("http://localhost:8100"+endpoint, {
        headers: {
            "Authorization": "Bearer "+token,
        }
    });
    const data = await response.json();
    console.log(data);
}

async function getToken(clientID, secret) {
    const response = await fetch("http://localhost:8100/token?grant_type=client_credentials&client_id="+clientID+"&client_secret="+secret);
    const data = await response.json();
    
    token = data.access_code
}