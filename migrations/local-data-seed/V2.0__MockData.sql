-- Init Mock Users

INSERT INTO footpaldb.public.footpal_user(forename, surname, email) VALUES ('John', 'Bloggs', 'jbloggs@gmailcom');
INSERT INTO footpaldb.public.footpal_user(forename, surname, email) VALUES ('Mary', 'Foster', 'mfoster@gmailcom');
INSERT INTO footpaldb.public.footpal_user(forename, surname, email) VALUES ('Pete', 'West', 'pwest@gmailcom');
INSERT INTO footpaldb.public.footpal_user(forename, surname, email) VALUES ('Mike', 'Turner', 'mturner-boyyy@gmailcom');
INSERT INTO footpaldb.public.footpal_user(forename, surname, email) VALUES ('Lucy', 'Luck', 'll123@gmailcom');
INSERT INTO footpaldb.public.footpal_user(forename, surname, email) VALUES ('Tom', 'Smith', 'tsmith@gmailcom');
INSERT INTO footpaldb.public.footpal_user(forename, surname, email) VALUES ('Gareth', 'Clone', 'gclone@gmailcom');
INSERT INTO footpaldb.public.footpal_user(forename, surname, email) VALUES ('Randy', 'Marsh', 'rthemarsh123@gmailcom');
INSERT INTO footpaldb.public.footpal_user(forename, surname, email) VALUES ('Vince', 'Vaughn', 'vincey@gmailcom');
INSERT INTO footpaldb.public.footpal_user(forename, surname, email) VALUES ('Owen', 'Wils', 'owils9800@gmailcom');
INSERT INTO footpaldb.public.footpal_user(forename, surname, email) VALUES ('Kevy', 'Fergie', 'fergieboy@gmailcom');
INSERT INTO footpaldb.public.footpal_user(forename, surname, email) VALUES ('Betty', 'Fulton', 'fulton.b@gmailcom');
INSERT INTO footpaldb.public.footpal_user(forename, surname, email) VALUES ('Ally', 'Baba', 'ally.baba@gmailcom');
INSERT INTO footpaldb.public.footpal_user(forename, surname, email) VALUES ('Timmy', 'Turner', 'tt@gmailcom');
INSERT INTO footpaldb.public.footpal_user(forename, surname, email) VALUES ('Drew', 'OnPage', 'haha@gmailcom');
INSERT INTO footpaldb.public.footpal_user(forename, surname, email) VALUES ('Joel', 'Crowe', 'thejoel@gmailcom');
INSERT INTO footpaldb.public.footpal_user(forename, surname, email) VALUES ('Yeti', 'Netty', 'yeti@gmailcom');
INSERT INTO footpaldb.public.footpal_user(forename, surname, email) VALUES ('Harold', 'Hart', 'thehartyboy@gmailcom');
INSERT INTO footpaldb.public.footpal_user(forename, surname, email) VALUES ('Beth', 'Brock', 'rockthebrock8382@gmailcom');
INSERT INTO footpaldb.public.footpal_user(forename, surname, email) VALUES ('Lionel', 'Messi', 'messi10@gmailcom');

-- Init Venues

INSERT INTO footpaldb.public.venue(venue_name, venue_address, postcode, city, phone_no, email) VALUES ('Shankill Sports Arena', '123 Shankill Street', 'BT87 3QW', 'Belfast', '07824712319', 'ssa@booking.com');
INSERT INTO footpaldb.public.venue(venue_name, venue_address, postcode, city, phone_no, email) VALUES ('Lurgan Liars Soccer Dome', '95 William Street', 'BT66 8RT', 'Belfast', '07824712319', 'lurgansoccerdome@info.co.uk');

-- Init Venue Admins
INSERT INTO footpaldb.public.venue_admin(footpal_user_id, venue_id) VALUES (2, 1);
INSERT INTO footpaldb.public.venue_admin(footpal_user_id, venue_id) VALUES (6, 2);
INSERT INTO footpaldb.public.venue_admin(footpal_user_id, venue_id) VALUES (4, 2);

