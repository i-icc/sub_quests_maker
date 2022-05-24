DROP TABLE comments;

DROP TABLE nices;

DROP TABLE quests;

DROP TABLE users;

DROP TABLE whos;

DROP TABLE timings;

DROP TABLE places;

DROP TABLE whats;

CREATE TABLE users (
    id int NOT NULL AUTO_INCREMENT,
    uid char(50) NOT NULL UNIQUE,
    nickname char(50) NOT NULL,
    image char(150),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    PRIMARY KEY (id)
);

CREATE TABLE whos (
    id int NOT NULL AUTO_INCREMENT,
    instruction char(50) NOT NULL UNIQUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    PRIMARY KEY (id)
);

CREATE TABLE timings (
    id int NOT NULL AUTO_INCREMENT,
    instruction char(100) NOT NULL UNIQUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    PRIMARY KEY (id)
);

CREATE TABLE places (
    id int NOT NULL AUTO_INCREMENT,
    instruction char(100) NOT NULL UNIQUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    PRIMARY KEY (id)
);

CREATE TABLE whats (
    id int NOT NULL AUTO_INCREMENT,
    instruction char(100) NOT NULL UNIQUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    PRIMARY KEY (id)
);

CREATE TABLE quests (
    id int NOT NULL AUTO_INCREMENT,
    user_uid char(50) NOT NULL,
    who_id int,
    timing_id int,
    place_id int,
    what_id int NOT NULL,
    comment char(200),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    PRIMARY KEY (id),
    FOREIGN KEY (user_uid) REFERENCES users(uid) ON DELETE CASCADE,
    FOREIGN KEY (who_id) REFERENCES whos(id) ON DELETE CASCADE,
    FOREIGN KEY (timing_id) REFERENCES timings(id) ON DELETE CASCADE,
    FOREIGN KEY (place_id) REFERENCES places(id) ON DELETE CASCADE,
    FOREIGN KEY (what_id) REFERENCES whats(id) ON DELETE CASCADE
);

CREATE TABLE comments (
    id int NOT NULL AUTO_INCREMENT,
    user_uid char(50) NOT NULL,
    quest_id int NOT NULL,
    comment char(200) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    PRIMARY KEY (id),
    FOREIGN KEY (user_uid) REFERENCES users(uid) ON DELETE CASCADE,
    FOREIGN KEY (quest_id) REFERENCES quests(id) ON DELETE CASCADE
);

CREATE TABLE nices (
    user_uid char(50) NOT NULL,
    quest_id int NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    PRIMARY KEY (user_uid, quest_id),
    FOREIGN KEY (user_uid) REFERENCES users(uid) ON DELETE CASCADE,
    FOREIGN KEY (quest_id) REFERENCES quests(id) ON DELETE CASCADE
);

/* テスト */
insert users(uid, nickname) values ('icc', 'bo-neko')