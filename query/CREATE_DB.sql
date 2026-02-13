CREATE TABLE "public"."tbl_users" ( 
  "id_user" SERIAL,
  "is_admin" SMALLINT NULL, -- hrs di alter
  "username" VARCHAR(30) NOT NULL,
  "password" VARCHAR(30) NOT NULL,
  "height_cm" DOUBLE PRECISION NULL,
  "weight kg" DOUBLE PRECISION NULL,
  "bmi" DOUBLE PRECISION NULL,
  "gender" VARCHAR(2) NULL,
  "dob" DATE NULL,
  "age" INTEGER NULL,
  "created_date" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ,
  "created_by" VARCHAR(30) NOT NULL,
  "updated_date" TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP ,
  "updated_by" VARCHAR(30) NULL,
  "deleted_date" TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP ,
  "deleted_by" VARCHAR(30) NULL,
  "is_active" SMALLINT NULL DEFAULT 1 ,
  CONSTRAINT "tbl_users_pkey" PRIMARY KEY ("id_user")
);

CREATE TABLE "public"."tbl_calorie_diary" ( 
  "id_diary" SERIAL,
  "id_user" INTEGER NOT NULL,
  "total_calories_in" INTEGER NOT NULL,
  "total_calories_out" INTEGER NOT NULL,
  "diary_date" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ,
  "is_active" INTEGER NOT NULL DEFAULT 1 ,
  CONSTRAINT "tbl_calorie_diary_pkey" PRIMARY KEY ("id_diary")
);

CREATE TABLE "public"."tbl_activities_log" ( 
  "id_activities" SERIAL,
  "id_sports" INTEGER NOT NULL,
  "id_user" INTEGER NOT NULL,
  "id_diary" INTEGER NOT NULL,
  "length_minutes" INTEGER NOT NULL,
  "calories_burned" INTEGER NOT NULL,
  "logged_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ,
  "is_active" SMALLINT NOT NULL DEFAULT 1 ,
  CONSTRAINT "tbl_activities_log_pkey" PRIMARY KEY ("id_activities")
);

CREATE TABLE "public"."tbl_food_log" ( 
  "id_food" SERIAL,
  "id_user" INTEGER NOT NULL,
  "id_diary" INTEGER NOT NULL,
  "food_name" VARCHAR(100) NOT NULL,
  "protein_grams" DOUBLE PRECISION NOT NULL,
  "carbohydrate_grams" DOUBLE PRECISION NOT NULL,
  "fat_grams" DOUBLE PRECISION NOT NULL,
  "weight_grams" DOUBLE PRECISION NOT NULL,
  "logged_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ,
  "is_active" SMALLINT NOT NULL DEFAULT 1 ,
  "calories" INTEGER NOT NULL,
  CONSTRAINT "tbl_food_log_pkey" PRIMARY KEY ("id_food")
);

CREATE TABLE "public"."tbl_sports" ( 
  "id_sports" SERIAL,
  "sports_name" VARCHAR(50) NOT NULL,
  "calories_burned_hourly" INTEGER NOT NULL,
  "created_date" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ,
  "created_by" VARCHAR(30) NOT NULL,
  "updated_date" TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP ,
  "updated_by" VARCHAR(30) NULL,
  "deleted_date" TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP ,
  "deleted_by" VARCHAR(30) NULL,
  "is_active" SMALLINT NOT NULL DEFAULT 1 ,
  CONSTRAINT "tbl_sports_pkey" PRIMARY KEY ("id_sports")
);
