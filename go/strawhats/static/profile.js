const queryString = window.location.search;
const queryParams = new URLSearchParams(queryString)
let userID = queryParams.get('user_id');
if (!userID) {
  queryParams.set('user_id', "1");
  userID = 1;
}
const userSelect = document.getElementById('userSelect');
userSelect.value=userID;
userSelect.addEventListener('change', function(event) {
  window.location = `/profile.html?user_id=${userSelect.value}`;
});

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
  const runsResponse = await fetch(`/runs?user_id=${userID}`);
  const { runs } = await runsResponse.json();
  return runs;
}

async function fetchBoulders() {
  const bouldersResponse = await fetch(`/climbs?user_id=${userID}&category=0`);
  const { climbs } = await bouldersResponse.json();
  return climbs;
}

async function fetchTopropes() {
  const topropesResponse = await fetch(`/climbs?user_id=${userID}&category=1`);
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

function formatDate(date) {
  return `${date.getUTCMonth() + 1}/${date.getUTCDate()}/${date.getUTCFullYear()}`
}

function populateRuns(runs) {
  runs?.forEach(({ date, distance, created_at }) => {
    const row = document.createElement("tr");

    // add date
    const dateCell = document.createElement("td");
    const mod_date = formatDate(new Date(date));
    dateCell.appendChild(document.createTextNode(mod_date));
    row.appendChild(dateCell);

    // add distance
    const distanceCell = document.createElement("td");
    console.log(distance);
    distanceCell.appendChild(document.createTextNode(distance));
    row.appendChild(distanceCell);

    // add created at
    const createdAtCell = document.createElement("td");
    const mod_createdAt = new Date(created_at).toString();
    createdAtCell.appendChild(document.createTextNode(mod_createdAt));
    row.appendChild(createdAtCell);

    // add row to table
    runsTable.appendChild(row);
  });
}

function populateBoulders(boulders) {
  boulders?.forEach(({ date, rating, is_challenge, created_at }) => {
    const row = document.createElement("tr");

    // add date
    const dateCell = document.createElement("td");
    const mod_date = formatDate(new Date(date));
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
    const mod_createdAt = new Date(created_at).toString();
    createdAtCell.appendChild(document.createTextNode(mod_createdAt));
    row.appendChild(createdAtCell);

    // add row to table
    bouldersTable.appendChild(row);
  });
}

function populateTopropes(topropes) {
  topropes?.forEach(({ date, rating, is_challenge, created_at }) => {
    const row = document.createElement("tr");

    // add date
    const dateCell = document.createElement("td");
    const mod_date = formatDate(new Date(date));
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
    const mod_createdAt = new Date(created_at).toString();
    createdAtCell.appendChild(document.createTextNode(mod_createdAt));
    row.appendChild(createdAtCell);

    // add row to table
    topropesTable.appendChild(row);
  });
}

populateTables();