-- Init Players
INSERT INTO footpaldb.public.player(footpal_user_id, nickname, phone_no, postcode, city) VALUES (1, null, '07858274821', 'BT34 5RW', 'Craigavon');
INSERT INTO footpaldb.public.player(footpal_user_id, nickname, phone_no, postcode, city) VALUES (3, 'Blazer', '07858274821', 'BT34 5RW', 'Belfast');
INSERT INTO footpaldb.public.player(footpal_user_id, nickname, phone_no, postcode, city) VALUES (5, null, '07858274821', 'BT98 6UY', 'Belfast');
INSERT INTO footpaldb.public.player(footpal_user_id, nickname, phone_no, postcode, city) VALUES (7, null, '07858274821', 'BT34 5RW', 'Lurgan');
INSERT INTO footpaldb.public.player(footpal_user_id, nickname, phone_no, postcode, city) VALUES (8, 'Rocket', '07858274821', 'BT34 5RW', 'Craigavon');
INSERT INTO footpaldb.public.player(footpal_user_id, nickname, phone_no, postcode, city) VALUES (9, null, '07858274821', 'BT98 6UY', 'Belfast');
INSERT INTO footpaldb.public.player(footpal_user_id, nickname, phone_no, postcode, city) VALUES (10, null, '07858274821', 'BT34 5RW', 'Craigavon');
INSERT INTO footpaldb.public.player(footpal_user_id, nickname, phone_no, postcode, city) VALUES (11, null, '07858274821', 'BT34 5RW', 'Larne');
INSERT INTO footpaldb.public.player(footpal_user_id, nickname, phone_no, postcode, city) VALUES (12, 'Looner', '07858274821', 'BT98 6UY', 'Lurgan');
INSERT INTO footpaldb.public.player(footpal_user_id, nickname, phone_no, postcode, city) VALUES (13, null, '07858274821', 'BT34 5RW', 'Lisburn');
INSERT INTO footpaldb.public.player(footpal_user_id, nickname, phone_no, postcode, city) VALUES (14, 'Top Bins', '07858274821', 'BT34 5RW', 'Craigavon');
INSERT INTO footpaldb.public.player(footpal_user_id, nickname, phone_no, postcode, city) VALUES (15, null, '07858274821', 'BT98 6UY', 'Belfast');
INSERT INTO footpaldb.public.player(footpal_user_id, nickname, phone_no, postcode, city) VALUES (16, 'Magic', '07858274821', 'BT98 6UY', 'Lisburn');
INSERT INTO footpaldb.public.player(footpal_user_id, nickname, phone_no, postcode, city) VALUES (17, null, '07858274821', 'BT34 5RW', 'Craigavon');
INSERT INTO footpaldb.public.player(footpal_user_id, nickname, phone_no, postcode, city) VALUES (18, null, '07858274821', 'BT34 5RW', 'Lisburn');
INSERT INTO footpaldb.public.player(footpal_user_id, nickname, phone_no, postcode, city) VALUES (19, null, '07858274821', 'BT98 6UY', 'Belfast');
INSERT INTO footpaldb.public.player(footpal_user_id, nickname, phone_no, postcode, city) VALUES (20, 'Leo', '07858274821', 'BT98 6UY', 'Lurgan');

-- Init Pitches
INSERT INTO footpaldb.public.pitch(venue_id, pitch_name, max_players, cost) VALUES(1, 'Pitch 1', 10, 30.0);
INSERT INTO footpaldb.public.pitch(venue_id, pitch_name, max_players, cost) VALUES(1, 'Pitch 2', 10, 30.0);
INSERT INTO footpaldb.public.pitch(venue_id, pitch_name, max_players, cost) VALUES(1, 'Pitch 3', 14, 50.0);
INSERT INTO footpaldb.public.pitch(venue_id, pitch_name, max_players, cost) VALUES(2, 'Lennon Arena', 10, 40.0);
INSERT INTO footpaldb.public.pitch(venue_id, pitch_name, max_players, cost) VALUES(2, 'The Big one', 14, 50.0);

-- Init Pitch Time Slots
INSERT INTO footpaldb.public.pitch_time_slot(pitch_id, day_of_week, start_time, end_time)  VALUES(1, 'Monday', '13:00:00', '14:00:00');
INSERT INTO footpaldb.public.pitch_time_slot(pitch_id, day_of_week, start_time, end_time)  VALUES(1, 'Monday', '14:00:00', '15:00:00');
INSERT INTO footpaldb.public.pitch_time_slot(pitch_id, day_of_week, start_time, end_time)  VALUES(1, 'Monday', '15:00:00', '16:00:00');
INSERT INTO footpaldb.public.pitch_time_slot(pitch_id, day_of_week, start_time, end_time)  VALUES(1, 'Monday', '16:00:00', '17:00:00');
INSERT INTO footpaldb.public.pitch_time_slot(pitch_id, day_of_week, start_time, end_time)  VALUES(1, 'Monday', '17:00:00', '18:00:00');
INSERT INTO footpaldb.public.pitch_time_slot(pitch_id, day_of_week, start_time, end_time)  VALUES(1, 'Monday', '18:00:00', '19:00:00');
INSERT INTO footpaldb.public.pitch_time_slot(pitch_id, day_of_week, start_time, end_time)  VALUES(1, 'Monday', '19:00:00', '20:00:00');
INSERT INTO footpaldb.public.pitch_time_slot(pitch_id, day_of_week, start_time, end_time)  VALUES(1, 'Monday', '20:00:00', '21:00:00');
INSERT INTO footpaldb.public.pitch_time_slot(pitch_id, day_of_week, start_time, end_time)  VALUES(1, 'Monday', '21:00:00', '22:00:00');

