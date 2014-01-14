# my Go demo app

### Just playing around, learning Cassandra and how Go might be useful in a RESTful / web framework context.

### Cassandra Schema used:

    CREATE KEYSPACE mykeyspace WITH replication = {
      'class': 'SimpleStrategy',
      'replication_factor': '1'
    };

    USE mykeyspace;

    CREATE TABLE posts (
      id int,
      content text,
      date timestamp,
      tags text,
      title text,
      PRIMARY KEY (id)
    ) WITH
      comment='Blog posts';

    CREATE INDEX posts_title ON posts (title);

    INSERT INTO posts (id, content, date, tags, title)
    VALUES (1, 'hello world', '2014-01-14 10:00-0600', '#demo #revel #go', 'First Post');

### Dependencies

    go get github.com/robfig/revel

    go get tux21b.org/v1/gocql

#### LICENSE

Copyright © 2014 Kevin Williams

MIT License