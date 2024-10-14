CREATE TABLE sessions (
    session_id UUID NOT NULL,
    user_id INT NOT NULL UNIQUE,
    expires_at TIMESTAMP NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (session_id),
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

COMMENT ON TABLE sessions IS 'ユーザーセッションを管理するテーブル';

COMMENT ON COLUMN sessions.user_id IS 'セッションに関連付けられたユーザーID';
COMMENT ON COLUMN sessions.session_id IS 'セッションを一意に識別するUUID';
COMMENT ON COLUMN sessions.expires_at IS 'セッションの有効期限';
COMMENT ON COLUMN sessions.created_at IS 'セッションの作成日時';
