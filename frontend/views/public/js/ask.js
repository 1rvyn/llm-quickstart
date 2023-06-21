console.log("ask.js has loaded");

document.querySelector('form').addEventListener('submit', function(event) {
    event.preventDefault();

    console.log("search button clicked");

    var question = document.querySelector('#question').value;

    console.log(question);

    var xhr = new XMLHttpRequest();

    xhr.open("POST", "api/search", true);
    xhr.setRequestHeader("Content-Type", "application/json");

    var answerDiv = document.querySelector('#answer');

     // Show loading image/gif
     if (answerDiv) {
        answerDiv.innerHTML = '<img src="images/rely-logo.jpeg" alt="Loading...">';
    }

    xhr.onreadystatechange = function() {
        if (xhr.readyState == 4 && xhr.status == 200) {
            var json = JSON.parse(xhr.responseText);
            console.log(json);

            // If there's an "answer" div, display the answer there.
            var answerDiv = document.querySelector('#answer');
            if (answerDiv) {
                // assuming the response structure is {"output": "Answer: The answer is..."}
                answerDiv.innerHTML = json.output; 
            }
        }
        else if (xhr.readyState == 4) {
            console.log("Request failed with status: " + xhr.status);
            // If there's an "answer" div, display the error message there.
            var answerDiv = document.querySelector('#answer');
            if (answerDiv) {
                answerDiv.innerHTML = "Error"; 
            }
        }
    }
    
    var data = JSON.stringify({"question": question});

    xhr.send(data);
});
