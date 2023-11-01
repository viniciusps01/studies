CREATE TABLE tasks (
    id SERIAL UNIQUE NOT NULL,
    description varchar not null,
    done boolean,

    primary key (id)
);