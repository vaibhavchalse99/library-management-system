CREATE TYPE book_status AS ENUM ('Available','Not Available','Issued');

CREATE TABLE IF NOT EXISTS book_copies(
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    book_id uuid REFERENCES users(id),
    status book_status DEFAULT 'Available',
    created_at timestamp with time zone NOT NULL DEFAULT current_timestamp,
    updated_at timestamp with time zone NOT NULL 
);
