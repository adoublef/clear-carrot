-- CREATE USER docker;
-- CREATE DATABASE docker;
-- GRANT ALL PRIVILEGES ON DATABASE docker TO docker;
CREATE TABLE IF NOT EXISTS testdb (id int, name varchar(255));
INSERT INTO testdb (id, name) VALUES (1, 'test')