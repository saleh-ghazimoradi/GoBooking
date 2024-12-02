CREATE TABLE IF NOT EXISTS events (
    id bigserial PRIMARY KEY,
    name varchar(255) NOT NULL,
    location varchar(255) NOT NULL,
    date TIMESTAMP(0) with time zone NOT NULL DEFAULT NOW(),
    created_at TIMESTAMP(0) with time zone NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP(0) with time zone NOT NULL DEFAULT NOW()
);

ALTER TABLE events ADD COLUMN version INT DEFAULT 0;

CREATE TABLE IF NOT EXISTS tickets (
    id bigserial PRIMARY KEY,
    event_id INT NOT NULL REFERENCES events(id) ON DELETE CASCADE ON UPDATE CASCADE,
    user_id INT NOT NULL,
    entered BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP(0) with time zone NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP(0) with time zone NOT NULL DEFAULT NOW()
);