INSERT INTO footpaldb.public.pitch_time_slot(pitch_id, day_of_week, start_time, end_time)  VALUES(1, 'Tuesday', '13:00:00', '14:00:00');
INSERT INTO footpaldb.public.pitch_time_slot(pitch_id, day_of_week, start_time, end_time)  VALUES(1, 'Tuesday', '14:00:00', '15:00:00');
INSERT INTO footpaldb.public.pitch_time_slot(pitch_id, day_of_week, start_time, end_time)  VALUES(1, 'Tuesday', '15:00:00', '16:00:00');
INSERT INTO footpaldb.public.pitch_time_slot(pitch_id, day_of_week, start_time, end_time)  VALUES(1, 'Tuesday', '16:00:00', '17:00:00');
INSERT INTO footpaldb.public.pitch_time_slot(pitch_id, day_of_week, start_time, end_time)  VALUES(1, 'Tuesday', '17:00:00', '18:00:00');
INSERT INTO footpaldb.public.pitch_time_slot(pitch_id, day_of_week, start_time, end_time)  VALUES(1, 'Tuesday', '18:00:00', '19:00:00');
INSERT INTO footpaldb.public.pitch_time_slot(pitch_id, day_of_week, start_time, end_time)  VALUES(1, 'Tuesday', '19:00:00', '20:00:00');
INSERT INTO footpaldb.public.pitch_time_slot(pitch_id, day_of_week, start_time, end_time)  VALUES(1, 'Tuesday', '20:00:00', '21:00:00');
INSERT INTO footpaldb.public.pitch_time_slot(pitch_id, day_of_week, start_time, end_time)  VALUES(1, 'Tuesday', '21:00:00', '22:00:00');

INSERT INTO footpaldb.public.pitch_time_slot(pitch_id, day_of_week, start_time, end_time)  VALUES(1, 'Wednesday', '13:00:00', '14:00:00');
INSERT INTO footpaldb.public.pitch_time_slot(pitch_id, day_of_week, start_time, end_time)  VALUES(1, 'Wednesday', '14:00:00', '15:00:00');
INSERT INTO footpaldb.public.pitch_time_slot(pitch_id, day_of_week, start_time, end_time)  VALUES(1, 'Wednesday', '15:00:00', '16:00:00');
INSERT INTO footpaldb.public.pitch_time_slot(pitch_id, day_of_week, start_time, end_time)  VALUES(1, 'Wednesday', '16:00:00', '17:00:00');
INSERT INTO footpaldb.public.pitch_time_slot(pitch_id, day_of_week, start_time, end_time)  VALUES(1, 'Wednesday', '17:00:00', '18:00:00');
INSERT INTO footpaldb.public.pitch_time_slot(pitch_id, day_of_week, start_time, end_time)  VALUES(1, 'Wednesday', '18:00:00', '19:00:00');
INSERT INTO footpaldb.public.pitch_time_slot(pitch_id, day_of_week, start_time, end_time)  VALUES(1, 'Wednesday', '19:00:00', '20:00:00');
INSERT INTO footpaldb.public.pitch_time_slot(pitch_id, day_of_week, start_time, end_time)  VALUES(1, 'Wednesday', '20:00:00', '21:00:00');
INSERT INTO footpaldb.public.pitch_time_slot(pitch_id, day_of_week, start_time, end_time)  VALUES(1, 'Wednesday', '21:00:00', '22:00:00');

INSERT INTO footpaldb.public.pitch_time_slot(pitch_id, day_of_week, start_time, end_time)  VALUES(1, 'Thursday', '13:00:00', '14:00:00');
INSERT INTO footpaldb.public.pitch_time_slot(pitch_id, day_of_week, start_time, end_time)  VALUES(1, 'Thursday', '14:00:00', '15:00:00');
INSERT INTO footpaldb.public.pitch_time_slot(pitch_id, day_of_week, start_time, end_time)  VALUES(1, 'Thursday', '15:00:00', '16:00:00');
INSERT INTO footpaldb.public.pitch_time_slot(pitch_id, day_of_week, start_time, end_time)  VALUES(1, 'Thursday', '16:00:00', '17:00:00');
INSERT INTO footpaldb.public.pitch_time_slot(pitch_id, day_of_week, start_time, end_time)  VALUES(1, 'Thursday', '17:00:00', '18:00:00');
INSERT INTO footpaldb.public.pitch_time_slot(pitch_id, day_of_week, start_time, end_time)  VALUES(1, 'Thursday', '18:00:00', '19:00:00');
INSERT INTO footpaldb.public.pitch_time_slot(pitch_id, day_of_week, start_time, end_time)  VALUES(1, 'Thursday', '19:00:00', '20:00:00');
INSERT INTO footpaldb.public.pitch_time_slot(pitch_id, day_of_week, start_time, end_time)  VALUES(1, 'Thursday', '20:00:00', '21:00:00');
INSERT INTO footpaldb.public.pitch_time_slot(pitch_id, day_of_week, start_time, end_time)  VALUES(1, 'Thursday', '21:00:00', '22:00:00');

