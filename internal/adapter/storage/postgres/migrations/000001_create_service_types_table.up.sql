CREATE TYPE "service_types_status_enum" AS ENUM ('active', 'inactive');

CREATE TABLE "service_types" (
    "id" varchar PRIMARY KEY,
    "name" varchar NOT NULL,
    "code" varchar NOT NULL,
    "status" service_types_status_enum DEFAULT 'active',
    "institution_id" varchar NOT NULL
);
