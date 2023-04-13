function toJSON(formData) {
  return {
    user_id: parseInt(formData.get('user_id')),
    category: 1, // top rope
    rating: parseInt(formData.get('rating')),
    date: formData.get('date'),
    is_challenge: formData.get('challenge') === 'on'
  };
}

function isNumber(value) {
  return typeof value === 'number' && !isNaN(value);
}

function isValid(jsonData) {
  return isNumber(jsonData.user_id) && isNumber(jsonData.category)  && isNumber(jsonData.rating) && jsonData.date?.length > 0;
}

function postData(data = {}) {
  return fetch('/climb', {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify(data),
  });
}

form = document.getElementById('climbForm');
form.addEventListener('submit', async function(e) {
  e.preventDefault();

  const formData = new FormData(this);
  const json = toJSON(formData);
  if (isValid(json)) {
    await postData(json);
    alert('climb submitted!');
    window.location = '/';
  } else {
    alert('please provide all fields');
  }
});
