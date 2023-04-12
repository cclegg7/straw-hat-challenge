CREATE TABLE crews (
   id INT NOT NULL AUTO_INCREMENT,
   name VARCHAR(200) NOT NULL,
   PRIMARY KEY (id)
);

CREATE TABLE users (
   id INT NOT NULL AUTO_INCREMENT,
   name VARCHAR(200) NOT NULL,
   crew_id INT,
   toprope VARCHAR(10),
   boulder VARCHAR(10),
   PRIMARY KEY (id),
   FOREIGN KEY (crew_id) REFERENCES crews (id)
);

CREATE TABLE runs (
   id INT NOT NULL AUTO_INCREMENT,
   user_id INT,
   date DATE,
   distance FLOAT,
   created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
   PRIMARY KEY (id),
   FOREIGN KEY (user_id) REFERENCES users (id)
);

CREATE TABLE climbs (
   id INT NOT NULL AUTO_INCREMENT,
   user_id INT,
   date DATE,
   category VARCHAR(10),
   difficulty VARCHAR(10),
   created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
   PRIMARY KEY (id),
   FOREIGN KEY (user_id) REFERENCES users (id)
);


CREATE TABLE climb_files (
   id INT NOT NULL AUTO_INCREMENT,
   climb_id INT,
   url VARCHAR(200),
   PRIMARY KEY (id), 
   FOREIGN KEY (climb_id) REFERENCES climbs (id)
)

