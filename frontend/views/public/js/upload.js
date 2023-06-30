// req to send new files
document.getElementById('upload-form').addEventListener('submit', function (event) {
    event.preventDefault();

    var selectedRadio = document.querySelector('input[name="value"]:checked');
    if (!selectedRadio) {
        alert("Please select a value.");
        return;
    }

    var sliderValue = selectedRadio.value; 

    var fileInput = document.getElementById('file');
    var files = fileInput.files; 
    var formData = new FormData();

    for (var i = 0; i < files.length; i++) {
        formData.append(`file${i}`, files[i]);
    }

    fetch(`/admin/api/upload/${sliderValue}`, { 
        method: 'POST',
        body: formData,
        credentials: 'include'
    })
    .then(response => response.json())
    .then(data => {
        console.log(data);
    })
    .catch(error => {
        console.error('Error:', error);
    });
});



    // makes radio buttons update
    var radios = document.getElementsByName('value');
    var valueParagraph = document.getElementById('radio-value');

    for(var i = 0; i < radios.length; i++) {
        radios[i].addEventListener('change', function() {
            valueParagraph.textContent = this.value;
        });
    }



    for(var i = 0; i < radios.length; i++) {
        radios[i].addEventListener('change', function() {
            valueParagraph.textContent = this.value;
            
            // Fetch the files when a radio button is selected
            fetch(`/admin/api/foldercontents/${this.value}`)
                .then(response => response.json())
                .then(files => populateFileList(files))
                .catch(error => console.error('Error:', error));
        });
    }
    
    document.querySelectorAll('.delete-button').forEach(button => {
        button.addEventListener('click', function() {
            var fileName = this.getAttribute('data-file');
            var folderId = document.querySelector('input[name="value"]:checked').value;

            // Delete the file
            fetch('/admin/api/files', {
                method: 'DELETE',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify({ file: fileName, folderId: folderId }),
            })
            .then(response => response.json())
            .then(response => console.log('Success:', response))
            .catch(error => console.error('Error:', error));
        });
    });

    function populateFileList(files) {
        var fileList = document.getElementById('file-list');
        fileList.innerHTML = '';
        files.forEach(fileName => {
            var fileElement = document.createElement('div');
            fileElement.innerHTML = `
                <label class="block">
                    <input type="checkbox" name="file" value="${fileName}">
                    ${fileName}
                    <button class="delete-button" data-file="${fileName}">Delete</button>
                </label>
            `;
            fileList.appendChild(fileElement);
        });
    
        // Add event listeners to delete buttons after the file list is populated
        Array.from(fileList.getElementsByClassName('delete-button')).forEach(button => {
            button.addEventListener('click', function() {
                var fileName = this.getAttribute('data-file');
                var folderId = document.querySelector('input[name="value"]:checked').value;
        
                // Delete the file
                fetch('/admin/api/files', {
                    method: 'DELETE',
                    headers: {
                        'Content-Type': 'application/json',
                    },
                    body: JSON.stringify({ files: [fileName], folderId: folderId }),
                })
                .then(response => response.json())
                .then(response => console.log('Success:', response))
                .catch(error => console.error('Error:', error));
            });
        });
    }