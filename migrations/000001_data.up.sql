CREATE TABLE IF NOT EXISTS "post" (
    "id" SERIAL PRIMARY KEY NOT NULL,
    "user_id" INTEGER NOT NULL,
    "title" TEXT NOT NULL,
    "body" TEXT NOT NULL,
    "page" INTEGER NOT NULL
);
