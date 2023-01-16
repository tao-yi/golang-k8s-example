CREATE TABLE book (
  id serial,
  isbn varchar NOT NULL,
  title varchar NOT NULL,
  description varchar NOT NULL ,
  created_at timestamptz DEFAULT now()
);