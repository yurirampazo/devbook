USE devbook;

DROP TABLE IF EXISTS followers; --seguimores

CREATE TABLE followers(
  user_id int not null,
  follower_id int not null,
  
  FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
  FOREIGN KEY (follower_id) REFERENCES users(id) ON DELETE CASCADE,

  PRIMARY KEY(user_id, follower_id)
)ENGINE=INNODB;
