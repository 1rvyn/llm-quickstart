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

window.onload = function() {
    // Get slider element
    var slider = document.getElementById('slider-id');
    // Get slider value display element
    var sliderValue = document.getElementById('slider-value');

    // Set initial value
    sliderValue.textContent = slider.value;

    // Update the displayed slider value whenever it changes
    slider.oninput = function() {
        sliderValue.textContent = this.value;
    }
}
