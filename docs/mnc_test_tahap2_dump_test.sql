-- -------------------------------------------------------------
-- TablePlus 6.1.8(574)
--
-- https://tableplus.com/
--
-- Database: mnc_test_tahap2
-- Generation Time: 2024-11-23 7:29:45.3000 PM
-- -------------------------------------------------------------


DROP TABLE IF EXISTS "public"."payments";
-- This script only contains the table creation statements and does not fully represent the table in the database. Do not use it as a backup.

-- Table Definition
CREATE TABLE "public"."payments" (
    "payment_id" uuid NOT NULL,
    "user_id" uuid NOT NULL,
    "amount" numeric(18,2) NOT NULL,
    "remarks" text,
    "created_date" timestamptz,
    "updated_date" timestamptz,
    PRIMARY KEY ("payment_id")
);

DROP TABLE IF EXISTS "public"."topups";
-- This script only contains the table creation statements and does not fully represent the table in the database. Do not use it as a backup.

-- Table Definition
CREATE TABLE "public"."topups" (
    "top_up_id" uuid NOT NULL,
    "user_id" uuid NOT NULL,
    "amount" numeric(18,2) NOT NULL,
    "created_date" timestamptz,
    "updated_date" timestamptz,
    PRIMARY KEY ("top_up_id")
);

DROP TABLE IF EXISTS "public"."transactions";
-- This script only contains the table creation statements and does not fully represent the table in the database. Do not use it as a backup.

DROP TYPE IF EXISTS "public"."transaction_type";
CREATE TYPE "public"."transaction_type" AS ENUM ('credit', 'debit');
DROP TYPE IF EXISTS "public"."transaction_status";
CREATE TYPE "public"."transaction_status" AS ENUM ('pending', 'success', 'failed');

-- Table Definition
CREATE TABLE "public"."transactions" (
    "transaction_id" uuid NOT NULL,
    "user_id" uuid NOT NULL,
    "type" "public"."transaction_type" NOT NULL,
    "reference_type" uuid NOT NULL,
    "reference_id" uuid NOT NULL,
    "amount" numeric(18,2) NOT NULL,
    "remarks" text,
    "balance_before" numeric(18,2) NOT NULL,
    "balance_after" numeric(18,2) NOT NULL,
    "status" "public"."transaction_status" NOT NULL,
    "created_date" timestamptz,
    "updated_date" timestamptz,
    PRIMARY KEY ("transaction_id")
);

DROP TABLE IF EXISTS "public"."transfers";
-- This script only contains the table creation statements and does not fully represent the table in the database. Do not use it as a backup.

DROP TYPE IF EXISTS "public"."transfer_status";
CREATE TYPE "public"."transfer_status" AS ENUM ('pending', 'success', 'failed');

-- Table Definition
CREATE TABLE "public"."transfers" (
    "transfer_id" uuid NOT NULL,
    "from_user_id" uuid NOT NULL,
    "to_user_id" uuid NOT NULL,
    "amount" numeric(18,2) NOT NULL,
    "remarks" text,
    "status" "public"."transfer_status" NOT NULL,
    "created_date" timestamptz,
    "updated_date" timestamptz,
    PRIMARY KEY ("transfer_id")
);

DROP TABLE IF EXISTS "public"."users";
-- This script only contains the table creation statements and does not fully represent the table in the database. Do not use it as a backup.

-- Table Definition
CREATE TABLE "public"."users" (
    "user_id" uuid NOT NULL,
    "first_name" varchar(100) NOT NULL,
    "last_name" varchar(100) NOT NULL,
    "phone_number" varchar(16) NOT NULL,
    "address" text,
    "pin" varchar NOT NULL,
    "balance" numeric(18,2) DEFAULT 0.000000,
    "created_date" timestamptz,
    "updated_date" timestamptz,
    PRIMARY KEY ("user_id")
);

