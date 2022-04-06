-- Booking Status
INSERT INTO booking_status_ref VALUES (1, 'approved');
INSERT INTO booking_status_ref VALUES (2, 'cancelled');
INSERT INTO booking_status_ref VALUES (3, 'denied');
INSERT INTO booking_status_ref VALUES (4, 'pending');

-- Match Status
INSERT INTO match_status_ref VALUES (1, 'active');
INSERT INTO match_status_ref VALUES (2, 'cancelled');
INSERT INTO match_status_ref VALUES (3, 'completed');
INSERT INTO match_status_ref VALUES (4, 'pending');

-- Match Access Level
INSERT INTO match_access_status_ref VALUES (1, 'public');
INSERT INTO match_access_status_ref VALUES (2, 'private');

-- Payment Types
INSERT INTO payment_type_ref VALUES (1, 'card');
INSERT INTO payment_type_ref VALUES (2, 'cash');