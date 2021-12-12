CREATE TABLE users
(
  id uuid not null unique,
  name varchar(255) not null,
  username varchar(255) not null unique,
  password_hash varchar(255) not null,
  budget integer default 0
);