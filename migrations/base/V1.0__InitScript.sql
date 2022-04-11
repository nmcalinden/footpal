-- FOOTPAL Tables

/*
	Reference Tables
*/

CREATE TABLE booking_status_ref
(
    id	                INTEGER PRIMARY KEY NOT NULL,
    "description"		VARCHAR(25) NOT NULL
);

CREATE TABLE squad_player_status_ref
(
    id 	                INTEGER PRIMARY KEY NOT NULL,
    "description"		VARCHAR(25) NOT NULL
);

CREATE TABLE match_status_ref
(
    id 	                INTEGER PRIMARY KEY NOT NULL,
    "description"		VARCHAR(25) NOT NULL
);

CREATE TABLE match_access_status_ref
(
    id	                INTEGER PRIMARY KEY NOT NULL,
    "description"		VARCHAR(25) NOT NULL
);

CREATE TABLE payment_type_ref
(
    id		            INTEGER PRIMARY KEY NOT NULL,
    "description"		VARCHAR(25) NOT NULL
);

/*
	Venue
*/
CREATE TABLE venue
(
    id			        INTEGER PRIMARY KEY NOT NULL,
    venue_name   		VARCHAR(100) NOT NULL,
    venue_address      	VARCHAR(100) NOT NULL,
    postcode			VARCHAR(8),
    city 				VARCHAR(50),
    phone_no			VARCHAR(15),
    email 				VARCHAR(100)
);

CREATE TABLE pitch
(
    id			        INTEGER PRIMARY KEY NOT NULL,
    venue_id 			INTEGER NOT NULL,
    pitch_name			VARCHAR(50) NOT NULL,
    max_players			INTEGER,
    cost 				DECIMAL,
    CONSTRAINT fk_venue_pitch_id FOREIGN KEY (venue_id)
        REFERENCES venue(id)
);

/*
	Users
*/
CREATE TABLE footpal_user
(
    id              INTEGER PRIMARY KEY NOT NULL,
    forename 		VARCHAR(50) NOT NULL,
    surname 		VARCHAR(50) NOT NULL,
    email			VARCHAR(100) NOT NULL
);

CREATE TABLE player
(
    id      		INTEGER PRIMARY KEY NOT NULL,
    footpal_user_id INTEGER NOT NULL,
    nickname		VARCHAR(16),
    phone_no		VARCHAR(15),
    postcode		VARCHAR(8),
    city			VARCHAR(50),
    CONSTRAINT fk_player_user_id FOREIGN KEY (footpal_user_id)
        REFERENCES footpal_user(id),
    CONSTRAINT user_unique UNIQUE (footpal_user_id)
);

CREATE TABLE venue_admin
(
    id	            INTEGER PRIMARY KEY NOT NULL,
    footpal_user_id INTEGER NOT NULL,
    venue_id 		INTEGER NOT NULL,
    CONSTRAINT fk_venue_user_id FOREIGN KEY (footpal_user_id)
        REFERENCES footpal_user(id),
    CONSTRAINT fk_venue_admin_id FOREIGN KEY (venue_id)
        REFERENCES venue(id)
);


/*
	Squads
*/

CREATE TABLE squad
(
    id		        INTEGER PRIMARY KEY NOT NULL,
    squad_name		VARCHAR(30),
    city			VARCHAR(50)
);

CREATE TABLE squad_player
(
    id	                    INTEGER PRIMARY KEY NOT NULL,
    squad_id 		        INTEGER NOT NULL,
    player_id 		        INTEGER NOT NULL,
    user_role               VARCHAR(20),
    squad_player_status_id  INTEGER NOT NULL,
    CONSTRAINT fk_squad_player_id FOREIGN KEY (player_id)
        REFERENCES player(id),
    CONSTRAINT fk_squad_player_squad_id FOREIGN KEY (squad_id)
        REFERENCES squad(id),
    CONSTRAINT fk_squad_player_status_id FOREIGN KEY (squad_player_status_id)
        REFERENCES squad_player_status_ref(id)
);

/*
	Match Bookings
*/
CREATE TABLE booking
(
    id	        		INTEGER PRIMARY KEY NOT NULL,
    booking_status_id	INTEGER,
    created_by          INTEGER,
    created 			TIMESTAMP,
    last_updated		TIMESTAMP,
    CONSTRAINT fk_booking_player_id FOREIGN KEY (created_by)
        REFERENCES player(id),
    CONSTRAINT fk_booking_status_id FOREIGN KEY (booking_status_id)
        REFERENCES booking_status_ref(id)
);

CREATE TABLE pitch_time_slot
(
    id 	                INTEGER PRIMARY KEY NOT NULL,
    pitch_id			INTEGER NOT NULL,
    day_of_week 		VARCHAR(10) NOT NULL,
    start_time			time,
    end_time			time,
    CONSTRAINT fk_pitch_time_slot_pitch_id FOREIGN KEY (pitch_id)
        REFERENCES pitch(id)
);

CREATE TABLE pitch_slot
(
    id          		INTEGER PRIMARY KEY NOT NULL,
    booking_id			INTEGER NOT NULL,
    pitch_time_slot_id 	INTEGER NOT NULL,
    match_date			DATE,
    booking_status_id 	INTEGER NOT NULL,
    CONSTRAINT fk_pitch_slot_booking_id FOREIGN KEY (booking_id)
        REFERENCES booking(id),
	CONSTRAINT fk_pitch_time_shot_id FOREIGN KEY (pitch_time_slot_id)
        REFERENCES pitch_time_slot(id),
    CONSTRAINT fk_pitch_slot_status_id FOREIGN KEY (booking_status_id)
        REFERENCES booking_status_ref(id)
);

CREATE TABLE match
(
    id	        			INTEGER PRIMARY KEY NOT NULL,
    booking_id				INTEGER NOT NULL,
    match_access_status_id	INTEGER NOT NULL,
    match_status_id			INTEGER NOT NULL,
    match_date              DATE,
    cost					DECIMAL,
    is_paid  				BOOLEAN,
    created 				TIMESTAMP,
    last_updated			TIMESTAMP,
    CONSTRAINT fk_match_booking_id FOREIGN KEY (booking_id)
        REFERENCES booking(id),
    CONSTRAINT fk_match_access_id FOREIGN KEY (match_access_status_id)
        REFERENCES match_access_status_ref(id),
    CONSTRAINT fk_match_status_id FOREIGN KEY (match_status_id)
        REFERENCES match_status_ref(id)
);

CREATE TABLE match_player
(
    id	            	INTEGER PRIMARY KEY NOT NULL,
    match_id			INTEGER NOT NULL,
    player_id			INTEGER NOT NULL,
    amount_to_pay 		DECIMAL,
    payment_type_id		INTEGER NOT NULL,
    CONSTRAINT fk_match_player_match_id FOREIGN KEY (match_id)
        REFERENCES match(id),
    CONSTRAINT fk_match_player_player_id FOREIGN KEY (player_id)
        REFERENCES player(id)
        ON DELETE CASCADE,
    CONSTRAINT fk_match_player_payment_id FOREIGN KEY (payment_type_id)
        REFERENCES payment_type_ref(id)
);

