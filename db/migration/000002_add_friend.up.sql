-- 好友表
CREATE TABLE friendships (
    id BIGSERIAL PRIMARY KEy,
    user_id INTEGER NOT NULL,
    friend_id INTEGER NOT NULL,
    comment VARCHAR(60) NOT NULL,
    status SMALLINT NOT NULL DEFAULT 1,
    created_at TIMESTAMPTZ NOT NULL DEFAULT (now()),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT '0001-01-01 00:00:00Z',
    FOREIGN KEY (user_id) REFERENCES users ON DELETE CASCADE,
    FOREIGN KEY (friend_id) REFERENCES users ON DELETE CASCADE
);

COMMENT ON COLUMN friendships.id IS '主键Id';

COMMENT ON COLUMN friendships.user_id IS '用户Id';

COMMENT ON COLUMN friendships.friend_id IS '好友Id';

COMMENT ON COLUMN friendships.comment IS '好友备注';

COMMENT ON COLUMN friendships.status IS '好友状态, 1: 待确认, 2: 已确认, 3: 已拒绝';

COMMENT ON COLUMN friendships.created_at IS '创建时间';

COMMENT ON COLUMN friendships.updated_at IS '更新时间';

-- 群组表
CREATE TABLE groups (
    id BIGSERIAL PRIMARY KEY,
    creator_id INTEGER NOT NULL,
    group_name VARCHAR(100) NOT NULL UNIQUE,
    group_member_quantity INTEGER NOT NULL DEFAULT 1,
    description VARCHAR(255) NOT NULL DEFAULT '',
    created_at TIMESTAMPTZ NOT NULL DEFAULT (now()),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT '0001-01-01 00:00:00Z',
    FOREIGN KEY (creator_id) REFERENCES users ON DELETE CASCADE
);

COMMENT ON COLUMN groups.id IS '主键Id';

COMMENT ON COLUMN groups.creator_id IS '创建者Id';

COMMENT ON COLUMN groups.group_name IS '群组名';

COMMENT ON COLUMN groups.group_member_quantity IS '群员人数';

COMMENT ON COLUMN groups.description IS '群组描述';

COMMENT ON COLUMN groups.created_at IS '创建时间';

COMMENT ON COLUMN groups.updated_at IS '更新时间';

-- 群组成员关系表
CREATE TABLE group_member_ships (
    id BIGSERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL,
    group_id INTEGER NOT NULL,
    role SMALLINT NOT NULL DEFAULT 3,
    status SMALLINT NOT NULL DEFAULT 1,
    created_at TIMESTAMPTZ NOT NULL DEFAULT (now()),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT '0001-01-01 00:00:00Z',
    FOREIGN KEY (user_id) REFERENCES users ON DELETE CASCADE,
    FOREIGN KEY (group_id) REFERENCES groups ON DELETE CASCADE
);

COMMENT ON COLUMN group_member_ships.id IS '主键Id';

COMMENT ON COLUMN group_member_ships.user_id IS '用户Id';

COMMENT ON COLUMN group_member_ships.group_id IS '群组Id';

COMMENT ON COLUMN group_member_ships.role IS '群员角色, 1: 群主, 2: 管理员, 3: 普通成员';

COMMENT ON COLUMN group_member_ships.status IS '群员状态, 1: 待确认, 2: 已确认, 3: 已拒绝';

COMMENT ON COLUMN group_member_ships.created_at IS '创建时间';

COMMENT ON COLUMN group_member_ships.updated_at IS '更新时间';

-- 好友/群组申请记录表
-- CREATE TABLE friend_group_applys (
--     id BIGSERIAL PRIMARY KEY,
--     user_id INTEGER NOT NULL,
--     target_id INTEGER NOT NULL,
--     description VARCHAR(100) NOT NULl DEFAULT '',
--     apply_type SMALLINT NOT NULL,
--     status SMALLINT NOT NULL DEFAULT 1,
--     created_at TIMESTAMPTZ NOT NULL DEFAULT (now()),
--     updated_at TIMESTAMPTZ NOT NULL DEFAULT '0001-01-01 00:00:00Z'
-- );

-- COMMENT ON COLUMN friend_group_applys.id IS '主键Id';

-- COMMENT ON COLUMN friend_group_applys.user_id IS '申请者用户Id';

-- COMMENT ON COLUMN friend_group_applys.target_id IS '响应者Id(好友/群组Id)';

-- COMMENT ON COLUMN friend_group_applys.description IS '申请描述';

-- COMMENT ON COLUMN friend_group_applys.apply_type IS '申请类型, 1: 好友, 2: 群组';

-- COMMENT ON COLUMN friend_group_applys.status IS '申请状态, 1: 待确认, 2: 已确认, 3: 已拒绝';

-- COMMENT ON COLUMN friend_group_applys.created_at IS '创建时间';

-- COMMENT ON COLUMN friend_group_applys.updated_at IS '响应时间';