CREATE SEQUENCE
IF NOT EXISTS users_id_seq;
CREATE SEQUENCE
IF NOT EXISTS events_id_seq;
CREATE SEQUENCE
IF NOT EXISTS venues_id_seq;
CREATE SEQUENCE
IF NOT EXISTS tickets_id_seq;
CREATE SEQUENCE
IF NOT EXISTS bookings_id_seq;
CREATE SEQUENCE
IF NOT EXISTS payments_id_seq;
CREATE SEQUENCE
IF NOT EXISTS seats_id_seq;
CREATE SEQUENCE
IF NOT EXISTS discounts_id_seq;
CREATE SEQUENCE
IF NOT EXISTS applied_discounts_id_seq;
CREATE SEQUENCE
IF NOT EXISTS refunds_id_seq;

CREATE TABLE
IF NOT EXISTS users
(
  id uuid NOT NULL PRIMARY KEY DEFAULT gen_random_uuid
(),
  name varchar
(500) NOT NULL,
  email varchar
(500) NOT NULL,
  password varchar
(500) NOT NULL,
  phone varchar
(500),
  created_at timestamp NOT NULL,
  last_login timestamp NOT NULL,
  role varchar
(500)
);

CREATE TABLE
IF NOT EXISTS events
(
  id uuid NOT NULL PRIMARY KEY DEFAULT gen_random_uuid
(),
  name varchar
(500) NOT NULL,
  "desc" text NOT NULL,
  organizer_id uuid NOT NULL,
  venue_id uuid NOT NULL,
  event_date timestamp NOT NULL,
  created_at timestamp NOT NULL
);

CREATE TABLE
IF NOT EXISTS venues
(
  id uuid NOT NULL PRIMARY KEY DEFAULT gen_random_uuid
(),
  name bigint NOT NULL,
  location bigint NOT NULL,
  capacity bigint NOT NULL,
  added_by uuid NOT NULL,
  created_at timestamp NOT NULL,
  seating_map json
);

CREATE TABLE
IF NOT EXISTS tickets
(
  id uuid NOT NULL PRIMARY KEY DEFAULT gen_random_uuid
(),
  event_id uuid NOT NULL,
  ticket_type varchar
(500) NOT NULL,
  price decimal NOT NULL,
  total_tickets int NOT NULL,
  available_tickets int NOT NULL,
  created_at timestamp,
  seat_id uuid NOT NULL
);

CREATE TABLE
IF NOT EXISTS bookings
(
  id uuid NOT NULL PRIMARY KEY DEFAULT gen_random_uuid
(),
  user_id uuid NOT NULL,
  ticket_id uuid NOT NULL,
  quantity int NOT NULL,
  total_price decimal NOT NULL,
  status varchar
(500) NOT NULL,
  created_at timestamp NOT NULL,
  seat_id uuid NOT NULL
);

CREATE TABLE
IF NOT EXISTS payments
(
  id uuid NOT NULL PRIMARY KEY DEFAULT gen_random_uuid
(),
  booking_id uuid NOT NULL,
  user_id uuid NOT NULL,
  amount decimal NOT NULL,
  payment_method varchar
(500) NOT NULL,
  status varchar
(500) NOT NULL,
  created_at timestamp NOT NULL,
  transaction_id varchar
(500) NOT NULL UNIQUE
);

CREATE TABLE
IF NOT EXISTS seats
(
  id uuid NOT NULL PRIMARY KEY DEFAULT gen_random_uuid
(),
  venue_id uuid NOT NULL,
  section varchar
(500),
  row_number varchar
(500),
  seat_number varchar
(500),
  is_available boolean
);

CREATE TABLE
IF NOT EXISTS discounts
(
  id uuid NOT NULL PRIMARY KEY DEFAULT gen_random_uuid
(),
  event_id uuid NOT NULL,
  code varchar
(500) NOT NULL,
  discount_type varchar
(500) NOT NULL,
  discount_value decimal NOT NULL,
  max_uses int NOT NULL,
  expires_at timestamp NOT NULL,
  created_at timestamp NOT NULL
);

CREATE TABLE
IF NOT EXISTS applied_discounts
(
  id bigint NOT NULL PRIMARY KEY,
  booking_id uuid NOT NULL,
  discount_id uuid NOT NULL,
  discount_amount decimal NOT NULL,
  created_at timestamp NOT NULL
);

CREATE TABLE
IF NOT EXISTS refunds
(
  id uuid NOT NULL PRIMARY KEY DEFAULT gen_random_uuid
(),
  booking_id uuid NOT NULL,
  user_id uuid NOT NULL,
  refund_reason text NOT NULL,
  amount decimal NOT NULL,
  status varchar
(500) NOT NULL,
  created_at timestamp NOT NULL
);

ALTER TABLE venues ADD CONSTRAINT venues_id_fk FOREIGN KEY (id) REFERENCES events (venue_id);
ALTER TABLE users ADD CONSTRAINT users_id_fk FOREIGN KEY (id) REFERENCES events (organizer_id);
ALTER TABLE users ADD CONSTRAINT users_added_by_fk FOREIGN KEY (id) REFERENCES venues (added_by);
ALTER TABLE tickets ADD CONSTRAINT tickets_event_id_fk FOREIGN KEY (event_id) REFERENCES events (id);
ALTER TABLE bookings ADD CONSTRAINT bookings_user_id_fk FOREIGN KEY (user_id) REFERENCES users (id);
ALTER TABLE bookings ADD CONSTRAINT bookings_ticket_id_fk FOREIGN KEY (ticket_id) REFERENCES tickets (id);
ALTER TABLE payments ADD CONSTRAINT payments_booking_id_fk FOREIGN KEY (booking_id) REFERENCES bookings (id);
ALTER TABLE users ADD CONSTRAINT users_payment_user_id_fk FOREIGN KEY (id) REFERENCES payments (user_id);
ALTER TABLE venues ADD CONSTRAINT venues_seat_venue_id_fk FOREIGN KEY (id) REFERENCES seats (venue_id);
ALTER TABLE tickets ADD CONSTRAINT tickets_seat_id_fk FOREIGN KEY (seat_id) REFERENCES seats (id);
ALTER TABLE bookings ADD CONSTRAINT bookings_seat_id_fk FOREIGN KEY (seat_id) REFERENCES seats (id);
ALTER TABLE events ADD CONSTRAINT events_discount_event_id_fk FOREIGN KEY (id) REFERENCES discounts (event_id);
ALTER TABLE applied_discounts ADD CONSTRAINT applied_discounts_booking_id_fk FOREIGN KEY (booking_id) REFERENCES bookings (id);
ALTER TABLE applied_discounts ADD CONSTRAINT applied_discounts_discount_id_fk FOREIGN KEY (discount_id) REFERENCES discounts (id);
ALTER TABLE refunds ADD CONSTRAINT refunds_booking_id_fk FOREIGN KEY (booking_id) REFERENCES bookings (id);
ALTER TABLE refunds ADD CONSTRAINT refunds_user_id_fk FOREIGN KEY (user_id) REFERENCES users (id);