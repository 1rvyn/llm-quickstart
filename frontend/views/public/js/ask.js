console.log("ask.js has loaded")

// on click for the button with id="ask"
$("#ask").on("click", function(event) {
    console.log("ask button clicked")
    console.log($("#question").val())
    //send POST to 127.0.0.1:5000/ask
    $.ajax({
        url: "api/search",
        method: "POST",
        contentType: "application/json",
        data: JSON.stringify({
            question: $("#question").val()
        })
    }).done(function(response) {
        // clear the answer div
        $("#answer").text("")
        console.log(response)
        // display response in the div with id="answer"
        $("#answer").append("<p>" + response.answer + "</p>" + "<br>" + "<p>" + response.refrences + "</p>")
    }
    )
})