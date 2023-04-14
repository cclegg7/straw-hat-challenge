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