INSERT INTO footpaldb.public.pitch_time_slot(pitch_id, day_of_week, start_time, end_time)  VALUES(1, 'Friday', '13:00:00', '14:00:00');
INSERT INTO footpaldb.public.pitch_time_slot(pitch_id, day_of_week, start_time, end_time)  VALUES(1, 'Friday', '14:00:00', '15:00:00');
INSERT INTO footpaldb.public.pitch_time_slot(pitch_id, day_of_week, start_time, end_time)  VALUES(1, 'Friday', '15:00:00', '16:00:00');
INSERT INTO footpaldb.public.pitch_time_slot(pitch_id, day_of_week, start_time, end_time)  VALUES(1, 'Friday', '16:00:00', '17:00:00');
INSERT INTO footpaldb.public.pitch_time_slot(pitch_id, day_of_week, start_time, end_time)  VALUES(1, 'Friday', '17:00:00', '18:00:00');
INSERT INTO footpaldb.public.pitch_time_slot(pitch_id, day_of_week, start_time, end_time)  VALUES(1, 'Friday', '18:00:00', '19:00:00');
INSERT INTO footpaldb.public.pitch_time_slot(pitch_id, day_of_week, start_time, end_time)  VALUES(1, 'Friday', '19:00:00', '20:00:00');
INSERT INTO footpaldb.public.pitch_time_slot(pitch_id, day_of_week, start_time, end_time)  VALUES(1, 'Friday', '20:00:00', '21:00:00');
INSERT INTO footpaldb.public.pitch_time_slot(pitch_id, day_of_week, start_time, end_time)  VALUES(1, 'Friday', '21:00:00', '22:00:00');

INSERT INTO footpaldb.public.pitch_time_slot(pitch_id, day_of_week, start_time, end_time)  VALUES(1, 'Saturday', '09:00:00', '10:00:00');
INSERT INTO footpaldb.public.pitch_time_slot(pitch_id, day_of_week, start_time, end_time)  VALUES(1, 'Saturday', '10:00:00', '11:00:00');
INSERT INTO footpaldb.public.pitch_time_slot(pitch_id, day_of_week, start_time, end_time)  VALUES(1, 'Saturday', '11:00:00', '12:00:00');
INSERT INTO footpaldb.public.pitch_time_slot(pitch_id, day_of_week, start_time, end_time)  VALUES(1, 'Saturday', '12:00:00', '13:00:00');
INSERT INTO footpaldb.public.pitch_time_slot(pitch_id, day_of_week, start_time, end_time)  VALUES(1, 'Saturday', '13:00:00', '14:00:00');
INSERT INTO footpaldb.public.pitch_time_slot(pitch_id, day_of_week, start_time, end_time)  VALUES(1, 'Saturday', '14:00:00', '15:00:00');
INSERT INTO footpaldb.public.pitch_time_slot(pitch_id, day_of_week, start_time, end_time)  VALUES(1, 'Saturday', '15:00:00', '16:00:00');
INSERT INTO footpaldb.public.pitch_time_slot(pitch_id, day_of_week, start_time, end_time)  VALUES(1, 'Saturday', '16:00:00', '17:00:00');

INSERT INTO footpaldb.public.pitch_time_slot(pitch_id, day_of_week, start_time, end_time)  VALUES(4, 'Sunday', '09:00:00', '10:00:00');
INSERT INTO footpaldb.public.pitch_time_slot(pitch_id, day_of_week, start_time, end_time)  VALUES(4, 'Sunday', '10:00:00', '11:00:00');
INSERT INTO footpaldb.public.pitch_time_slot(pitch_id, day_of_week, start_time, end_time)  VALUES(4, 'Sunday', '11:00:00', '12:00:00');
INSERT INTO footpaldb.public.pitch_time_slot(pitch_id, day_of_week, start_time, end_time)  VALUES(4, 'Sunday', '12:00:00', '13:00:00');
INSERT INTO footpaldb.public.pitch_time_slot(pitch_id, day_of_week, start_time, end_time)  VALUES(4, 'Sunday', '13:00:00', '14:00:00');
INSERT INTO footpaldb.public.pitch_time_slot(pitch_id, day_of_week, start_time, end_time)  VALUES(4, 'Sunday', '14:00:00', '15:00:00');
INSERT INTO footpaldb.public.pitch_time_slot(pitch_id, day_of_week, start_time, end_time)  VALUES(4, 'Sunday', '15:00:00', '16:00:00');
INSERT INTO footpaldb.public.pitch_time_slot(pitch_id, day_of_week, start_time, end_time)  VALUES(4, 'Sunday', '16:00:00', '17:00:00');

