const scoreboard = document.getElementById("scoreboard");

async function populateScoreboard() {
  const scoresResponse = await fetch("/scores");
  const { scores } = await scoresResponse.json();
  scores.forEach(({ rank, user_id, user_name, character_token, score }) => {
    const row = document.createElement("tr");
    row.classList.add('scoreboard__row');

    // add rank
    const rankCell = document.createElement("td");
    rankCell.appendChild(document.createTextNode(rank));
    row.appendChild(rankCell);

    // add image
    const image = document.createElement("img");
    image.setAttribute("style", "width: 50px; height: 50px");
    image.setAttribute("src", `images/character_icons/${character_token}.svg`);
    const imageCell = document.createElement("td");
    imageCell.appendChild(image);
    row.appendChild(imageCell);

    // add user
    const userCell = document.createElement("td");
    userCell.appendChild(document.createTextNode(user_name));
    row.appendChild(userCell);

    // add score
    const scoreCell = document.createElement("td");
    scoreCell.appendChild(document.createTextNode(score));
    row.appendChild(scoreCell);

    // add row to table
    scoreboard.appendChild(row);

    row.addEventListener('click', async function(e) {
      e.preventDefault();
      window.location = `/profile.html?user_id=${user_id}`;
    });
  });
}

populateScoreboard();
