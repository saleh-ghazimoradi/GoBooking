CREATE TABLE IF NOT EXISTS events (
    id bigserial PRIMARY KEY,
    name varchar(255) NOT NULL,
    location varchar(255) NOT NULL,
    date TIMESTAMP(0) with time zone NOT NULL DEFAULT NOW(),
    created_at TIMESTAMP(0) with time zone NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP(0) with time zone NOT NULL DEFAULT NOW()
);