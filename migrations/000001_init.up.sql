CREATE TYPE statuses AS ENUM ('waiting_ocr', 'waiting_validation', 'waiting_approval', 'closed','rejected');

CREATE TABLE ticket
(
    id           TEXT PRIMARY KEY,
    user_id      TEXT      NOT NULL,
    shop_address TEXT      NOT NULL,
    created_at   TIMESTAMP NOT NULL DEFAULT NOW(),
    image_url    TEXT      NOT NULL,
    updated_at   TIMESTAMP,
    status       statuses  NOT NULL DEFAULT 'waiting_ocr'
);

CREATE TABLE tickets_item
(
    ticket_id   TEXT    not null unique references ticket (id) on delete cascade,
    product     TEXT    not null,
    description TEXT    not null,
    price       decimal not null,
    amount      int     not null,
    unit        TEXT    not null,
    overprice   decimal
);

CREATE FUNCTION get_ticket_status(ticket_id TEXT) RETURNS statuses AS
$$
BEGIN
    RETURN (SELECT status
            FROM ticket
            WHERE id = ticket_id);
END
$$ LANGUAGE plpgsql;

CREATE TABLE rejection_reason
(
    ticket_id TEXT NOT NULL UNIQUE REFERENCES ticket (id) ON DELETE CASCADE CHECK (get_ticket_status(ticket_id) = 'rejected'),
    reason    TEXT NOT NULL

);
CREATE INDEX ON rejection_reason (ticket_id);

CREATE TABLE ticket_log
(
    ticket_id TEXT      NOT NULL UNIQUE REFERENCES ticket (id) ON DELETE CASCADE CHECK (get_ticket_status(ticket_id) = 'closed'),
    who       TEXT      NOT NULL,
    closed_at TIMESTAMP NOT NULL DEFAULT NOW()
);