CREATE TABLE IF NOT EXISTS comments (
  id serial PRIMARY KEY,
  content text NOT NULL,
  post_id int NOT NULL,
  created_at timestamptz NOT NULL DEFAULT NOW(),
  FOREIGN KEY (post_id) REFERENCES posts(id) ON DELETE CASCADE
);