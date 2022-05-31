
CREATE TABLE posts(
    id UUID PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    description VARCHAR(50) NOT NULL,
    user_id UUID NOT NULL
);
CREATE TABLE medias(
    id UUID NOT NULL PRIMARY KEY,
    post_id UUID NOT NULL,
    type TEXT NOT NULL,
    link TEXT NOT NULL,
    FOREIGN KEY(post_id)
      REFERENCES posts(id)
      ON DELETE CASCADE
);
SELECT * FROM posts;
SELECT * FROM medias;