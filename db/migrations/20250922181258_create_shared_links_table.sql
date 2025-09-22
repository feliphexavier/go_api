-- migrate:up
CREATE TABLE IF NOT EXISTS shared_links (
    id UUID PRIMARY KEY,
    trip_id UUID,
    shared_links_token TEXT,
    expired_at TIMESTAMP,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_user_id_refresh_tokens FOREIGN KEY (trip_id) REFERENCES trips(id)
)

-- migrate:down
DROP TABLE IF EXISTS shared_links
