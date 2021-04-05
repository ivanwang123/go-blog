CREATE TABLE IF NOT EXISTS posts (
  id serial PRIMARY KEY,
  title varchar(255) NOT NULL,
  content text NOT NULL,
  user_id int NOT NULL,
  created_at timestamptz NOT NULL DEFAULT NOW(),
  FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);