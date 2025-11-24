USE devbook;

INSERT INTO users (name, nick, email, password)
VALUES
('user3', "user_3", 'user3@gmail.com', '$2a$10$5f3xzaYBBOFtYC0tgyUQ6.jK8Urey1JLiynEsE2zWjCQaTj0fgtaG'),
('user4', "user_4", 'user4@gmail.com', '$2a$10$5f3xzaYBBOFtYC0tgyUQ6.jK8Urey1JLiynEsE2zWjCQaTj0fgtaG'),
('user5', "user_5", 'user5@gmail.com', '$2a$10$5f3xzaYBBOFtYC0tgyUQ6.jK8Urey1JLiynEsE2zWjCQaTj0fgtaG');

INSERT INTO followers(user_id, follower_id)
VALUES
(1,2),
(3,1),
(1,3),
(2,1);



