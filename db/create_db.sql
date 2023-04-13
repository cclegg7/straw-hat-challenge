CREATE TABLE characters (
   id INT NOT NULL AUTO_INCREMENT,
   token VARCHAR(200) NOT NULL,
   display_name VARCHAR(200) NOT NULL,
   PRIMARY KEY (id)
);

CREATE TABLE users (
   id INT NOT NULL AUTO_INCREMENT,
   name VARCHAR(200) NOT NULL,
   character_id INT,
   toprope INT,
   boulder INT,
   PRIMARY KEY (id),
   FOREIGN KEY (character_id) REFERENCES characters (id)
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
   category INT,
   rating INT,
   is_challenge BOOLEAN,
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
);

