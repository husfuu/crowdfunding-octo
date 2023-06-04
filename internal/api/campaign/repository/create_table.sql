CREATE TABLE [IF NOT EXISTS] user (
    user_id serial PRIMARY KEY,
    name varchar(50) NOT NULL,
    password varchar(50) NOT NULL,
    email varchar(255) UNIQUE NOT NULL,
    avatar_file_name varchar(255),
    role varchar NOT NULL, 
)