INSERT INTO "public"."payments" ("payment_id", "user_id", "amount", "remarks", "created_date", "updated_date") VALUES
('2d2d70e2-12cb-46d0-a6c7-1ab8081bd4ce', 'a899ce8f-af0b-40c0-b636-c4aaeecef63e', 10000.00, 'Pulsa Telkomsel 10k', '2024-11-23 15:22:33.852141+07', '2024-11-23 15:22:33.852141+07'),
('8b8f3b1f-ca0b-4757-a7fe-2bc2312d98e2', 'a899ce8f-af0b-40c0-b636-c4aaeecef63e', 50000.00, 'Pulsa Telkomsel 50k', '2024-11-23 15:20:17.709959+07', '2024-11-23 15:20:17.70996+07');

INSERT INTO "public"."topups" ("top_up_id", "user_id", "amount", "created_date", "updated_date") VALUES
('050f29df-7652-488a-b0ab-817d549c1714', 'a899ce8f-af0b-40c0-b636-c4aaeecef63e', 100000.00, '2024-11-23 15:48:11.385893+07', '2024-11-23 15:48:11.385893+07'),
('39a3fe00-7430-4f70-9e00-86214c314ac7', 'a899ce8f-af0b-40c0-b636-c4aaeecef63e', 50000.00, '2024-11-23 14:50:57.544456+07', NULL),
('5982f3d1-970c-4bc9-b8e5-660f799bac65', 'a899ce8f-af0b-40c0-b636-c4aaeecef63e', 100000.00, '2024-11-23 18:41:23.833692+07', '2024-11-23 18:41:23.833692+07'),
('8c0f91d3-7bc3-4514-aa4e-6c375c30e917', 'a899ce8f-af0b-40c0-b636-c4aaeecef63e', 50000.00, '2024-11-23 15:06:25.707761+07', NULL);

