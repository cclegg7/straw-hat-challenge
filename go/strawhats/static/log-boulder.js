function toJSON(formData) {
    return {
        user_id: parseInt(formData.get('user_id')),
        category: 0, // boulder
        rating: parseInt(formData.get('rating')),
        date: formData.get('date'),
        is_challenge: formData.get('challenge') === 'on'
    };
}

function isValid(jsonData) {
    return isNumber(jsonData.user_id) && isNumber(jsonData.category)  && isNumber(jsonData.rating) && jsonData.date?.length > 0;
 }

form = document.getElementById('climbForm');
form.addEventListener('submit', async function(e) {
    e.preventDefault();

    const formData = new FormData(this);
    const json = toJSON(formData);
    if (isValid(json)) {
        setLoadingState();
        const file = formData.get('file');
        if (hasFile()) {
            json.file_token = await uploadFile(file);
        }
        await postData('/climb', json);
        alert('climb submitted!');
        window.location = '/';
      } else {
        alert('please provide all fields');
      }
});
