CREATE TABLE password_reset_tokens (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL UNIQUE,
    token UUID NOT NULL,
    expires_at TIMESTAMP NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

COMMENT ON TABLE password_reset_tokens IS 'パスワードリセットのトークンを管理するテーブル';

COMMENT ON COLUMN password_reset_tokens.id IS 'トークンの一意のID（プライマリキー）';
COMMENT ON COLUMN password_reset_tokens.user_id IS 'パスワードリセットを要求したユーザーのID（usersテーブルの外部キー）';
COMMENT ON COLUMN password_reset_tokens.token IS 'パスワードリセット用の一意なトークン';
COMMENT ON COLUMN password_reset_tokens.expires_at IS 'トークンの有効期限。これを過ぎると無効となる';
COMMENT ON COLUMN password_reset_tokens.created_at IS 'トークンの生成日時。作成時に自動的に設定される';