INSERT INTO "public"."transactions" ("transaction_id", "user_id", "type", "reference_type", "reference_id", "amount", "remarks", "balance_before", "balance_after", "status", "created_date", "updated_date") VALUES
('4d4d0cb5-1e16-41f3-85e6-6bddbd6a2d28', '907ed676-8acb-4b83-91ac-1fc4023725b7', 'credit', '8c80b231-c806-42ae-ac8a-38b042dcec8a', 'c808bd92-34aa-4768-9701-985f4f6e0208', 20000.00, 'Hadiah Ultah', 0.00, 20000.00, 'success', '2024-11-23 15:45:05.646104+07', '2024-11-23 15:45:05.646105+07'),
('530e581e-0631-4cf7-8bf9-aac35790dd58', 'a899ce8f-af0b-40c0-b636-c4aaeecef63e', 'credit', '8c80b231-c806-42ae-ac8a-38b042dcec8c', '5982f3d1-970c-4bc9-b8e5-660f799bac65', 100000.00, '', 65000.00, 165000.00, 'success', '2024-11-23 18:41:27.542635+07', '2024-11-23 18:41:27.542635+07'),
('8006458b-7f8a-4374-8538-41d93ceff969', 'a899ce8f-af0b-40c0-b636-c4aaeecef63e', 'debit', '8c80b231-c806-42ae-ac8a-38b042dcec8a', '7930021f-879d-47da-8f6a-ce630d26179b', 15000.00, 'Hadiah Pernikahan 2', 95000.00, 80000.00, 'success', '2024-11-23 17:12:42.717328+07', '2024-11-23 17:12:42.717328+07'),
('9a9f386c-e910-48de-9cb2-4aa9a0e694f7', 'a899ce8f-af0b-40c0-b636-c4aaeecef63e', 'debit', '8c80b231-c806-42ae-ac8a-38b042dcec8b', '8b8f3b1f-ca0b-4757-a7fe-2bc2312d98e2', 50000.00, 'Pulsa Telkomsel 50k', 100000.00, 50000.00, 'success', '2024-11-23 15:20:22.113455+07', '2024-11-23 15:20:22.113455+07'),
('a06d9f41-716b-4c7e-a311-5e81582749bd', 'a899ce8f-af0b-40c0-b636-c4aaeecef63e', 'credit', '8c80b231-c806-42ae-ac8a-38b042dcec8c', '39a3fe00-7430-4f70-9e00-86214c314ac7', 50000.00, '', 0.00, 50000.00, 'success', '2024-11-23 14:50:57.546448+07', '2024-11-23 14:50:57.546448+07'),
('a39a53b5-4560-4fdd-8e7f-45750214bfde', 'a899ce8f-af0b-40c0-b636-c4aaeecef63e', 'debit', '8c80b231-c806-42ae-ac8a-38b042dcec8a', '7e2dc5d5-5c17-4de0-a14a-49fe29d277c0', 10000.00, 'Hadiah Pernikahan', 120000.00, 110000.00, 'success', '2024-11-23 15:48:59.978131+07', '2024-11-23 15:48:59.978132+07'),
('ad05c921-c2d2-4160-b563-ddd260388f60', 'a899ce8f-af0b-40c0-b636-c4aaeecef63e', 'credit', '8c80b231-c806-42ae-ac8a-38b042dcec8c', '050f29df-7652-488a-b0ab-817d549c1714', 100000.00, '', 20000.00, 120000.00, 'success', '2024-11-23 15:48:11.402539+07', '2024-11-23 15:48:11.402539+07'),
('b85041f4-ff44-4f0b-b580-b130207a00c1', '907ed676-8acb-4b83-91ac-1fc4023725b7', 'credit', '8c80b231-c806-42ae-ac8a-38b042dcec8a', 'c849e0ff-11b7-4a4c-a3a0-c2cb533e84fb', 15000.00, 'Hadiah Pernikahan', 30000.00, 45000.00, 'success', '2024-11-23 15:50:06.275634+07', '2024-11-23 15:50:06.275635+07'),
('b8cc5906-0cde-46ec-a5da-9206e11806ab', '907ed676-8acb-4b83-91ac-1fc4023725b7', 'credit', '8c80b231-c806-42ae-ac8a-38b042dcec8a', 'c067cfb7-559a-4d24-bf76-542fb203bca2', 15000.00, 'Hadiah Pernikahan 2', 60000.00, 75000.00, 'success', '2024-11-23 17:13:57.075912+07', '2024-11-23 17:13:57.075912+07'),
('bfa2e894-d805-46a6-8f06-b2434e2280ae', 'a899ce8f-af0b-40c0-b636-c4aaeecef63e', 'debit', '8c80b231-c806-42ae-ac8a-38b042dcec8a', 'c849e0ff-11b7-4a4c-a3a0-c2cb533e84fb', 15000.00, 'Hadiah Pernikahan', 110000.00, 95000.00, 'success', '2024-11-23 15:50:03.536982+07', '2024-11-23 15:50:03.536983+07'),
('c4f9d6aa-c765-4fef-80ee-bcda5f4375de', 'a899ce8f-af0b-40c0-b636-c4aaeecef63e', 'debit', '8c80b231-c806-42ae-ac8a-38b042dcec8a', 'c808bd92-34aa-4768-9701-985f4f6e0208', 20000.00, 'Hadiah Ultah', 40000.00, 20000.00, 'success', '2024-11-23 15:44:10.420969+07', '2024-11-23 15:44:10.420969+07'),
('d1e46c02-c61a-4f3d-b6cd-02e3df1b2d40', '907ed676-8acb-4b83-91ac-1fc4023725b7', 'credit', '8c80b231-c806-42ae-ac8a-38b042dcec8a', '7930021f-879d-47da-8f6a-ce630d26179b', 15000.00, 'Hadiah Pernikahan 2', 45000.00, 60000.00, 'success', '2024-11-23 17:12:42.720624+07', '2024-11-23 17:12:42.720624+07'),
('d7c186fd-2527-49a7-84e0-49ff9d145653', 'a899ce8f-af0b-40c0-b636-c4aaeecef63e', 'debit', '8c80b231-c806-42ae-ac8a-38b042dcec8b', '2d2d70e2-12cb-46d0-a6c7-1ab8081bd4ce', 10000.00, 'Pulsa Telkomsel 10k', 50000.00, 40000.00, 'success', '2024-11-23 15:22:35.358322+07', '2024-11-23 15:22:35.358322+07'),
('ded9d9c0-e2e7-4d74-8f86-73e1dd1922be', '907ed676-8acb-4b83-91ac-1fc4023725b7', 'credit', '8c80b231-c806-42ae-ac8a-38b042dcec8a', '7e2dc5d5-5c17-4de0-a14a-49fe29d277c0', 10000.00, 'Hadiah Pernikahan', 20000.00, 30000.00, 'success', '2024-11-23 15:49:04.415583+07', '2024-11-23 15:49:04.415583+07'),
('ec5a21b3-0200-448b-a368-241c4caffcb2', 'a899ce8f-af0b-40c0-b636-c4aaeecef63e', 'credit', '8c80b231-c806-42ae-ac8a-38b042dcec8c', '8c0f91d3-7bc3-4514-aa4e-6c375c30e917', 50000.00, '', 50000.00, 100000.00, 'success', '2024-11-23 15:06:25.709878+07', '2024-11-23 15:06:25.709879+07'),
('f4ab4ffe-b8c8-47f5-bf2d-fa4b24331b47', 'a899ce8f-af0b-40c0-b636-c4aaeecef63e', 'debit', '8c80b231-c806-42ae-ac8a-38b042dcec8a', 'c067cfb7-559a-4d24-bf76-542fb203bca2', 15000.00, 'Hadiah Pernikahan 2', 80000.00, 65000.00, 'success', '2024-11-23 17:13:57.072161+07', '2024-11-23 17:13:57.072161+07');

