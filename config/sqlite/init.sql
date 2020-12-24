CREATE TABLE "t_admin" (
  "id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
  "username" TEXT,
  "email" TEXT,
  "password" TEXT
);

CREATE TABLE "t_channel" (
  "id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
  "name" TEXT,
  "path" TEXT,
  "avatar" TEXT,
  "visibility" integer,
  "meta_title" TEXT,
  "meta_description" TEXT,
  "sort" integer,
  "parent_id" INTEGER
);