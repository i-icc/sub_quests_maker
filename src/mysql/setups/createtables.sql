DROP TABLE comment;

DROP TABLE nice;

DROP TABLE quest;

DROP TABLE user;

DROP TABLE who;

DROP TABLE timing;

DROP TABLE place;

DROP TABLE what;

CREATE TABLE user (
    id int NOT NULL AUTO_INCREMENT,
    uid char(50) NOT NULL UNIQUE,
    nickname char(50) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    PRIMARY KEY (id)
);

CREATE TABLE who (
    id int NOT NULL AUTO_INCREMENT,
    instruction char(50) NOT NULL UNIQUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    PRIMARY KEY (id)
);

CREATE TABLE timing (
    id int NOT NULL AUTO_INCREMENT,
    instruction char(100) NOT NULL UNIQUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    PRIMARY KEY (id)
);

CREATE TABLE place (
    id int NOT NULL AUTO_INCREMENT,
    instruction char(100) NOT NULL UNIQUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    PRIMARY KEY (id)
);

CREATE TABLE what (
    id int NOT NULL AUTO_INCREMENT,
    instruction char(100) NOT NULL UNIQUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    PRIMARY KEY (id)
);

CREATE TABLE quest (
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
    FOREIGN KEY (user_uid) REFERENCES user(uid) ON DELETE CASCADE,
    FOREIGN KEY (who_id) REFERENCES who(id) ON DELETE CASCADE,
    FOREIGN KEY (timing_id) REFERENCES timing(id) ON DELETE CASCADE,
    FOREIGN KEY (place_id) REFERENCES place(id) ON DELETE CASCADE,
    FOREIGN KEY (what_id) REFERENCES what(id) ON DELETE CASCADE
);

CREATE TABLE comment (
    id int NOT NULL AUTO_INCREMENT,
    user_uid char(50) NOT NULL,
    quest_id int NOT NULL,
    comment char(200) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    PRIMARY KEY (id),
    FOREIGN KEY (user_uid) REFERENCES user(uid) ON DELETE CASCADE,
    FOREIGN KEY (quest_id) REFERENCES quest(id) ON DELETE CASCADE
);

CREATE TABLE nice (
    user_uid char(50) NOT NULL,
    quest_id int NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    PRIMARY KEY (user_uid, quest_id),
    FOREIGN KEY (user_uid) REFERENCES user(uid) ON DELETE CASCADE,
    FOREIGN KEY (quest_id) REFERENCES quest(id) ON DELETE CASCADE
);

/* テスト */
insert user(uid, nickname) values ('icc', 'bo-neko')