console.log("this is the login page");

document.querySelector('form').addEventListener('submit', function(event) {
    event.preventDefault();

    console.log("login button clicked");

    var username = document.querySelector('#username').value;
    var password = document.querySelector('#password').value;

    console.log(username);
    console.log(password);

    var xhr = new XMLHttpRequest();

    xhr.open("POST", "api/login", true);
    xhr.setRequestHeader("Content-Type", "application/json");

    xhr.onreadystatechange = function() {
        if (xhr.readyState == 4 && xhr.status == 200) {
            var json = JSON.parse(xhr.responseText);
            console.log(json);

            // redirect to /search page
            window.location.href = "/search";
        }
        else if (xhr.readyState == 4) {
            console.log("Request failed with status: " + xhr.status);
        }
    }
    
    var data = JSON.stringify({"username": username, "password": password});

    xhr.send(data);
});
