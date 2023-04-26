// set up clickable logo
const logo = document.getElementById('logo');
logo?.addEventListener('click', function() {
  window.location = '/';
});

function postData(url, data = {}) {
  return fetch(url, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify(data),
  });
}

function isNumber(value) {
  return typeof value === 'number' && !isNaN(value);
}

function hasFile() {
  const fileInput = document.getElementById('file');
  return fileInput?.value?.length > 0;
}

async function uploadFile(file) {
  const formData = new FormData();
  formData.append('file', file, file.name);

  const response = await fetch('/upload-file', {
    method: 'POST',
    body: formData
  });
  const jsonResponse = await response.json();
  return jsonResponse.file_token;
}

function setLoadingState() {
  const submitButton = document.querySelector('button[type="submit"]')
  submitButton?.setAttribute('disabled', true);
  submitButton.innerHTML = '';
  const spinner = document.createElement('span');
  spinner.setAttribute('role', 'status');
  spinner.classList.add('spinner-border', 'spinner-border-sm');
  submitButton.append(spinner, ' Submitting...');
}
