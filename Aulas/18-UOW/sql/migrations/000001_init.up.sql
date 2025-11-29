CREATE TABLE categories(
    id varchar(36) NOT NULL PRIMARY KEY,
    name text NOT NULL,
    description text
);


CREATE TABLE courses(
    id varchar(36) NOT NULL PRIMARY KEY,
    category_id varchar(36) NOT NULL,
    name text NOT NULL,
    description text,
    FOREIGN KEY (category_id) references categories(id)

);