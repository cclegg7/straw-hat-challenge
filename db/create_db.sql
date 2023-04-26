CREATE TABLE
    characters (
        id INT NOT NULL AUTO_INCREMENT,
        token VARCHAR(200) NOT NULL,
        display_name VARCHAR(200) NOT NULL,
        PRIMARY KEY (id)
    );

CREATE TABLE
    users (
        id INT NOT NULL AUTO_INCREMENT,
        name VARCHAR(200) NOT NULL,
        character_id INT,
        toprope INT,
        boulder INT,
        PRIMARY KEY (id),
        FOREIGN KEY (character_id) REFERENCES characters (id)
    );

CREATE TABLE
    runs (
        id INT NOT NULL AUTO_INCREMENT,
        user_id INT,
        date DATE,
        distance FLOAT,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        PRIMARY KEY (id),
        FOREIGN KEY (user_id) REFERENCES users (id)
    );

CREATE TABLE
    climbs (
        id INT NOT NULL AUTO_INCREMENT,
        user_id INT,
        date DATE,
        category INT,
        rating INT,
        is_challenge BOOLEAN NOT NULL DEFAULT FALSE,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        PRIMARY KEY (id),
        FOREIGN KEY (user_id) REFERENCES users (id)
    );

CREATE TABLE
    files (
        id INT NOT NULL AUTO_INCREMENT,
        token VARCHAR(100),
        url VARCHAR(200),
        content_type varchar(100),
        type SMALLINT,
        climb_id INT,
        run_id INT,
        PRIMARY KEY (id),
        UNIQUE KEY (token),
        FOREIGN KEY (climb_id) REFERENCES climbs (id),
        FOREIGN KEY (run_id) REFERENCES runs (id),
        INDEX (token)
    );

CREATE VIEW runs_with_week_info AS 
	SELECT
	    date,
	    CASE
	        WHEN DAYOFWEEK(date) IN (1, 7) THEN WEEK(date, 1) + 1
	        ELSE WEEK(date, 1)
	    END as week_num,
	    CASE
	        WHEN DAYOFWEEK(date) IN (1, 7) THEN 1
	        ELSE 0
	    END as is_weekend,
	    user_id,
	    distance
	FROM
runs; 

CREATE VIEW climbs_with_week_info AS 
	SELECT
	    date,
	    CASE
	        WHEN DAYOFWEEK(date) IN (1, 7) THEN WEEK(date, 1) + 1
	        ELSE WEEK(date, 1)
	    END as week_num,
	    user_id,
	    category,
	    rating,
	    is_challenge
	FROM
climbs; 