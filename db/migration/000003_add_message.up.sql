-- 消息表
CREATE TABLE messages (
    id BIGSERIAL PRIMARy KEY,
    sender_id INTEGER NOT NULL,
    receiver_id INTEGER NOT NULL,
    message_type SMALLINT NOT NULL,
    content VARCHAR(255) NOT NULL,
    send_type SMALLINT NOT NULl,
    -- message_status SMALLINT NOT NUll,
    pic VARCHAR(255) NOT NULL,
    url VARCHAR(255) NOT NULL,
    sending_time TIMESTAMPTZ NOT NULL DEFAULT now(),
    receive_time TIMESTAMPTZ NOT NULL DEFAULT '0001-01-01 00:00:00Z'
);

COMMENT ON COLUMN messages.id IS '主键Id';

COMMENT ON COLUMN messages.sender_id IS '消息发送者Id';

COMMENT ON COLUMN messages.receiver_id IS '消息接收者Id';

COMMENT ON COLUMN messages.message_type IS '消息类型, 1: 文字, 2: 图片, 3: 文件, 4: 语音, 5: 视频, 6: 语音通话, 7: 视频通话';

COMMENT ON COLUMN messages.content IS '消息内容';

COMMENT ON COLUMN messages.send_type IS '发送类型, 1: 私聊, 2: 群聊';

-- COMMENT ON COLUMN messages.message_status IS '消息状态, 1: 发送失败, 2: 发送成功, 3: 已读, 4: 未读';

COMMENT ON COLUMN messages.pic IS '缩略图';

COMMENT ON COLUMN messages.url IS '文件或图片地址';

COMMENT ON COLUMN messages.sending_time IS '消息发送时间';

COMMENT ON COLUMN messages.receive_time IS '消息读取时间';

CREATE INDEX messages_sender_id_index ON messages (sender_id);

CREATE INDEX messages_receiver_id_index ON messages (receiver_id);