INSERT INTO footpaldb.public.pitch_time_slot(pitch_id, day_of_week, start_time, end_time)  VALUES(4, 'Monday', '13:00:00', '14:00:00');
INSERT INTO footpaldb.public.pitch_time_slot(pitch_id, day_of_week, start_time, end_time)  VALUES(4, 'Monday', '14:00:00', '15:00:00');
INSERT INTO footpaldb.public.pitch_time_slot(pitch_id, day_of_week, start_time, end_time)  VALUES(4, 'Monday', '15:00:00', '16:00:00');
INSERT INTO footpaldb.public.pitch_time_slot(pitch_id, day_of_week, start_time, end_time)  VALUES(4, 'Monday', '16:00:00', '17:00:00');
INSERT INTO footpaldb.public.pitch_time_slot(pitch_id, day_of_week, start_time, end_time)  VALUES(4, 'Monday', '17:00:00', '18:00:00');
INSERT INTO footpaldb.public.pitch_time_slot(pitch_id, day_of_week, start_time, end_time)  VALUES(4, 'Monday', '18:00:00', '19:00:00');
INSERT INTO footpaldb.public.pitch_time_slot(pitch_id, day_of_week, start_time, end_time)  VALUES(4, 'Monday', '19:00:00', '20:00:00');
INSERT INTO footpaldb.public.pitch_time_slot(pitch_id, day_of_week, start_time, end_time)  VALUES(4, 'Monday', '20:00:00', '21:00:00');
INSERT INTO footpaldb.public.pitch_time_slot(pitch_id, day_of_week, start_time, end_time)  VALUES(4, 'Monday', '21:00:00', '22:00:00');

INSERT INTO footpaldb.public.pitch_time_slot(pitch_id, day_of_week, start_time, end_time)  VALUES(4, 'Tuesday', '13:00:00', '14:00:00');
INSERT INTO footpaldb.public.pitch_time_slot(pitch_id, day_of_week, start_time, end_time)  VALUES(4, 'Tuesday', '14:00:00', '15:00:00');
INSERT INTO footpaldb.public.pitch_time_slot(pitch_id, day_of_week, start_time, end_time)  VALUES(4, 'Tuesday', '15:00:00', '16:00:00');
INSERT INTO footpaldb.public.pitch_time_slot(pitch_id, day_of_week, start_time, end_time)  VALUES(4, 'Tuesday', '16:00:00', '17:00:00');
INSERT INTO footpaldb.public.pitch_time_slot(pitch_id, day_of_week, start_time, end_time)  VALUES(4, 'Tuesday', '17:00:00', '18:00:00');
INSERT INTO footpaldb.public.pitch_time_slot(pitch_id, day_of_week, start_time, end_time)  VALUES(4, 'Tuesday', '18:00:00', '19:00:00');
INSERT INTO footpaldb.public.pitch_time_slot(pitch_id, day_of_week, start_time, end_time)  VALUES(4, 'Tuesday', '19:00:00', '20:00:00');
INSERT INTO footpaldb.public.pitch_time_slot(pitch_id, day_of_week, start_time, end_time)  VALUES(4, 'Tuesday', '20:00:00', '21:00:00');
INSERT INTO footpaldb.public.pitch_time_slot(pitch_id, day_of_week, start_time, end_time)  VALUES(4, 'Tuesday', '21:00:00', '22:00:00');

