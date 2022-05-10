CREATE TABLE user (
    id int NOT NULL AUTO_INCREMENT,
    uid char NOT NULL UNIQUE,
    nickname char(50),
    PRIMARY KEY (id)
);

CREATE TABLE who (
    id int NOT NULL AUTO_INCREMENT,
    instruction char(50) NOT NULL UNIQUE,
    PRIMARY KEY (id)
);

CREATE TABLE timing (
    id int NOT NULL AUTO_INCREMENT,
    instruction char(100) NOT NULL UNIQUE,
    PRIMARY KEY (id)
);

CREATE TABLE place (
    id int NOT NULL AUTO_INCREMENT,
    instruction char(100) NOT NULL UNIQUE,
    PRIMARY KEY (id)
);

CREATE TABLE what (
    id int NOT NULL AUTO_INCREMENT,
    instruction char(100) NOT NULL UNIQUE,
    PRIMARY KEY (id)
);

CREATE TABLE quest (
    id int NOT NULL AUTO_INCREMENT,
    user_uid char NOT NULL,
    who_id int,
    timing_id int,
    place_id int,
    what_id int NOT NULL,
    comment char(200),
    PRIMARY KEY (id),
    FOREIGN KEY (user_uid) REFERENCES user(uid) ON DELETE CASCADE,
    FOREIGN KEY (who_id) REFERENCES who(id) ON DELETE CASCADE,
    FOREIGN KEY (timing_id) REFERENCES timing(id) ON DELETE CASCADE,
    FOREIGN KEY (place_id) REFERENCES place(id) ON DELETE CASCADE,
    FOREIGN KEY (what_id) REFERENCES what(id) ON DELETE CASCADE
);

CREATE TABLE comment (
    id int NOT NULL AUTO_INCREMENT,
    user_uid char NOT NULL,
    quest_id int NOT NULL,
    comment char(200) NOT NULL,
    PRIMARY KEY (id),
    FOREIGN KEY (user_uid) REFERENCES user(uid) ON DELETE CASCADE,
    FOREIGN KEY (quest_id) REFERENCES quest(id) ON DELETE CASCADE
);

CREATE TABLE nice (
    user_uid char NOT NULL,
    quest_id int NOT NULL,
    PRIMARY KEY (user_uid, quest_id),
    FOREIGN KEY (user_uid) REFERENCES user(uid) ON DELETE CASCADE,
    FOREIGN KEY (quest_id) REFERENCES quest(id) ON DELETE CASCADE
);
