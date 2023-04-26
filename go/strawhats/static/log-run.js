function toJSON(formData) {
    return {
        user_id: parseInt(formData.get('user_id')),
        distance: parseInt(formData.get('distance')),
        date: formData.get('date'),
    };
}

function isValid(jsonData) {
    return isNumber(jsonData.user_id) && isNumber(jsonData.distance) && jsonData.date?.length > 0;
 }

form = document.getElementById('runForm');
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
        await postData('/run', json);
        alert('run submitted!');
        window.location = '/';
      } else {
        alert('please provide all fields');
      }
});
