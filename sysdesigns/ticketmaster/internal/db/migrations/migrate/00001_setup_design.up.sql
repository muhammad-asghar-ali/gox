CREATE OR REPLACE FUNCTION generate_custom_uuid()
RETURNS UUID AS $$
BEGIN
  RETURN gen_random_uuid();
END;
$$ LANGUAGE plpgsql;

CREATE TABLE IF NOT EXISTS users (
  id UUID PRIMARY KEY DEFAULT generate_custom_uuid(),
  name VARCHAR(255) NOT NULL,
  email VARCHAR(255) NOT NULL UNIQUE,
  password VARCHAR(255) NOT NULL,
  phone VARCHAR(255),
  created_at TIMESTAMP DEFAULT NOW() NOT NULL,
  last_login TIMESTAMP DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS venues (
  id UUID PRIMARY KEY DEFAULT generate_custom_uuid(),
  name VARCHAR(255) NOT NULL,
  location VARCHAR(255) NOT NULL,
  capacity INT NOT NULL,
  added_by UUID NOT NULL,
  created_at TIMESTAMP DEFAULT NOW() NOT NULL,
  FOREIGN KEY (added_by) REFERENCES users(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS events (
  id UUID PRIMARY KEY DEFAULT generate_custom_uuid(),
  name VARCHAR(255) NOT NULL,
  description TEXT NOT NULL,
  added_by UUID NOT NULL,
  venue_id UUID NOT NULL,
  event_date TIMESTAMP NOT NULL,
  total_tickets INT NOT NULL,
  available_tickets INT NOT NULL,
  created_at TIMESTAMP DEFAULT NOW() NOT NULL,
  FOREIGN KEY (added_by) REFERENCES users(id) ON DELETE CASCADE,
  FOREIGN KEY (venue_id) REFERENCES venues(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS tickets (
  id UUID PRIMARY KEY DEFAULT generate_custom_uuid(),
  event_id UUID NOT NULL,
  ticket_type VARCHAR(255) NOT NULL,
  price DECIMAL(10, 2) NOT NULL,
  total_tickets INT NOT NULL,
  available_tickets INT NOT NULL,
  created_at TIMESTAMP DEFAULT NOW(),
  FOREIGN KEY (event_id) REFERENCES events(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS bookings (
  id UUID PRIMARY KEY DEFAULT generate_custom_uuid(),
  user_id UUID NOT NULL,
  ticket_id UUID NOT NULL,
  quantity INT NOT NULL,
  total_price DECIMAL(10, 2) NOT NULL,
  status VARCHAR(255) NOT NULL CHECK (status IN ('pending', 'confirmed', 'canceled')),
  created_at TIMESTAMP DEFAULT NOW() NOT NULL,
  FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
  FOREIGN KEY (ticket_id) REFERENCES tickets(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS payments (
  id UUID PRIMARY KEY DEFAULT generate_custom_uuid(),
  booking_id UUID NOT NULL,
  user_id UUID NOT NULL,
  amount DECIMAL(10, 2) NOT NULL,
  payment_method VARCHAR(255) NOT NULL CHECK (payment_method IN ('credit_card', 'paypal', 'bank_transfer', 'cash')),
  status VARCHAR(255) NOT NULL CHECK (status IN ('pending', 'completed', 'failed', 'refunded')),
  created_at TIMESTAMP DEFAULT NOW(),
  FOREIGN KEY (booking_id) REFERENCES bookings(id) ON DELETE CASCADE,
  FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS performers (
  id uuid NOT NULL PRIMARY KEY DEFAULT gen_random_uuid(),
  name varchar(500) NOT NULL,
  genre varchar(500) NOT NULL,
  bio text NOT NULL,
  created_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS event_performers (
  id uuid NOT NULL PRIMARY KEY DEFAULT gen_random_uuid(),
  event_id uuid NOT NULL REFERENCES events(id) ON DELETE CASCADE,
  performer_id uuid NOT NULL REFERENCES performers(id) ON DELETE CASCADE
);