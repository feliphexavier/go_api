-- migrate:up
CREATE TABLE IF NOT EXISTS pictures (
    id UUID PRIMARY KEY,
    trip_id UUID,
    url TEXT,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_trip_id_pictures FOREIGN KEY (trip_id) REFERENCES trips(id)
)

-- migrate:down

