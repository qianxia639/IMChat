-- 群组成员关系表
CREATE TABLE group_member_ships (
    id BIGSERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL REFERENCES users ON DELETE CASCADE,
    group_id INTEGER NOT NULL REFERENCES groups ON DELETE CASCADE,
    role INTEGER NOT NULL DEFAULT 0,
    joined_at TIMESTAMPTZ NOT NULL DEFAULT (now())
);

COMMENT ON COLUMN group_member_ships.id IS '群组成员的唯一标识符';

COMMENT ON COLUMN group_member_ships.user_id IS '用户Id';

COMMENT ON COLUMN group_member_ships.group_id IS '群组Id';

COMMENT ON COLUMN group_member_ships.role IS '角色, 0: 管理员, 1: 普通成员';

COMMENT ON COLUMN group_member_ships.joined_at IS '加入时间';