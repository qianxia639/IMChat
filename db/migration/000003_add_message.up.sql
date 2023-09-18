CREATE TABLE messages (
    id BIGSERIAL PRIMARy KEY,
    sender_id INTEGER NOT NULL,
    receiver_id INTEGER NOT NULL,
    content VARCHAR(255) NOT NULL,
    -- type SMALLINT NOT NULL,
    sender_time TIMESTAMPTZ NOT NULL DEFAULT now(),
    receiver_time TIMESTAMPTZ NOT NULL DEFAULT '0001-01-01 00:00:00Z'
);

COMMENT ON COLUMN messages.id IS '主键Id';

COMMENT ON COLUMN messages.sender_id IS '消息发送者Id';

COMMENT ON COLUMN messages.receiver_id IS '消息接收者Id';

COMMENT ON COLUMN messages.content IS '消息内容';

-- COMMENT ON COLUMN messages.type IS '消息类型';

COMMENT ON COLUMN messages.sender_time IS '消息发送时间';

COMMENT ON COLUMN messages.receiver_time IS '消息读取时间';