INSERT INTO footpaldb.public.pitch_time_slot(pitch_id, day_of_week, start_time, end_time)  VALUES(4, 'Wednesday', '13:00:00', '14:00:00');
INSERT INTO footpaldb.public.pitch_time_slot(pitch_id, day_of_week, start_time, end_time)  VALUES(4, 'Wednesday', '14:00:00', '15:00:00');
INSERT INTO footpaldb.public.pitch_time_slot(pitch_id, day_of_week, start_time, end_time)  VALUES(4, 'Wednesday', '15:00:00', '16:00:00');
INSERT INTO footpaldb.public.pitch_time_slot(pitch_id, day_of_week, start_time, end_time)  VALUES(4, 'Wednesday', '16:00:00', '17:00:00');
INSERT INTO footpaldb.public.pitch_time_slot(pitch_id, day_of_week, start_time, end_time)  VALUES(4, 'Wednesday', '17:00:00', '18:00:00');
INSERT INTO footpaldb.public.pitch_time_slot(pitch_id, day_of_week, start_time, end_time)  VALUES(4, 'Wednesday', '18:00:00', '19:00:00');
INSERT INTO footpaldb.public.pitch_time_slot(pitch_id, day_of_week, start_time, end_time)  VALUES(4, 'Wednesday', '19:00:00', '20:00:00');
INSERT INTO footpaldb.public.pitch_time_slot(pitch_id, day_of_week, start_time, end_time)  VALUES(4, 'Wednesday', '20:00:00', '21:00:00');
INSERT INTO footpaldb.public.pitch_time_slot(pitch_id, day_of_week, start_time, end_time)  VALUES(4, 'Wednesday', '21:00:00', '22:00:00');

INSERT INTO footpaldb.public.pitch_time_slot(pitch_id, day_of_week, start_time, end_time)  VALUES(4, 'Thursday', '13:00:00', '14:00:00');
INSERT INTO footpaldb.public.pitch_time_slot(pitch_id, day_of_week, start_time, end_time)  VALUES(4, 'Thursday', '14:00:00', '15:00:00');
INSERT INTO footpaldb.public.pitch_time_slot(pitch_id, day_of_week, start_time, end_time)  VALUES(4, 'Thursday', '15:00:00', '16:00:00');
INSERT INTO footpaldb.public.pitch_time_slot(pitch_id, day_of_week, start_time, end_time)  VALUES(4, 'Thursday', '16:00:00', '17:00:00');
INSERT INTO footpaldb.public.pitch_time_slot(pitch_id, day_of_week, start_time, end_time)  VALUES(4, 'Thursday', '17:00:00', '18:00:00');
INSERT INTO footpaldb.public.pitch_time_slot(pitch_id, day_of_week, start_time, end_time)  VALUES(4, 'Thursday', '18:00:00', '19:00:00');
INSERT INTO footpaldb.public.pitch_time_slot(pitch_id, day_of_week, start_time, end_time)  VALUES(4, 'Thursday', '19:00:00', '20:00:00');
INSERT INTO footpaldb.public.pitch_time_slot(pitch_id, day_of_week, start_time, end_time)  VALUES(4, 'Thursday', '20:00:00', '21:00:00');
INSERT INTO footpaldb.public.pitch_time_slot(pitch_id, day_of_week, start_time, end_time)  VALUES(4, 'Thursday', '21:00:00', '22:00:00');

INSERT INTO footpaldb.public.pitch_time_slot(pitch_id, day_of_week, start_time, end_time)  VALUES(4, 'Friday', '13:00:00', '14:00:00');
INSERT INTO footpaldb.public.pitch_time_slot(pitch_id, day_of_week, start_time, end_time)  VALUES(4, 'Friday', '14:00:00', '15:00:00');
INSERT INTO footpaldb.public.pitch_time_slot(pitch_id, day_of_week, start_time, end_time)  VALUES(4, 'Friday', '15:00:00', '16:00:00');
INSERT INTO footpaldb.public.pitch_time_slot(pitch_id, day_of_week, start_time, end_time)  VALUES(4, 'Friday', '16:00:00', '17:00:00');
INSERT INTO footpaldb.public.pitch_time_slot(pitch_id, day_of_week, start_time, end_time)  VALUES(4, 'Friday', '17:00:00', '18:00:00');
INSERT INTO footpaldb.public.pitch_time_slot(pitch_id, day_of_week, start_time, end_time)  VALUES(4, 'Friday', '18:00:00', '19:00:00');
INSERT INTO footpaldb.public.pitch_time_slot(pitch_id, day_of_week, start_time, end_time)  VALUES(4, 'Friday', '19:00:00', '20:00:00');
INSERT INTO footpaldb.public.pitch_time_slot(pitch_id, day_of_week, start_time, end_time)  VALUES(4, 'Friday', '20:00:00', '21:00:00');
INSERT INTO footpaldb.public.pitch_time_slot(pitch_id, day_of_week, start_time, end_time)  VALUES(4, 'Friday', '21:00:00', '22:00:00');

