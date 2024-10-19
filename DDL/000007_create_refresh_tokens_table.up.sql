CREATE TABLE refresh_tokens (
    refresh_token UUID NOT NULL,
    user_id INT NOT NULL,
    expires_at TIMESTAMP NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (refresh_token),
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

COMMENT ON TABLE refresh_tokens IS 'ユーザーのリフレッシュトークンを管理するテーブル';

COMMENT ON COLUMN refresh_tokens.user_id IS 'リフレッシュトークンに関連付けられたユーザーID';
COMMENT ON COLUMN refresh_tokens.refresh_token IS 'リフレッシュトークンを一意に識別するUUID';
COMMENT ON COLUMN refresh_tokens.expires_at IS 'リフレッシュトークンの有効期限';
COMMENT ON COLUMN refresh_tokens.created_at IS 'リフレッシュトークンの作成日時';
