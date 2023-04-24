const mediaContainer = document.getElementById('mediaContainer')

const queryString = window.location.search;
const queryParams = new URLSearchParams(queryString);
const src = queryParams.get('src');
const contentType = queryParams.get('content_type');
if (src?.length > 0) {
    const primaryMimeType = contentType?.split('/')[0];
    if (primaryMimeType === 'image') {
        const img = document.createElement('img');
        img.setAttribute('src', src);
        img.classList.add('full-width-img');
        mediaContainer.appendChild(img);
    } else if (primaryMimeType === 'video') {
        const video = document.createElement('video');
        video.setAttribute('controls', 'true');
        const videoSource = document.createElement('source');
        videoSource.setAttribute('src', src);
        videoSource.setAttribute('type', contentType);
        video.appendChild(videoSource)
        video.classList.add('full-width-img');
        mediaContainer.appendChild(video);
    } else {
        mediaContainer.appendChild(document.createTextNode(`unsupported content type: ${contentType}`));
    }
}

