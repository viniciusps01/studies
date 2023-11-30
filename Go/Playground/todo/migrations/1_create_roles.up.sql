CREATE TABLE role(
    id INT NOT NULL UNIQUE,
    name TEXT NOT NULL,

    PRIMARY KEY (id)
);

CREATE TABLE permission(
    id INT NOT NULL UNIQUE,
    name TEXT NOT NULL,

    PRIMARY KEY (id)
);


CREATE TABLE role_permission(
    role_id INT NOT NULL,
    permission_id INT NOT NULL,

    CONSTRAINT fk_role
        FOREIGN KEY (role_id) 
            REFERENCES role(id),
    CONSTRAINT fk_permission
        FOREIGN KEY (permission_id) 
            REFERENCES permission(id)
);

INSERT INTO role(id, name) VALUES (1, 'basic'), (2, 'admin');
INSERT INTO permission(id, name) VALUES (1, 'read_all_users');
INSERT INTO role_permission(role_id, permission_id) VALUES (2, 1);