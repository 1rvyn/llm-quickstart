document.addEventListener('DOMContentLoaded', function () {
    document.getElementById('upload-form').addEventListener('submit', function (event) {
        event.preventDefault();
        
        // Assuming sliderValue is the value from your slider, ranging from 1-10
        var sliderValue = document.getElementById('slider-id').value; 

        var fileInput = document.getElementById('file');
        var files = fileInput.files; // this is a FileList object, not an array
        var formData = new FormData();

        for (var i = 0; i < files.length; i++) {
            formData.append(`file${i}`, files[i]);
        }

        // Add sliderValue to the API URL as a parameter
        fetch(`/admin/api/upload/${sliderValue}`, { 
            method: 'POST',
            body: formData,
            credentials: 'include'
        })
        .then(response => response.json())
        .then(data => {
            console.log(data);
            // handle response data as needed
        })
        .catch(error => {
            console.error('Error:', error);
        });
    });
});



    var radios = document.getElementsByName('value');
    var valueParagraph = document.getElementById('radio-value');

    for(var i = 0; i < radios.length; i++) {
        radios[i].addEventListener('change', function() {
            valueParagraph.textContent = this.value;
        });
    }
