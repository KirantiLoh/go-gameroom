CREATE TABLE IF NOT EXISTS "users" (
  name varchar(100) not null,
  email varchar(150) unique not null,
  password varchar(200) not null,
  is_verified boolean default false,
  id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY
);
