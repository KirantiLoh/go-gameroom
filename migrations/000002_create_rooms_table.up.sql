CREATE TABLE IF NOT EXISTS "rooms" ( 
  id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  uuid UUID DEFAULT gen_random_uuid(),
  name VARCHAR(100) NOT NULL,
  description VARCHAR(500),
  created_at timestamp DEFAULT now(),

  leader_id BIGINT NOT NULL,
  FOREIGN KEY(leader_id) REFERENCES users(id) ON DELETE CASCADE
);

