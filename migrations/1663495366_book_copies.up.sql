CREATE TYPE book_status AS ENUM ('Available','Not Available','Issued');

CREATE TABLE IF NOT EXISTS book_copies(
    ISBN text PRIMARY KEY,
    book_id uuid REFERENCES books(id),
    status book_status DEFAULT 'Available',
    created_at timestamp with time zone NOT NULL DEFAULT current_timestamp,
    updated_at timestamp with time zone NOT NULL 
);
