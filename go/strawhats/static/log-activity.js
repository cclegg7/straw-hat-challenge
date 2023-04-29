function setUpRunForm() {
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
    
    const runForm = document.getElementById('runForm');
    runForm.addEventListener('submit', async function(e) {
        e.preventDefault();
    
        const formData = new FormData(this);
        const json = toJSON(formData);
        if (isValid(json)) {
            setLoadingState('runForm');
            const file = formData.get('file');
            if (hasFile('runForm')) {
                json.file_token = await uploadFile(file);
            }
            await postData('/run', json);
            alert('run submitted!');
            window.location = '/';
          } else {
            alert('please provide all fields');
          }
    });
}

function setUpBoulderForm() {
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
    
    const boulderForm = document.getElementById('boulderForm');
    boulderForm.addEventListener('submit', async function(e) {
        e.preventDefault();
    
        const formData = new FormData(this);
        const json = toJSON(formData);
        if (isValid(json)) {
            setLoadingState('boulderForm');
            const file = formData.get('file');
            if (hasFile('boulderForm')) {
                json.file_token = await uploadFile(file);
            }
            await postData('/climb', json);
            alert('climb submitted!');
            window.location = '/';
          } else {
            alert('please provide all fields');
          }
    });    
}

function setUpTopRopeForm() {
    function toJSON(formData) {
        return {
          user_id: parseInt(formData.get('user_id')),
          category: 1, // top rope
          rating: parseInt(formData.get('rating')),
          date: formData.get('date'),
          is_challenge: formData.get('challenge') === 'on'
        };
      }
      
      function isValid(jsonData) {
        return isNumber(jsonData.user_id) && isNumber(jsonData.category)  && isNumber(jsonData.rating) && jsonData.date?.length > 0;
      }
      
      const topRopeForm = document.getElementById('topRopeForm');
      topRopeForm.addEventListener('submit', async function(e) {
        e.preventDefault();
      
        const formData = new FormData(this);
        const json = toJSON(formData);
        if (isValid(json)) {
          setLoadingState('topRopeForm');
          const file = formData.get('file');
          if (hasFile('topRopeForm')) {
            json.file_token = await uploadFile(file);
          }
          await postData('/climb', json);
          alert('climb submitted!');
          window.location = '/';
        } else {
          alert('please provide all fields');
        }
      });
}

setUpRunForm();
setUpTopRopeForm();
setUpBoulderForm();

const formIds = ['runForm', 'boulderForm', 'topRopeForm'];
const activitySelect = document.getElementById('activitySelect');
activitySelect.addEventListener('change', function(event) {
  const selectedForm = activitySelect.value;
  formIds.forEach(formId => {
      const form = document.getElementById(formId);
      if (selectedForm === formId) {
        form.style.display = 'block';
      } else {
        form.style.display = 'none';
      }
  })
});