CREATE TABLE IF NOT EXISTS "banners" (
  "id" bigserial PRIMARY KEY,
  "title" varchar(512) NOT NULL,
  "data" text NOT NULL,
  "url" varchar(2048) NOT NULL
);

CREATE TABLE IF NOT EXISTS "features" (
  "id" integer PRIMARY KEY NOT NULL,
  "banner_id" integer UNIQUE NOT NULL
);

CREATE TABLE IF NOT EXISTS "tags" (
  "id" integer NOT NULL,
  "banner_id" integer NOT NULL
);

COMMENT ON COLUMN "banners"."url" IS 'max url length for browsers';

CREATE TABLE IF NOT EXISTS "features_banners" (
  "features_banner_id" integer,
  "banners_id" bigserial,
  PRIMARY KEY ("features_banner_id", "banners_id")
);

ALTER TABLE "features_banners" ADD FOREIGN KEY ("features_banner_id") REFERENCES "features" ("banner_id");

ALTER TABLE "features_banners" ADD FOREIGN KEY ("banners_id") REFERENCES "banners" ("id");

ALTER TABLE "tags" ADD FOREIGN KEY ("banner_id") REFERENCES "banners" ("id");


CREATE TABLE IF NOT EXISTS "user_session" (
  "token" uuid PRIMARY KEY,
  "user_role_id" integer NOT NULL
);

CREATE TABLE IF NOT EXISTS "user_role" (
  "id" bigserial PRIMARY KEY,
  "value" varchar(128) UNIQUE NOT NULL 
);

INSERT INTO
  "user_role"(value)
VALUES
  ('admin'),
  ('user')
ON CONFLICT
  DO NOTHING;

WITH data AS (
  SELECT gen_random_uuid() AS token, 1 AS user_role_id
  UNION ALL
  SELECT gen_random_uuid(), 2
)
INSERT INTO "user_session" (token, user_role_id)
SELECT token, user_role_id
FROM data
WHERE NOT EXISTS (SELECT 1 FROM "user_session");


ALTER TABLE "user_session" ADD FOREIGN KEY ("user_role_id") REFERENCES "user_role" ("id");
