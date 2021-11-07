CREATE TABLE myblog
(
  blog_id bigint NOT NULL,
  title text ,
  content text ,
  CONSTRAINT myblog_pkey PRIMARY KEY (blog_id)
);

INSERT INTO myblog(id, title, content) VALUES
(1, 'test', 'testtesttesttesttest'),
(2, 'test2', 'testtesttesttesttest');