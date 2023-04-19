CREATE TABLE "books" (
                         "id" SERIAL PRIMARY KEY,
                         "created_at" TIMESTAMP WITH TIME ZONE,
                         "updated_at" TIMESTAMP WITH TIME ZONE,
                         "deleted_at" TIMESTAMP WITH TIME ZONE,
                         "title" TEXT,
                         "author" TEXT
);


CREATE TABLE "users" (
                         "id" SERIAL PRIMARY KEY,
                         "created_at" TIMESTAMP WITH TIME ZONE,
                         "updated_at" TIMESTAMP WITH TIME ZONE,
                         "deleted_at" TIMESTAMP WITH TIME ZONE,
                         "name" TEXT,
                         "email" TEXT UNIQUE,
                         "password" BYTEA
);

CREATE TABLE "records" (
                           "id" SERIAL PRIMARY KEY,
                           "created_at" TIMESTAMP WITH TIME ZONE,
                           "updated_at" TIMESTAMP WITH TIME ZONE,
                           "deleted_at" TIMESTAMP WITH TIME ZONE,
                           "user_id" INTEGER,
                           "book_id" INTEGER,
                           "taken_at" TIMESTAMP WITH TIME ZONE,
                           "returned_at" TIMESTAMP WITH TIME ZONE,
                           "borrowed" BOOLEAN,
                           FOREIGN KEY ("user_id") REFERENCES "users" ("id"),
                           FOREIGN KEY ("book_id") REFERENCES "books" ("id")
);
