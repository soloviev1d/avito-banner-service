CREATE TABLE IF NOT EXISTS "banners" (
  "id" bigserial PRIMARY KEY,
  "title" varchar(512) NOT NULL,
  "data" text NOT NULL,
  "url" varchar(2048) NOT NULL,
  "is_active" boolean DEFAULT true
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
  SELECT 'f139b375-1c04-4295-8897-901653e2bc02' AS token, 1 AS user_role_id
  UNION ALL
  SELECT '247406c0-2591-4e4a-87cc-45803a795dbd', 2
)
INSERT INTO "user_session" (token, user_role_id)
SELECT token, user_role_id
FROM data
WHERE NOT EXISTS (SELECT 1 FROM "user_session");


ALTER TABLE "user_session" ADD FOREIGN KEY ("user_role_id") REFERENCES "user_role" ("id");
