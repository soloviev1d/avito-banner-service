CREATE TABLE "banners" (
  "id" bigserial PRIMARY KEY,
  "title" varchar(512) NOT NULL,
  "data" text NOT NULL,
  "url" varchar(2048) NOT NULL
);

CREATE TABLE "features" (
  "id" integer PRIMARY KEY NOT NULL,
  "banner_id" integer UNIQUE NOT NULL
);

CREATE TABLE "tags" (
  "id" integer NOT NULL,
  "banner_id" integer NOT NULL
);

COMMENT ON COLUMN "banners"."url" IS 'max url length for browsers';

CREATE TABLE "features_banners" (
  "features_banner_id" integer,
  "banners_id" bigserial,
  PRIMARY KEY ("features_banner_id", "banners_id")
);

ALTER TABLE "features_banners" ADD FOREIGN KEY ("features_banner_id") REFERENCES "features" ("banner_id");

ALTER TABLE "features_banners" ADD FOREIGN KEY ("banners_id") REFERENCES "banners" ("id");

ALTER TABLE "tags" ADD FOREIGN KEY ("banner_id") REFERENCES "banners" ("id");
