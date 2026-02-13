--ALTER tbl_users
ALTER TABLE public.tbl_users
ALTER COLUMN is_admin SET DEFAULT 0;

ALTER TABLE public.tbl_users
ALTER COLUMN is_admin SET NOT NULL;

ALTER TABLE public.tbl_users
ALTER COLUMN is_active SET NOT NULL;

ALTER TABLE public.tbl_users
ALTER COLUMN "updated_date" DROP DEFAULT;

ALTER TABLE public.tbl_users
ALTER COLUMN "deleted_date" DROP DEFAULT;

--ALTER tbl_calorie_diary
ALTER TABLE "public"."tbl_calorie_diary"
ADD COLUMN net_calore INTEGER NOT NULL

ALTER TABLE tbl_food_log
RENAME COLUMN id_food to id_food_log

ALTER TABLE "tbl_master_food"
ALTER COLUMN food_id TYPE VARCHAR(30)

ALTER TABLE "tbl_food_log"
RENAME COLUMN logged_at to created_date