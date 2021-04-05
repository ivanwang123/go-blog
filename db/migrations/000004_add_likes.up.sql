BEGIN;

CREATE TABLE IF NOT EXISTS likes (
  id serial PRIMARY KEY,
  user_id int NOT NULL,
  post_id int NOT NULL,
  FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
  FOREIGN KEY (post_id) REFERENCES posts(id) ON DELETE CASCADE
);

ALTER TABLE likes ADD CONSTRAINT unique_user_post UNIQUE(user_id, post_id);

COMMIT;