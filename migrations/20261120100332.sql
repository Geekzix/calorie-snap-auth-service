-- Create "permissions" table
CREATE TABLE "public"."permissions" (
  "id" uuid NOT NULL DEFAULT gen_random_uuid(),
  "code" character varying(100) NOT NULL,
  "description" text NULL,
  "created_by" uuid NULL,
  "updated_by" uuid NULL,
  "created_at" timestamptz NOT NULL DEFAULT now(),
  "updated_at" timestamptz NOT NULL DEFAULT now(),
  PRIMARY KEY ("id")
);
-- Create index "ux_permissions_code" to table: "permissions"
CREATE UNIQUE INDEX "ux_permissions_code" ON "public"."permissions" ("code");
-- Create "roles" table
CREATE TABLE "public"."roles" (
  "id" uuid NOT NULL DEFAULT gen_random_uuid(),
  "name" character varying(50) NOT NULL,
  "description" character varying(255) NULL,
  "created_by" uuid NULL,
  "updated_by" uuid NULL,
  "created_at" timestamptz NOT NULL DEFAULT now(),
  "updated_at" timestamptz NOT NULL DEFAULT now(),
  PRIMARY KEY ("id")
);
-- Create index "ux_role_name" to table: "roles"
CREATE UNIQUE INDEX "ux_role_name" ON "public"."roles" ("name");
-- Create "role_permission" table
CREATE TABLE "public"."role_permission" (
  "role_id" uuid NOT NULL DEFAULT gen_random_uuid(),
  "permission_id" uuid NOT NULL DEFAULT gen_random_uuid(),
  PRIMARY KEY ("role_id", "permission_id"),
  CONSTRAINT "fk_role_permission_permission" FOREIGN KEY ("permission_id") REFERENCES "public"."permissions" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT "fk_role_permission_role" FOREIGN KEY ("role_id") REFERENCES "public"."roles" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION
);
-- Create "users" table
CREATE TABLE "public"."users" (
  "id" uuid NOT NULL DEFAULT gen_random_uuid(),
  "username" character varying(255) NOT NULL,
  "email" character varying(100) NOT NULL,
  "phone" character varying(15) NOT NULL,
  "password" character varying(255) NOT NULL,
  "gender" smallint NOT NULL,
  "status" character varying(32) NOT NULL DEFAULT 'PENDING',
  "role_id" uuid NOT NULL,
  "created_by" uuid NULL,
  "updated_by" uuid NULL,
  "created_at" timestamptz NOT NULL DEFAULT now(),
  "updated_at" timestamptz NOT NULL DEFAULT now(),
  PRIMARY KEY ("id"),
  CONSTRAINT "fk_users_role" FOREIGN KEY ("role_id") REFERENCES "public"."roles" ("id") ON UPDATE CASCADE ON DELETE RESTRICT
);
-- Create index "idx_users_role_id" to table: "users"
CREATE INDEX "idx_users_role_id" ON "public"."users" ("role_id");
-- Create index "idx_users_status" to table: "users"
CREATE INDEX "idx_users_status" ON "public"."users" ("status");
-- Create index "ux_user_email" to table: "users"
CREATE UNIQUE INDEX "ux_user_email" ON "public"."users" ("email");
-- Create index "ux_user_phone" to table: "users"
CREATE UNIQUE INDEX "ux_user_phone" ON "public"."users" ("phone");
