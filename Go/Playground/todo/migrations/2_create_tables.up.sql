CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE user_account (
    id uuid not null default uuid_generate_v4(),
    email text not null unique,
    first_name text not null,
    last_name text not null,
    password text not null,
    role_id int not null default 1,

     primary key (id),
     CONSTRAINT fk_role 
        FOREIGN KEY (role_id)
            REFERENCES role(id)
);

CREATE TABLE task (
    id SERIAL UNIQUE NOT NULL,
    user_id uuid not null,
    title text not null,
    description text,
    done boolean default false,

    primary key (id),
    foreign key (user_id) references user_account(id)
);