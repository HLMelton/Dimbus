-- Create the short_links table
CREATE TABLE IF NOT EXISTS short_links (
    id SERIAL PRIMARY KEY,
    code VARCHAR(10) NOT NULL UNIQUE,
    url TEXT NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    last_accessed TIMESTAMP WITH TIME ZONE
);

-- Create an index on the code column for faster lookups
CREATE INDEX IF NOT EXISTS idx_short_links_code ON short_links(code); 