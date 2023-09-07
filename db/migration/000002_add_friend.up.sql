CREATE TABLE friends (
    user_id INTEGER NOT NULL REFERENCES users ON DELETE CASCADE,
    friend_id INTEGER NOT NULL REFERENCES users ON DELETE CASCADE,
    note VARCHAR(20) NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT (now()),
    PRIMARY KEY (user_id, friend_id)
);

COMMENT ON COLUMN friends.user_id IS '用户Id';

COMMENT ON COLUMN friends.friend_id IS '好友Id';

COMMENT ON COLUMN friends.note IS '好友备注';

COMMENT ON COLUMN friends.created_at IS '创建时间';

CREATE TABLE friend_apply (
    apply_id INTEGER NOT NULL REFERENCES users ON DELETE CASCADE,
    reply_id INTEGER NOT NULL REFERENCES users ON DELETE CASCADE,
    apply_desc VARCHAR(20) NOT NULL,
    note VARCHAR(20) NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT (now()),
    PRIMARY KEY(apply_id, reply_id)
);

COMMENT ON COLUMN friend_apply.apply_id IS '所有者Id';

COMMENT ON COLUMN friend_apply.reply_id IS '目标对象Id';

COMMENT ON COLUMN friend_apply.apply_desc IS '申请描述';

COMMENT ON COLUMN friend_apply.note IS '备注';

COMMENT ON COLUMN friend_apply.created_at IS '创建时间';