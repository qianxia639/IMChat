-- 好友表
CREATE TABLE friends (
    user_id INTEGER NOT NULL,
    friend_id INTEGER NOT NULL,
    note VARCHAR(60) NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT (now()),
    PRIMARY KEY (user_id, friend_id),
    FOREIGN KEY (user_id) REFERENCES users ON DELETE CASCADE,
    FOREIGN KEY (friend_id) REFERENCES users ON DELETE CASCADE
);

COMMENT ON COLUMN friends.user_id IS '用户Id';

COMMENT ON COLUMN friends.friend_id IS '好友Id';

COMMENT ON COLUMN friends.note IS '好友备注';

COMMENT ON COLUMN friends.created_at IS '创建时间';

-- 群组表
CREATE TABLE groups (
    id BIGSERIAL PRIMARY KEY,
    creator_id INTEGER NOT NULL,
    group_name VARCHAR(100) NOT NULL UNIQUE,
    current_quantity INTEGER NOT NULL DEFAULT 1,
    created_at TIMESTAMPTZ NOT NULL DEFAULT (now()),
    FOREIGN KEY (creator_id) REFERENCES users ON DELETE CASCADE
);

COMMENT ON COLUMN groups.id IS '主键Id';

COMMENT ON COLUMN groups.creator_id IS '创建者Id';

COMMENT ON COLUMN groups.group_name IS '群组名';

COMMENT ON COLUMN groups.current_quantity IS '群员人数';

COMMENT ON COLUMN groups.created_at IS '创建时间';

-- 好友/群组申请表
-- CREATE TABLE friend_group_applys (
--     id BIGSERIAL PRIMARY KEY,
--     sender_id INTEGER NOT NULL,
--     receiver_id INTEGER NOT NULL,
--     apply_desc VARCHAR(30) NOT NULL,
--     status SMALLINT NOT NULL DEFAULT 0,
--     apply_type SMALLINT NOT NULL,
--     apply_time TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
--     reply_time TIMESTAMPTZ NOT NULL DEFAULT '0001-01-01 00:00:00Z'
-- );

-- COMMENT ON COLUMN friend_group_applys.id IS '主键Id';

-- COMMENT ON COLUMN friend_group_applys.sender_id IS '申请者的用户Id';

-- COMMENT ON COLUMN friend_group_applys.receiver_id IS '接收者的用户Id或群组Id';

-- COMMENT ON COLUMN friend_group_applys.apply_desc IS '申请描述';

-- COMMENT ON COLUMN friend_group_applys.status IS '申请状态, 0: 待确定, 1: 同意, 2: 拒绝';

-- COMMENT ON COLUMN friend_group_applys.apply_type IS '申请类型, 0: 好友, 1: 群组';

-- COMMENT ON COLUMN friend_group_applys.apply_time IS '申请时间';

-- COMMENT ON COLUMN friend_group_applys.reply_time IS '响应时间';