INSERT INTO "public"."transfers" ("transfer_id", "from_user_id", "to_user_id", "amount", "remarks", "status", "created_date", "updated_date") VALUES
('7930021f-879d-47da-8f6a-ce630d26179b', 'a899ce8f-af0b-40c0-b636-c4aaeecef63e', '907ed676-8acb-4b83-91ac-1fc4023725b7', 15000.00, 'Hadiah Pernikahan 2', 'success', '2024-11-23 17:12:39.102093+07', '2024-11-23 17:12:39.102094+07'),
('7e2dc5d5-5c17-4de0-a14a-49fe29d277c0', 'a899ce8f-af0b-40c0-b636-c4aaeecef63e', '907ed676-8acb-4b83-91ac-1fc4023725b7', 10000.00, 'Hadiah Pernikahan', 'success', '2024-11-23 15:48:58.031956+07', '2024-11-23 15:48:58.031957+07'),
('c067cfb7-559a-4d24-bf76-542fb203bca2', 'a899ce8f-af0b-40c0-b636-c4aaeecef63e', '907ed676-8acb-4b83-91ac-1fc4023725b7', 15000.00, 'Hadiah Pernikahan 2', 'success', '2024-11-23 17:13:57.067755+07', '2024-11-23 17:13:57.067755+07'),
('c808bd92-34aa-4768-9701-985f4f6e0208', 'a899ce8f-af0b-40c0-b636-c4aaeecef63e', '907ed676-8acb-4b83-91ac-1fc4023725b7', 20000.00, 'Hadiah Ultah', 'success', '2024-11-23 15:43:53.757021+07', '2024-11-23 15:43:53.757021+07'),
('c849e0ff-11b7-4a4c-a3a0-c2cb533e84fb', 'a899ce8f-af0b-40c0-b636-c4aaeecef63e', '907ed676-8acb-4b83-91ac-1fc4023725b7', 15000.00, 'Hadiah Pernikahan', 'success', '2024-11-23 15:50:02.646292+07', '2024-11-23 15:50:02.646292+07');

INSERT INTO "public"."users" ("user_id", "first_name", "last_name", "phone_number", "address", "pin", "balance", "created_date", "updated_date") VALUES
('907ed676-8acb-4b83-91ac-1fc4023725b7', 'Abdul', 'Hanif', '62813128489442', 'Jln. Merdeka, No. 60, jakarta', '$2a$08$iBT0Z0Y./j1Linng4GXOROI3B5eq9WPpOl4IyyGeH.4ZbccdVhuz.', 75000.00, '2024-11-23 15:42:26.161544+07', '2024-11-23 17:13:57.076293+07'),
('a899ce8f-af0b-40c0-b636-c4aaeecef63e', 'Admin', 'MNC', '62813128489441', 'Jln. Merdeka, No. 59, jakarta', '$2a$08$2GaRTGTfLc9eBXC8nHN70O2ubEPFJTdu1XkkXvC73GYOXxAOCZX/u', 165000.00, '2024-11-23 14:13:26.4155+07', '2024-11-23 18:41:27.543572+07');



-- Indices
CREATE UNIQUE INDEX users_phone_number_key ON public.users USING btree (phone_number);
