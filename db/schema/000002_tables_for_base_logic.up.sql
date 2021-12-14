CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE film
(
    id           uuid primary key default uuid_generate_v1(),
    name         varchar(255) not null,
    description  varchar(1024),
    duration     time         not null,
    release_date timestamp
);

CREATE TABLE session
(
    id        uuid primary key default uuid_generate_v1(),
    film_id   uuid references film (id) on delete cascade not null,
    show_date timestamp                                   not null
);

CREATE TABLE user_session
(
    id         uuid primary key default uuid_generate_v1(),
    user_id    uuid references users (id) on delete cascade   not null,
    session_id uuid references session (id) on delete cascade not null,
    place      int                                            not null check ( place > 0 and place < 51 ),
    unique (session_id, place)
)