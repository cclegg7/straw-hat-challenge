const runsTable = document.getElementById("runsTable");
const bouldersTable = document.getElementById("bouldersTable");
const topropesTable = document.getElementById("topropesTable");

async function fetchRuns() {
  const runsResponse = await fetch('/runs');
  const { runs } = await runsResponse.json()
  return runs;
}

async function fetchBoulders() {
  const bouldersResponse = await fetch('/boulders');
  const { boulders } = await bouldersResponse.json()
  return boulders;
}

async function fetchTopropes() {
  const topropesResponse = await fetch('/topropes');
  const { topropes } = await topropesResponse.json()
  return topropes;
}

async function populateTables() {
  const [runs, boulders, topropes] = await Promise.all([fetchRuns(), fetchBoulders(), fetchTopropes()]);

  // populate all three tables, the same way we do in `index.js`
}

populateTables();
