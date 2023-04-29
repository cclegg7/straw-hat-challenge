

function setProgressBar() {
  const progressBar = document.getElementById('progressBar');
  const today = new Date();
  const challenge_start = new Date("04/15/2023");
  const challenge_end = new Date("08/18/2023");
  const challenge_diff = challenge_end - challenge_start;
  const today_end_diff = today - challenge_start;
  const bar_width = today_end_diff / challenge_diff * 100;
  progressBar.style.width = bar_width.toString().concat("%");   
}

async function populateScoreboard() {
  const scoreboard = document.getElementById("scoreboard");

  const scoresResponse = await fetch("/scores");
  const { scores } = await scoresResponse.json();
  scores.forEach(({ rank, user_id, user_name, character_token, score }) => {
    const row = document.createElement("div");
    row.classList.add('row', 'm-3', 'px-3', 'rounded-pill', 'bg-light');
 

    // add image
    const image = document.createElement("img");
    image.setAttribute("style", "width: 50px; height: 50px");
    image.setAttribute("src", `images/character_icons/${character_token}.svg`);
    const imageCell = document.createElement("span");
    imageCell.appendChild(image);
    imageCell.classList.add('col-2', 'g-0');
    row.appendChild(imageCell);


    // add user
    const userCell = document.createElement("span");
    userCell.appendChild(document.createTextNode(user_name.toUpperCase()));
    userCell.classList.add('col-8', 'text-center', 'align-self-center', 'fw-bold');
    row.appendChild(userCell);

    // add scores
    const scoreCell = document.createElement("span");
    scoreCell.appendChild(document.createTextNode(score));
    row.appendChild(scoreCell);
    scoreCell.classList.add('col-2', 'text-end', 'align-self-center','fw-bold');

    // add row to table
    scoreboard.appendChild(row);

    row.addEventListener('click', async function(e) {
      e.preventDefault();
      window.location = `/profile.html?user_id=${user_id}`;
    });
  });
}

setProgressBar();
populateScoreboard();
