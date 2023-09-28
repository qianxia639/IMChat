CREATE TABLE "users" (
  "id" SERIAL PRIMARY KEY,
  "username" VARCHAR(20) UNIQUE NOT NULL,
  "nickname" VARCHAR(60) UNIQUE NOT NULL,
  "password" VARCHAR(100) NOT NULL,
  "email" VARCHAR(64) UNIQUE NOT NULL,
  "gender" SMALLINT NOT NULL DEFAULT 0,
  "profile_picture_url" VARCHAR(255) NOT NULL DEFAULT 'default.jpg',
  "status" SMALLINT NOT NULL DEFAULT 0,
  "password_changed_at" TIMESTAMPTZ NOT NULL DEFAULT '0001-01-01 00:00:00Z',
  "last_login_at" TIMESTAMPTZ NOT NULL DEFAULT '0001-01-01 00:00:00Z',
  "is_active" BOOLEAN DEFAULT false,
  "created_at" TIMESTAMPTZ NOT NULL DEFAULT (now()),
  "updated_at" TIMESTAMPTZ NOT NULL DEFAULT '0001-01-01 00:00:00Z'
);

COMMENT ON COLUMN "users"."id" IS '用户Id';

COMMENT ON COLUMN "users"."username" IS '用户名';

COMMENT ON COLUMN "users"."nickname" IS '用户昵称';

COMMENT ON COLUMN "users"."password" IS '密码';

COMMENT ON COLUMN "users"."email" IS '邮箱';

COMMENT ON COLUMN "users"."gender" IS '性别, 0 未知, 1 男, 2 女';

COMMENT ON COLUMN "users"."profile_picture_url" IS '个人头像图片';

COMMENT ON COLUMN "users"."status" IS '在线状态(在线/离线)';

COMMENT ON COLUMN "users"."password_changed_at" IS '密码更新时间';

COMMENT ON COLUMN "users"."last_login_at" IS '最后在线时间';

COMMENT ON COLUMN "users"."is_active" IS '活动状态(激活/未激活)';

COMMENT ON COLUMN "users"."created_at" IS '创建时间';

COMMENT ON COLUMN "users"."updated_at" IS '更新时间';


CREATE TABLE "user_login_logs" (
  "id" BIGSERIAL PRIMARY KEY,
  "user_id" INTEGER NOT NULL,
  "login_time" TIMESTAMPTZ NOT NULL DEFAULT (now()),
  "login_ip" VARCHAR(20) NOT NULL,
  "login_ip_region" VARCHAR(20) NOT NULL,
  "is_login_exceptional" BOOLEAN NOT NULL DEFAULT false,
  "platform" VARCHAR(16) NOT NULL,
  "user_agent" VARCHAR(150) NOT NULL
);

COMMENT ON COLUMN "user_login_logs"."id" IS '日志Id';

COMMENT ON COLUMN "user_login_logs"."user_id" IS '用户id';

COMMENT ON COLUMN "user_login_logs"."login_time" IS '登录时间';

COMMENT ON COLUMN "user_login_logs"."login_ip" IS '登录ip';

COMMENT ON COLUMN "user_login_logs"."login_ip_region" IS '登录ip所在地';

COMMENT ON COLUMN "user_login_logs"."is_login_exceptional" IS '是否登录异常, t: 异常, f: 非异常';

COMMENT ON COLUMN "user_login_logs"."platform" IS '登录平台';

COMMENT ON COLUMN "user_login_logs"."user_agent" IS 'user_agent';