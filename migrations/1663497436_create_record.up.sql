CREATE TABLE IF NOT EXISTS records(
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id uuid NOT NULL REFERENCES users(id) ,
    book_copy_id text NOT NULL REFERENCES book_copies(ISBN),
    book_id uuid NOT NULL REFERENCES books(id) ,
    issued_at timestamp  with time zone NOT NULL DEFAULT current_timestamp,
    returned_at timestamp with time zone
);
