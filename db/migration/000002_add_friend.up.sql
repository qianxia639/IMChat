CREATE TABLE friends (
    user_id INTEGER NOT NULL REFERENCES users ON DELETE CASCADE,
    friend_id INTEGER NOT NULL REFERENCES users ON DELETE CASCADE,
    status SMALLINT NOT NULL,
    note VARCHAR(20) NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT (now()),
    PRIMARY KEY (user_id, friend_id)
);

COMMENT ON COLUMN friends.user_id IS '用户Id';

COMMENT ON COLUMN friends.friend_id IS '好友Id';

COMMENT ON COLUMN friends.status IS '好友状态, 1: 在线, 2: 离线';

COMMENT ON COLUMN friends.note IS '好友备注';

COMMENT ON COLUMN friends.created_at IS '创建时间';

-- CREATE TABLE friend_apply (
--     apply_id INTEGER NOT NULL REFERENCES users ON DELETE CASCADE,
--     reply_id INTEGER NOT NULL REFERENCES users ON DELETE CASCADE,
--     apply_desc VARCHAR(20) NOT NULL,
--     note VARCHAR(20) NOT NULL,
--     created_at TIMESTAMPTZ NOT NULL DEFAULT (now()),
--     PRIMARY KEY(apply_id, reply_id)
-- );

-- COMMENT ON COLUMN friend_apply.apply_id IS '所有者Id';

-- COMMENT ON COLUMN friend_apply.reply_id IS '目标对象Id';

-- COMMENT ON COLUMN friend_apply.apply_desc IS '申请描述';

-- COMMENT ON COLUMN friend_apply.note IS '备注';

-- COMMENT ON COLUMN friend_apply.created_at IS '创建时间';

CREATE TABLE friend_cluster_apply (
    id BIGSERIAL PRIMARY KEY,
    apply_id INTEGER NOT NULL,
    receiver_id INTEGER NOT NULL,
    apply_desc VARCHAR(30) NOT NULL,
    status SMALLINT NOT NULL DEFAULT 0,
    flag SMALLINT NOT NULL,
    apply_time TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    response_time TIMESTAMPTZ NOT NULL DEFAULT '0001-01-01 00:00:00Z'
);

COMMENT ON COLUMN friend_cluster_apply.id IS '主键Id';

COMMENT ON COLUMN friend_cluster_apply.apply_id IS '申请者的用户Id';

COMMENT ON COLUMN friend_cluster_apply.receiver_id IS '接收者的用户或群组Id';

COMMENT ON COLUMN friend_cluster_apply.apply_desc IS '申请描述';

COMMENT ON COLUMN friend_cluster_apply.status IS '申请状态, 0: 等待中, 1: 同意, 2: 拒绝';

COMMENT ON COLUMN friend_cluster_apply.flag IS '申请标识, 0: 好友, 1: 群组';

COMMENT ON COLUMN friend_cluster_apply.apply_time IS '申请时间';

COMMENT ON COLUMN friend_cluster_apply.response_time IS '响应时间';