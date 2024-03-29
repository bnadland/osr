-- migrate:up
CREATE TABLE feeds (
    feed_id INT GENERATED ALWAYS AS IDENTITY, 
    title TEXT NOT NULL,
    link TEXT UNIQUE NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    PRIMARY KEY(feed_id)
);
CREATE TABLE items (
    item_id INT GENERATED ALWAYS AS IDENTITY,
    feed_id INT,
    title TEXT NOT NULL,
    link TEXT UNIQUE NOT NULL,
    categories TEXT[],
    content TEXT NOT NULL DEFAULT '',
    published_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    PRIMARY KEY(item_id),
    CONSTRAINT fk_feed FOREIGN KEY(feed_id) REFERENCES feeds(feed_id)
);
-- migrate:down
DROP TABLE items;
DROP TABLE feeds;