INSERT INTO footpaldb.public.pitch_time_slot(pitch_id, day_of_week, start_time, end_time)  VALUES(4, 'Saturday', '09:00:00', '10:00:00');
INSERT INTO footpaldb.public.pitch_time_slot(pitch_id, day_of_week, start_time, end_time)  VALUES(4, 'Saturday', '10:00:00', '11:00:00');
INSERT INTO footpaldb.public.pitch_time_slot(pitch_id, day_of_week, start_time, end_time)  VALUES(4, 'Saturday', '11:00:00', '12:00:00');
INSERT INTO footpaldb.public.pitch_time_slot(pitch_id, day_of_week, start_time, end_time)  VALUES(4, 'Saturday', '12:00:00', '13:00:00');
INSERT INTO footpaldb.public.pitch_time_slot(pitch_id, day_of_week, start_time, end_time)  VALUES(4, 'Saturday', '13:00:00', '14:00:00');
INSERT INTO footpaldb.public.pitch_time_slot(pitch_id, day_of_week, start_time, end_time)  VALUES(4, 'Saturday', '14:00:00', '15:00:00');
INSERT INTO footpaldb.public.pitch_time_slot(pitch_id, day_of_week, start_time, end_time)  VALUES(4, 'Saturday', '15:00:00', '16:00:00');
INSERT INTO footpaldb.public.pitch_time_slot(pitch_id, day_of_week, start_time, end_time)  VALUES(4, 'Saturday', '16:00:00', '17:00:00');

INSERT INTO footpaldb.public.pitch_time_slot(pitch_id, day_of_week, start_time, end_time)  VALUES(4, 'Sunday', '09:00:00', '10:00:00');
INSERT INTO footpaldb.public.pitch_time_slot(pitch_id, day_of_week, start_time, end_time)  VALUES(4, 'Sunday', '10:00:00', '11:00:00');
INSERT INTO footpaldb.public.pitch_time_slot(pitch_id, day_of_week, start_time, end_time)  VALUES(4, 'Sunday', '11:00:00', '12:00:00');
INSERT INTO footpaldb.public.pitch_time_slot(pitch_id, day_of_week, start_time, end_time)  VALUES(4, 'Sunday', '12:00:00', '13:00:00');
INSERT INTO footpaldb.public.pitch_time_slot(pitch_id, day_of_week, start_time, end_time)  VALUES(4, 'Sunday', '13:00:00', '14:00:00');
INSERT INTO footpaldb.public.pitch_time_slot(pitch_id, day_of_week, start_time, end_time)  VALUES(4, 'Sunday', '14:00:00', '15:00:00');
INSERT INTO footpaldb.public.pitch_time_slot(pitch_id, day_of_week, start_time, end_time)  VALUES(4, 'Sunday', '15:00:00', '16:00:00');
INSERT INTO footpaldb.public.pitch_time_slot(pitch_id, day_of_week, start_time, end_time)  VALUES(4, 'Sunday', '16:00:00', '17:00:00');


-- Init Squads
INSERT INTO footpaldb.public.squad(squad_name, city) VALUES('Leos 5 a side', 'Belfast');
INSERT INTO footpaldb.public.squad(squad_name, city) VALUES('CR7 5 a side', 'Belfast');

-- Init Squad Players
INSERT INTO footpaldb.public.squad_player(squad_id, player_id, user_role, squad_player_status_id)  VALUES(1, 7, 'player', 1);
INSERT INTO footpaldb.public.squad_player(squad_id, player_id, user_role, squad_player_status_id)  VALUES(1, 11, 'player', 1);
INSERT INTO footpaldb.public.squad_player(squad_id, player_id, user_role, squad_player_status_id)  VALUES(1, 5, 'player', 1);
INSERT INTO footpaldb.public.squad_player(squad_id, player_id, user_role, squad_player_status_id)  VALUES(1, 15, 'player', 1);
INSERT INTO footpaldb.public.squad_player(squad_id, player_id, user_role, squad_player_status_id)  VALUES(1, 1, 'player', 1);
INSERT INTO footpaldb.public.squad_player(squad_id, player_id, user_role, squad_player_status_id)  VALUES(1, 14, 'player', 1);
INSERT INTO footpaldb.public.squad_player(squad_id, player_id, user_role, squad_player_status_id)  VALUES(1, 10, 'player', 1);
INSERT INTO footpaldb.public.squad_player(squad_id, player_id, user_role, squad_player_status_id)  VALUES(1, 17, 'player', 1);
INSERT INTO footpaldb.public.squad_player(squad_id, player_id, user_role, squad_player_status_id)  VALUES(1, 8, 'player', 1);
INSERT INTO footpaldb.public.squad_player(squad_id, player_id, user_role, squad_player_status_id)  VALUES(1, 6, 'admin', 1);

