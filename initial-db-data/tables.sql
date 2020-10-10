-- UUID tool
CREATE EXTENSION pgcrypto;

-- Base tables
CREATE TABLE IF NOT EXISTS products (
   id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
   price_in_cents BIGINT,
   title VARCHAR(255),
   description TEXT
);

CREATE TABLE IF NOT EXISTS users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    first_name VARCHAR(50),
    last_name VARCHAR(50),
    date_of_birth DATE
);
