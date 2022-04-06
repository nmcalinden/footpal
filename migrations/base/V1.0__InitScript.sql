-- FOOTPAL Tables

/*
	Reference Tables
*/

CREATE TABLE booking_status_ref
(
    booking_status_id	INTEGER PRIMARY KEY NOT NULL,
    status_name			VARCHAR(10) NOT NULL
);

CREATE TABLE match_status_ref
(
    match_status_id 	INTEGER PRIMARY KEY NOT NULL,
    status_name			VARCHAR(10) NOT NULL
);

CREATE TABLE match_access_status_ref
(
    match_access_status_id	INTEGER PRIMARY KEY NOT NULL,
    status_name				VARCHAR(10) NOT NULL
);

CREATE TABLE payment_type_ref
(
    payment_type_id		INTEGER PRIMARY KEY NOT NULL,
    payment				VARCHAR(10) NOT NULL
);

/*
	Venue
*/
CREATE TABLE venue
(
    venue_id			INTEGER PRIMARY KEY NOT NULL,
    venue_name 			VARCHAR(100) NOT NULL,
    venue_address		VARCHAR(100) NOT NULL,
    postcode			VARCHAR(8),
    city 				VARCHAR(50),
    phone_no			VARCHAR(15),
    email 				VARCHAR(100),
    opening_hours		jsonb
);

CREATE TABLE pitch
(
    pitch_id			INTEGER PRIMARY KEY NOT NULL,
    venue_id 			INTEGER NOT NULL,
    pitch_name			VARCHAR(50) NOT NULL,
    max_players			INTEGER,
    cost 				DECIMAL,
    CONSTRAINT fk_venue_pitch_id FOREIGN KEY (venue_id)
        REFERENCES venue(venue_id)
);

/*
	Users
*/
CREATE TABLE footpal_user
(
    footpal_user_id INTEGER PRIMARY KEY NOT NULL,
    forename 		VARCHAR(50) NOT NULL,
    surname 		VARCHAR(50) NOT NULL,
    email			VARCHAR(100) NOT NULL
);

CREATE TABLE player
(
    player_id 		INTEGER PRIMARY KEY NOT NULL,
    footpal_user_id INTEGER NOT NULL,
    nickname		VARCHAR(16),
    phone_no		VARCHAR(15),
    postcode		VARCHAR(8),
    city			VARCHAR(50),
    CONSTRAINT fk_player_user_id FOREIGN KEY (footpal_user_id)
        REFERENCES footpal_user(footpal_user_id)
);

CREATE TABLE venue_admin
(
    venue_admin_id	INTEGER PRIMARY KEY NOT NULL,
    footpal_user_id INTEGER NOT NULL,
    venue_id 		INTEGER NOT NULL,
    CONSTRAINT fk_venue_user_id FOREIGN KEY (footpal_user_id)
        REFERENCES footpal_user(footpal_user_id),
    CONSTRAINT fk_venue_admin_id FOREIGN KEY (venue_id)
        REFERENCES venue(venue_id)
);


/*
	Groups
*/

CREATE TABLE group_details
(
    group_id		INTEGER PRIMARY KEY NOT NULL,
    group_name		VARCHAR(30),
    city			VARCHAR(50)
);

CREATE TABLE player_group
(
    player_group_id	INTEGER PRIMARY KEY NOT NULL,
    group_id 		INTEGER NOT NULL,
    player_id 		INTEGER NOT NULL,
    CONSTRAINT fk_player_group_id FOREIGN KEY (player_id)
        REFERENCES player(player_id)
        ON DELETE CASCADE,
    CONSTRAINT fk_player_group_group_id FOREIGN KEY (group_id)
        REFERENCES group_details(group_id)
        ON DELETE CASCADE
);

CREATE TABLE player_group_admin
(
    player_group_admin_id	INTEGER PRIMARY KEY NOT NULL,
    group_id 				INTEGER NOT NULL,
    player_id 				INTEGER NOT NULL,
    CONSTRAINT fk_player_admin_group_id FOREIGN KEY (group_id)
        REFERENCES group_details(group_id)
        ON DELETE CASCADE,
    CONSTRAINT fk_player_admin_player_id FOREIGN KEY (player_id)
        REFERENCES player(player_id)
        ON DELETE CASCADE
);


/*
	Match Bookings
*/
CREATE TABLE booking
(
    booking_id			INTEGER PRIMARY KEY NOT NULL,
    player_id			INTEGER,
    booking_status_id	INTEGER,
    created 			TIMESTAMP,
    last_updated		TIMESTAMP,
    CONSTRAINT fk_booking_player_id FOREIGN KEY (player_id)
        REFERENCES player(player_id)
            ON DELETE SET NULL,
    CONSTRAINT fk_booking_status_id FOREIGN KEY (booking_status_id)
        REFERENCES booking_status_ref(booking_status_id)
);

CREATE TABLE pitch_time_slot
(
    pitch_time_slot_id 	INTEGER PRIMARY KEY NOT NULL,
    pitch_id			INTEGER NOT NULL,
    day_of_week 		VARCHAR(10) NOT NULL,
    start_time			time,
    end_time			time,
    CONSTRAINT fk_pitch_time_slot_pitch_id FOREIGN KEY (pitch_id)
        REFERENCES pitch(pitch_id)
);

CREATE TABLE pitch_slot
(
    pitch_slot_id 		INTEGER PRIMARY KEY NOT NULL,
    booking_id			INTEGER NOT NULL,
    pitch_time_slot_id 	INTEGER NOT NULL,
    match_date			DATE,
    booking_status_id 	INTEGER NOT NULL,
    CONSTRAINT fk_pitch_slot_booking_id FOREIGN KEY (booking_id)
        REFERENCES booking(booking_id),
	CONSTRAINT fk_pitch_time_shot_id FOREIGN KEY (pitch_time_slot_id)
        REFERENCES pitch_time_slot(pitch_time_slot_id),
    CONSTRAINT fk_pitch_slot_status_id FOREIGN KEY (booking_status_id)
        REFERENCES booking_status_ref(booking_status_id)
);

CREATE TABLE match
(
    match_id				INTEGER PRIMARY KEY NOT NULL,
    booking_id				INTEGER NOT NULL,
    match_access_status_id	INTEGER NOT NULL,
    match_status_id			INTEGER NOT NULL,
    cost					DECIMAL,
    is_paid  				BOOLEAN,
    created 				TIMESTAMP,
    last_updated			TIMESTAMP,
    CONSTRAINT fk_match_booking_id FOREIGN KEY (booking_id)
        REFERENCES booking(booking_id),
    CONSTRAINT fk_match_access_id FOREIGN KEY (match_access_status_id)
        REFERENCES match_access_status_ref(match_access_status_id),
    CONSTRAINT fk_match_status_id FOREIGN KEY (match_status_id)
        REFERENCES match_status_ref(match_status_id)
);

CREATE TABLE match_player
(
    match_player_id		INTEGER PRIMARY KEY NOT NULL,
    match_id			INTEGER NOT NULL,
    player_id			INTEGER NOT NULL,
    amount_to_pay 		DECIMAL,
    payment_type_id		INTEGER NOT NULL,
    CONSTRAINT fk_match_player_match_id FOREIGN KEY (match_id)
        REFERENCES match(match_id),
    CONSTRAINT fk_match_player_player_id FOREIGN KEY (player_id)
        REFERENCES player(player_id)
        ON DELETE CASCADE,
    CONSTRAINT fk_match_player_payment_id FOREIGN KEY (payment_type_id)
        REFERENCES payment_type_ref(payment_type_id)
);

