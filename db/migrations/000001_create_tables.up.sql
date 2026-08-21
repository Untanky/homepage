CREATE TABLE IF NOT EXISTS authors (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  name VARCHAR(128) NOT NULL,
  picture_id UUID
);

CREATE TABLE IF NOT EXISTS blogs (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  title VARCHAR(128) NOT NULL,
  summary TEXT,
  banner_id UUID NOT NULL,
  authority VARCHAR() NOT NULL
);

CREATE TABLE IF NOT EXISTS posts (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  blog_id UUID NOT NULL REFERENCES blogs(id) ON DELETE CASCADE,
  author_id UUID NOT NULL REFERENCES authors(id) ON DELETE RESTRICT,
  slug VARCHAR(32),
  title varchar(256) NOT NULL,
  summary TEXT,
  content TEXT,
  banner_id UUID,
  created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),
  updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now()
);

ALTER TABLE my_table
ADD CONSTRAINT my_table_foreign_key_foo_key UNIQUE (foreign_key, foo);
CREATE INDEX IF NOT EXISTS idx_posts_author_id ON posts(author_id);

CREATE SCHEMA IF NOT EXISTS media;

CREATE TABLE IF NOT EXISTS media.assets ( 
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  path VARCHAR NOT NULL,
  alt VARCHAR NOT NULL DEFAULT '',
  width SMALLINT NOT NULL,
  height SMALLINT NOT NULL,
  created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now()
);

CREATE TABLE IF NOT EXISTS media.asset_versions (
  asset_id UUID NOT NULL REFERENCES media.assets(id) ON DELETE CASCADE,
  mimetype VARCHAR NOT NULL,
  width SMALLINT NOT NULL,
  height SMALLINT NOT NULL,
  data BYTEA NOT NULL,
  PRIMARY KEY (asset_id, width, mimetype)
);
