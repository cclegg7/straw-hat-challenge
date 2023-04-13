# straw-hat-challenge

## How to contribute
1. Create new branch: `git checkout -b <your username>/<branch name>`
2. Make code changes
3. Stage changes you want to put to main: `git add <file name>` / `git add --all` / Use SourceTree to stage changes
4. Commit to your branch: `git commit`
5. Push to remote: `git push`
6. Create pull request on GitHub website

# How to set up local database
1. Install and start MySQL
2. Create database with name `straw_hat_challenge`
3. Make sure it can be accessed with username `root`, no password
4. Run `create_db.sql` (Run `reset_db.sql` first if you need to make schema changes to drop existing tables)
5. Run `insert_base_data.sql`
6. (optional) Run `test_data/seed_test_data.sql`

## How to run server locally
1. Setup local database
2. Install go
3. Navigate to `go/strawhats` directory
4. Run `go run ./server/main/main.go`
5. Navigate to `localhost:8080` in browser

Note: The server does not need to be restarted for changes to static files, but _does_ need to be restarted for any changes to the server (.go files)