INSERT INTO footpaldb.public.squad_player(squad_id, player_id, user_role, squad_player_status_id)  VALUES(2, 9, 'player', 1);
INSERT INTO footpaldb.public.squad_player(squad_id, player_id, user_role, squad_player_status_id)  VALUES(2, 3, 'player', 1);
INSERT INTO footpaldb.public.squad_player(squad_id, player_id, user_role, squad_player_status_id)  VALUES(2, 12, 'player', 1);
INSERT INTO footpaldb.public.squad_player(squad_id, player_id, user_role, squad_player_status_id)  VALUES(2, 15, 'player', 1);
INSERT INTO footpaldb.public.squad_player(squad_id, player_id, user_role, squad_player_status_id)  VALUES(2, 16, 'player', 1);
INSERT INTO footpaldb.public.squad_player(squad_id, player_id, user_role, squad_player_status_id)  VALUES(2, 17, 'player', 1);
INSERT INTO footpaldb.public.squad_player(squad_id, player_id, user_role, squad_player_status_id)  VALUES(2, 16, 'player', 1);
INSERT INTO footpaldb.public.squad_player(squad_id, player_id, user_role, squad_player_status_id)  VALUES(2, 17, 'admin', 1);
INSERT INTO footpaldb.public.squad_player(squad_id, player_id, user_role, squad_player_status_id)  VALUES(2, 13, 'player', 1);

-- Init Bookings
INSERT INTO footpaldb.public.booking(booking_status_id, created_by, created, last_updated) VALUES(1, 17, current_timestamp, current_timestamp);
INSERT INTO footpaldb.public.booking(booking_status_id, created_by, created, last_updated) VALUES(4, 6, current_timestamp, current_timestamp);

-- Init Pitch Slots
INSERT INTO footpaldb.public.pitch_slot(booking_id, pitch_time_slot_id, match_date, booking_status_id) VALUES(1, 70, DATE '2022-04-20', 1);
INSERT INTO footpaldb.public.pitch_slot(booking_id, pitch_time_slot_id, match_date, booking_status_id) VALUES(1, 70, DATE '2022-04-27', 1);
INSERT INTO footpaldb.public.pitch_slot(booking_id, pitch_time_slot_id, match_date, booking_status_id) VALUES(1, 70, DATE '2022-05-04', 1);
INSERT INTO footpaldb.public.pitch_slot(booking_id, pitch_time_slot_id, match_date, booking_status_id) VALUES(2, 52, DATE '2022-04-25', 4);

-- Init Match
INSERT INTO footpaldb.public.match(booking_id, match_access_status_id, match_status_id, match_date, cost, is_paid, created, last_updated) VALUES(1, 2, 1, DATE '2022-04-20', 30.0, false, current_timestamp, current_timestamp);
INSERT INTO footpaldb.public.match(booking_id, match_access_status_id, match_status_id, match_date, cost, is_paid, created, last_updated) VALUES(1, 2, 4, DATE '2022-04-27', 30.0, false, current_timestamp, current_timestamp);
INSERT INTO footpaldb.public.match(booking_id, match_access_status_id, match_status_id, match_date, cost, is_paid, created, last_updated) VALUES(1, 2, 4, DATE '2022-05-04', 30.0, false, current_timestamp, current_timestamp);

-- Init Match Player
INSERT INTO footpaldb.public.match_player(match_id, player_id, amount_to_pay, payment_type_id)  VALUES(1, 4, 3.0, 1);
INSERT INTO footpaldb.public.match_player(match_id, player_id, amount_to_pay, payment_type_id)  VALUES(1, 8, 3.0, 1);
INSERT INTO footpaldb.public.match_player(match_id, player_id, amount_to_pay, payment_type_id)  VALUES(1, 3, 3.0, 1);
INSERT INTO footpaldb.public.match_player(match_id, player_id, amount_to_pay, payment_type_id)  VALUES(1, 15, 3.0, 2);
INSERT INTO footpaldb.public.match_player(match_id, player_id, amount_to_pay, payment_type_id)  VALUES(1, 1, 3.0, 1);
INSERT INTO footpaldb.public.match_player(match_id, player_id, amount_to_pay, payment_type_id)  VALUES(1, 11, 3.0, 1);
INSERT INTO footpaldb.public.match_player(match_id, player_id, amount_to_pay, payment_type_id)  VALUES(1, 7, 3.0, 2);
INSERT INTO footpaldb.public.match_player(match_id, player_id, amount_to_pay, payment_type_id)  VALUES(1, 17, 3.0, 2);
INSERT INTO footpaldb.public.match_player(match_id, player_id, amount_to_pay, payment_type_id)  VALUES(1, 5, 3.0, 1);
INSERT INTO footpaldb.public.match_player(match_id, player_id, amount_to_pay, payment_type_id)  VALUES(1, 10, 3.0, 1);