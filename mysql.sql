CREATE DATABASE cetec;
USE cetec;

CREATE TABLE person (
    id INT NOT NULL AUTO_INCREMENT,
    name VARCHAR(255),
    age INT,
    PRIMARY KEY (id)
);

INSERT INTO person (id, name, age) VALUES (1, 'mike', 31), (2, 'John', 45), (3, 'Joseph', 20);

CREATE TABLE phone (
    id INT NOT NULL AUTO_INCREMENT,
    person_id INT,
    number VARCHAR(255),
    PRIMARY KEY (id)
);

INSERT INTO phone (id, person_id, number) VALUES (1, 1, '444-444-4444'), (2, 2, '123-444-7777'), (3, 3, '445-222-1234');

CREATE TABLE address (
    id INT NOT NULL AUTO_INCREMENT,
    city VARCHAR(255),
    state VARCHAR(255),
    street1 VARCHAR(255),
    street2 VARCHAR(255),
    zip_code VARCHAR(255),
    PRIMARY KEY (id)
);

INSERT INTO address (id, city, state, street1, street2, zip_code) VALUES 
(1, 'Eugene', 'OR', '111 Main St', '', '98765'),
(2, 'Sacramento', 'CA', '432 First St', 'Apt 1', '22221'),
(3, 'Austin', 'TX', '213 South 1st St', '', '78704');

CREATE TABLE address_join (
    id INT NOT NULL AUTO_INCREMENT,
    person_id INT,
    address_id INT,
    PRIMARY KEY (id)
);

INSERT INTO address_join (id, person_id, address_id) VALUES (1, 1, 3), (2, 2, 1), (3, 3, 2);
