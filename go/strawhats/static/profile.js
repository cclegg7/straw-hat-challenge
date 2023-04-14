const runsTable = document.getElementById("runsTable");
const bouldersTable = document.getElementById("bouldersTable");
const topropesTable = document.getElementById("topropesTable");

const boulderRatings = {
  0: "VB",
  1: "V0-1",
  2: "V1-2",
  3: "V2-4",
  4: "V4-6",
};

const topropeRatings = {
  0: "5.6",
  1: "5.7",
  2: "5.8",
  3: "5.9",
  4: "5.10",
  5: "5.11",
};

async function fetchRuns() {
  const runsResponse = await fetch("/runs?user_id=2");
  const { runs } = await runsResponse.json();
  return runs;
}

async function fetchBoulders() {
  const bouldersResponse = await fetch("/climbs?user_id=2&category=0");
  const { climbs } = await bouldersResponse.json();
  return climbs;
}

async function fetchTopropes() {
  const topropesResponse = await fetch("/climbs?user_id=2&category=1");
  const { climbs } = await topropesResponse.json();
  return climbs;
}

async function populateTables() {
  const [runs, boulders, topropes] = await Promise.all([
    fetchRuns(),
    fetchBoulders(),
    fetchTopropes(),
  ]);

  populateRuns(runs);
  populateBoulders(boulders);
  populateTopropes(topropes);
}

// convert date string to date: new Date(runs[0].Date).toDateString()

function populateRuns(runs) {
  runs.forEach(({ date, distance, created_at }) => {
    const row = document.createElement("tr");

    // add date
    const dateCell = document.createElement("td");
    const mod_date = new Date(date).toDateString();
    dateCell.appendChild(document.createTextNode(mod_date));
    row.appendChild(dateCell);

    // add distance
    const distanceCell = document.createElement("td");
    console.log(distance);
    distanceCell.appendChild(document.createTextNode(distance));
    row.appendChild(distanceCell);

    // add created at
    const createdAtCell = document.createElement("td");
    const mod_createdAt = new Date(created_at).toDateString();
    createdAtCell.appendChild(document.createTextNode(mod_createdAt));
    row.appendChild(createdAtCell);

    // add row to table
    runsTable.appendChild(row);
  });
}

function populateBoulders(boulders) {
  boulders.forEach(({ date, rating, is_challenge, created_at }) => {
    const row = document.createElement("tr");

    // add date
    const dateCell = document.createElement("td");
    const mod_date = new Date(date).toDateString();
    dateCell.appendChild(document.createTextNode(mod_date));
    row.appendChild(dateCell);

    // add rating
    const ratingCell = document.createElement("td");
    // console.log(distance);
    ratingCell.appendChild(document.createTextNode(boulderRatings[rating]));
    row.appendChild(ratingCell);

    // add rating
    const is_challengeCell = document.createElement("td");
    // console.log(distance);
    is_challengeCell.appendChild(document.createTextNode(is_challenge));
    row.appendChild(is_challengeCell);

    // add created at
    const createdAtCell = document.createElement("td");
    const mod_createdAt = new Date(created_at).toDateString();
    createdAtCell.appendChild(document.createTextNode(mod_createdAt));
    row.appendChild(createdAtCell);

    // add row to table
    bouldersTable.appendChild(row);
  });
}

function populateTopropes(topropes) {
  topropes.forEach(({ date, rating, is_challenge, created_at }) => {
    const row = document.createElement("tr");

    // add date
    const dateCell = document.createElement("td");
    const mod_date = new Date(date).toDateString();
    dateCell.appendChild(document.createTextNode(mod_date));
    row.appendChild(dateCell);

    // add rating
    const ratingCell = document.createElement("td");
    // console.log(distance);
    ratingCell.appendChild(document.createTextNode(topropeRatings[rating]));
    row.appendChild(ratingCell);

    // add rating
    const is_challengeCell = document.createElement("td");
    // console.log(distance);
    is_challengeCell.appendChild(document.createTextNode(is_challenge));
    row.appendChild(is_challengeCell);

    // add created at
    const createdAtCell = document.createElement("td");
    const mod_createdAt = new Date(created_at).toDateString();
    createdAtCell.appendChild(document.createTextNode(mod_createdAt));
    row.appendChild(createdAtCell);

    // add row to table
    topropesTable.appendChild(row);
  });
}

populateTables();
