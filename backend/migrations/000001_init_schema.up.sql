CREATE TABLE "users" (
                         "id" SERIAL PRIMARY KEY,
                         "name" varchar NOT NULL,
                         "email" varchar UNIQUE NOT NULL,
                         "password" varchar NOT NULL,
                         "created_at" timestamptz DEFAULT (now())
);

CREATE TABLE "courses" (
                           "id" SERIAL PRIMARY KEY,
                           "title" varchar NOT NULL,
                           "description" text,
                           "price" float NOT NULL,
                           "user_id" integer REFERENCES "users" ("id") ON DELETE CASCADE,
                           "created_at" timestamptz DEFAULT (now())
);