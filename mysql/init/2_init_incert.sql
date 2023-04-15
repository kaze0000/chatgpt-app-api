-- 投入用データ初期化
TRUNCATE TABLE users;

INSERT INTO users
  (name, email, password)
VALUES
  ('user1', 'user1@ex.com', 'password'),
  ('user2', 'user2@ex.com', 'password'),
  ('user3', 'user3@ex.com', 'password');

