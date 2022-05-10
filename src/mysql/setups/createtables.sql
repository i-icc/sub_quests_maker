CREATE TABLE user (
    id int NOT NULL PRIMARY KEY,
    uid varchar NOT NULL,
    nickname varchar(50),
);

CREATE TABLE who (
    id int NOT NULL,
    order varchar(100) NOT NULL,
);

CREATE TABLE whene (
    id int NOT NULL,
    order varchar(100) NOT NULL,
);

CREATE TABLE place (
    id int NOT NULL,
    order varchar(100) NOT NULL,
);

CREATE TABLE what (
    id int NOT NULL,
    order varchar(100) NOT NULL,
);

CREATE TABLE quest (
    id int NOT NULL,
    user_uid varchar NOT NULL,
    who_id int,
    whene_id int,
    place_id int,
    what_id int NOT NULL,
    comment varchar(200)
);

CREATE TABLE comment (
    id int NOT NULL,
    quest_id int NOT NULL,
    comment varchar(200) NOT NULL
);

CREATE TABLE nice (user_uid varchar, quest_id int,);