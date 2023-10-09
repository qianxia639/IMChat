-- 消息表
CREATE TABLE messages (
    id BIGSERIAL PRIMARy KEY,
    sender_id INTEGER NOT NULL,
    receiver_id INTEGER NOT NULL,
    message_type SMALLINT NOT NULL,
    content VARCHAR(255) NOT NULL,
    send_type SMALLINT NOT NULl,
    sending_time TIMESTAMPTZ NOT NULL DEFAULT now(),
    receiv_time TIMESTAMPTZ NOT NULL DEFAULT '0001-01-01 00:00:00Z'
);

COMMENT ON COLUMN messages.id IS '主键Id';

COMMENT ON COLUMN messages.sender_id IS '消息发送者Id';

COMMENT ON COLUMN messages.receiver_id IS '消息接收者Id';

COMMENT ON COLUMN messages.message_type IS '消息类型, 0: 文字';

COMMENT ON COLUMN messages.content IS '消息内容';

COMMENT ON COLUMN messages.send_type IS '发送类型, 0: 私聊, 1: 群聊';

COMMENT ON COLUMN messages.sending_time IS '消息发送时间';

COMMENT ON COLUMN messages.receiv_time IS '消息读取时间';