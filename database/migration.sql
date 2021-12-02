CREATE TABLE IF NOT EXISTS `genre`
(
    genre_id   bigint auto_increment,
    genreName varchar(100) NOT NULL DEFAULT '0',
    PRIMARY KEY (`genre_id`)
);

INSERT INTO `genre` (`genreName`)
VALUES ('Adventure'),
       ('Classics'),
       ('Fantasy');

CREATE TABLE IF NOT EXISTS `books`
(
    id   bigint auto_increment,
    name varchar(100) NOT NULL DEFAULT 'None' UNIQUE,
    price float NOT NULL DEFAULT '0' CHECK (`price` >= 0),
    genre bigint NOT NULL DEFAULT '0' CHECK (`genre` > 0),
    amount bigint NOT NULL DEFAULT '0' CHECK (`amount` >= 0),
    FOREIGN KEY (genre) REFERENCES genre(genre_id),
    PRIMARY KEY (`id`)
);

INSERT INTO `books` (id, `name`, `price`, `genre`, `amount`)
VALUES ('1', 'The Three Musketeers', '10.44', '1', '5');

INSERT INTO `books` (id, `name`, `price`, `genre`, `amount`)
VALUES ('2', 'Roadside Picnic', '12.21', '3', '2');