-- Init Mock Users

INSERT INTO footpal_user VALUES (1, 'John', 'Bloggs', 'jbloggs@gmailcom');
INSERT INTO footpal_user VALUES (2, 'Mary', 'Foster', 'mfoster@gmailcom');
INSERT INTO footpal_user VALUES (3, 'Pete', 'West', 'pwest@gmailcom');
INSERT INTO footpal_user VALUES (4, 'Mike', 'Turner', 'mturner-boyyy@gmailcom');
INSERT INTO footpal_user VALUES (5, 'Lucy', 'Luck', 'll123@gmailcom');
INSERT INTO footpal_user VALUES (6, 'Tom', 'Smith', 'tsmith@gmailcom');
INSERT INTO footpal_user VALUES (7, 'Gareth', 'Clone', 'gclone@gmailcom');
INSERT INTO footpal_user VALUES (8, 'Randy', 'Marsh', 'rthemarsh123@gmailcom');
INSERT INTO footpal_user VALUES (9, 'Vince', 'Vaughn', 'vincey@gmailcom');
INSERT INTO footpal_user VALUES (10, 'Owen', 'Wils', 'owils9800@gmailcom');
INSERT INTO footpal_user VALUES (11, 'Kevy', 'Fergie', 'fergieboy@gmailcom');
INSERT INTO footpal_user VALUES (12, 'Betty', 'Fulton', 'fulton.b@gmailcom');
INSERT INTO footpal_user VALUES (13, 'Ally', 'Baba', 'ally.baba@gmailcom');
INSERT INTO footpal_user VALUES (14, 'Timmy', 'Turner', 'tt@gmailcom');
INSERT INTO footpal_user VALUES (15, 'Drew', 'OnPage', 'haha@gmailcom');
INSERT INTO footpal_user VALUES (16, 'Joel', 'Crowe', 'thejoel@gmailcom');
INSERT INTO footpal_user VALUES (17, 'Yeti', 'Netty', 'yeti@gmailcom');
INSERT INTO footpal_user VALUES (18, 'Harold', 'Hart', 'thehartyboy@gmailcom');
INSERT INTO footpal_user VALUES (19, 'Beth', 'Brock', 'rockthebrock8382@gmailcom');
INSERT INTO footpal_user VALUES (20, 'Lionel', 'Messi', 'messi10@gmailcom');

-- Init Venues

INSERT INTO venue VALUES (1, 'Shankill Sports Arena', '123 Shankill Street', 'BT87 3QW', 'Belfast', '07824712319', 'ssa@booking.com');
INSERT INTO venue VALUES (2, 'Lurgan Liars Soccer Dome', '95 William Street', 'BT66 8RT', 'Belfast', '07824712319', 'lurgansoccerdome@info.co.uk');

-- Init Venue Admins
INSERT INTO venue_admin VALUES (1, 2, 1);
INSERT INTO venue_admin VALUES (2, 6, 2);
INSERT INTO venue_admin VALUES (3, 4, 2);

-- Init Players
INSERT INTO player VALUES (1, 1, null, '07858274821', 'BT34 5RW', 'Craigavon');
INSERT INTO player VALUES (2, 3, 'Blazer', '07858274821', 'BT34 5RW', 'Belfast');
INSERT INTO player VALUES (3, 5, null, '07858274821', 'BT98 6UY', 'Belfast');
INSERT INTO player VALUES (4, 7, null, '07858274821', 'BT34 5RW', 'Lurgan');
INSERT INTO player VALUES (5, 8, 'Rocket', '07858274821', 'BT34 5RW', 'Craigavon');
INSERT INTO player VALUES (6, 9, null, '07858274821', 'BT98 6UY', 'Belfast');
INSERT INTO player VALUES (7, 10, null, '07858274821', 'BT34 5RW', 'Craigavon');
INSERT INTO player VALUES (8, 11, null, '07858274821', 'BT34 5RW', 'Larne');
INSERT INTO player VALUES (9, 12, 'Looner', '07858274821', 'BT98 6UY', 'Lurgan');
INSERT INTO player VALUES (10, 13, null, '07858274821', 'BT34 5RW', 'Lisburn');
INSERT INTO player VALUES (11, 14, 'Top Bins', '07858274821', 'BT34 5RW', 'Craigavon');
INSERT INTO player VALUES (12, 15, null, '07858274821', 'BT98 6UY', 'Belfast');
INSERT INTO player VALUES (13, 16, 'Magic', '07858274821', 'BT98 6UY', 'Lisburn');
INSERT INTO player VALUES (14, 17, null, '07858274821', 'BT34 5RW', 'Craigavon');
INSERT INTO player VALUES (15, 18, null, '07858274821', 'BT34 5RW', 'Lisburn');
INSERT INTO player VALUES (16, 19, null, '07858274821', 'BT98 6UY', 'Belfast');
INSERT INTO player VALUES (17, 20, 'Leo', '07858274821', 'BT98 6UY', 'Lurgan');

-- Init Pitches
INSERT INTO pitch VALUES(1, 1, 'Pitch 1', 10, 30.0);
INSERT INTO pitch VALUES(2, 1, 'Pitch 2', 10, 30.0);
INSERT INTO pitch VALUES(3, 1, 'Pitch 3', 14, 50.0);
INSERT INTO pitch VALUES(4, 2, 'Lennon Arena', 10, 40.0);
INSERT INTO pitch VALUES(5, 2, 'The Big one', 14, 50.0);

-- Init Pitch Time Slots
INSERT INTO pitch_time_slot VALUES(1, 1, 'Monday', '13:00:00', '14:00:00');
INSERT INTO pitch_time_slot VALUES(2, 1, 'Monday', '14:00:00', '15:00:00');
INSERT INTO pitch_time_slot VALUES(3, 1, 'Monday', '15:00:00', '16:00:00');
INSERT INTO pitch_time_slot VALUES(4, 1, 'Monday', '16:00:00', '17:00:00');
INSERT INTO pitch_time_slot VALUES(5, 1, 'Monday', '17:00:00', '18:00:00');
INSERT INTO pitch_time_slot VALUES(6, 1, 'Monday', '18:00:00', '19:00:00');
INSERT INTO pitch_time_slot VALUES(7, 1, 'Monday', '19:00:00', '20:00:00');
INSERT INTO pitch_time_slot VALUES(8, 1, 'Monday', '20:00:00', '21:00:00');
INSERT INTO pitch_time_slot VALUES(9, 1, 'Monday', '21:00:00', '22:00:00');

INSERT INTO pitch_time_slot VALUES(10, 1, 'Tuesday', '13:00:00', '14:00:00');
INSERT INTO pitch_time_slot VALUES(11, 1, 'Tuesday', '14:00:00', '15:00:00');
INSERT INTO pitch_time_slot VALUES(12, 1, 'Tuesday', '15:00:00', '16:00:00');
INSERT INTO pitch_time_slot VALUES(13, 1, 'Tuesday', '16:00:00', '17:00:00');
INSERT INTO pitch_time_slot VALUES(14, 1, 'Tuesday', '17:00:00', '18:00:00');
INSERT INTO pitch_time_slot VALUES(15, 1, 'Tuesday', '18:00:00', '19:00:00');
INSERT INTO pitch_time_slot VALUES(16, 1, 'Tuesday', '19:00:00', '20:00:00');
INSERT INTO pitch_time_slot VALUES(17, 1, 'Tuesday', '20:00:00', '21:00:00');
INSERT INTO pitch_time_slot VALUES(18, 1, 'Tuesday', '21:00:00', '22:00:00');

INSERT INTO pitch_time_slot VALUES(19, 1, 'Wednesday', '13:00:00', '14:00:00');
INSERT INTO pitch_time_slot VALUES(20, 1, 'Wednesday', '14:00:00', '15:00:00');
INSERT INTO pitch_time_slot VALUES(21, 1, 'Wednesday', '15:00:00', '16:00:00');
INSERT INTO pitch_time_slot VALUES(22, 1, 'Wednesday', '16:00:00', '17:00:00');
INSERT INTO pitch_time_slot VALUES(23, 1, 'Wednesday', '17:00:00', '18:00:00');
INSERT INTO pitch_time_slot VALUES(24, 1, 'Wednesday', '18:00:00', '19:00:00');
INSERT INTO pitch_time_slot VALUES(25, 1, 'Wednesday', '19:00:00', '20:00:00');
INSERT INTO pitch_time_slot VALUES(26, 1, 'Wednesday', '20:00:00', '21:00:00');
INSERT INTO pitch_time_slot VALUES(27, 1, 'Wednesday', '21:00:00', '22:00:00');

INSERT INTO pitch_time_slot VALUES(28, 1, 'Thursday', '13:00:00', '14:00:00');
INSERT INTO pitch_time_slot VALUES(29, 1, 'Thursday', '14:00:00', '15:00:00');
INSERT INTO pitch_time_slot VALUES(30, 1, 'Thursday', '15:00:00', '16:00:00');
INSERT INTO pitch_time_slot VALUES(31, 1, 'Thursday', '16:00:00', '17:00:00');
INSERT INTO pitch_time_slot VALUES(32, 1, 'Thursday', '17:00:00', '18:00:00');
INSERT INTO pitch_time_slot VALUES(33, 1, 'Thursday', '18:00:00', '19:00:00');
INSERT INTO pitch_time_slot VALUES(34, 1, 'Thursday', '19:00:00', '20:00:00');
INSERT INTO pitch_time_slot VALUES(35, 1, 'Thursday', '20:00:00', '21:00:00');
INSERT INTO pitch_time_slot VALUES(36, 1, 'Thursday', '21:00:00', '22:00:00');

INSERT INTO pitch_time_slot VALUES(37, 1, 'Friday', '13:00:00', '14:00:00');
INSERT INTO pitch_time_slot VALUES(38, 1, 'Friday', '14:00:00', '15:00:00');
INSERT INTO pitch_time_slot VALUES(39, 1, 'Friday', '15:00:00', '16:00:00');
INSERT INTO pitch_time_slot VALUES(40, 1, 'Friday', '16:00:00', '17:00:00');
INSERT INTO pitch_time_slot VALUES(41, 1, 'Friday', '17:00:00', '18:00:00');
INSERT INTO pitch_time_slot VALUES(42, 1, 'Friday', '18:00:00', '19:00:00');
INSERT INTO pitch_time_slot VALUES(43, 1, 'Friday', '19:00:00', '20:00:00');
INSERT INTO pitch_time_slot VALUES(44, 1, 'Friday', '20:00:00', '21:00:00');
INSERT INTO pitch_time_slot VALUES(45, 1, 'Friday', '21:00:00', '22:00:00');

INSERT INTO pitch_time_slot VALUES(46, 1, 'Saturday', '09:00:00', '10:00:00');
INSERT INTO pitch_time_slot VALUES(47, 1, 'Saturday', '10:00:00', '11:00:00');
INSERT INTO pitch_time_slot VALUES(48, 1, 'Saturday', '11:00:00', '12:00:00');
INSERT INTO pitch_time_slot VALUES(49, 1, 'Saturday', '12:00:00', '13:00:00');
INSERT INTO pitch_time_slot VALUES(50, 1, 'Saturday', '13:00:00', '14:00:00');
INSERT INTO pitch_time_slot VALUES(51, 1, 'Saturday', '14:00:00', '15:00:00');
INSERT INTO pitch_time_slot VALUES(52, 1, 'Saturday', '15:00:00', '16:00:00');
INSERT INTO pitch_time_slot VALUES(53, 1, 'Saturday', '16:00:00', '17:00:00');

INSERT INTO pitch_time_slot VALUES(54, 4, 'Sunday', '09:00:00', '10:00:00');
INSERT INTO pitch_time_slot VALUES(55, 4, 'Sunday', '10:00:00', '11:00:00');
INSERT INTO pitch_time_slot VALUES(56, 4, 'Sunday', '11:00:00', '12:00:00');
INSERT INTO pitch_time_slot VALUES(57, 4, 'Sunday', '12:00:00', '13:00:00');
INSERT INTO pitch_time_slot VALUES(58, 4, 'Sunday', '13:00:00', '14:00:00');
INSERT INTO pitch_time_slot VALUES(59, 4, 'Sunday', '14:00:00', '15:00:00');
INSERT INTO pitch_time_slot VALUES(60, 4, 'Sunday', '15:00:00', '16:00:00');
INSERT INTO pitch_time_slot VALUES(61, 4, 'Sunday', '16:00:00', '17:00:00');

INSERT INTO pitch_time_slot VALUES(62, 4, 'Monday', '13:00:00', '14:00:00');
INSERT INTO pitch_time_slot VALUES(63, 4, 'Monday', '14:00:00', '15:00:00');
INSERT INTO pitch_time_slot VALUES(64, 4, 'Monday', '15:00:00', '16:00:00');
INSERT INTO pitch_time_slot VALUES(65, 4, 'Monday', '16:00:00', '17:00:00');
INSERT INTO pitch_time_slot VALUES(66, 4, 'Monday', '17:00:00', '18:00:00');
INSERT INTO pitch_time_slot VALUES(67, 4, 'Monday', '18:00:00', '19:00:00');
INSERT INTO pitch_time_slot VALUES(68, 4, 'Monday', '19:00:00', '20:00:00');
INSERT INTO pitch_time_slot VALUES(69, 4, 'Monday', '20:00:00', '21:00:00');
INSERT INTO pitch_time_slot VALUES(70, 4, 'Monday', '21:00:00', '22:00:00');

INSERT INTO pitch_time_slot VALUES(71, 4, 'Tuesday', '13:00:00', '14:00:00');
INSERT INTO pitch_time_slot VALUES(72, 4, 'Tuesday', '14:00:00', '15:00:00');
INSERT INTO pitch_time_slot VALUES(73, 4, 'Tuesday', '15:00:00', '16:00:00');
INSERT INTO pitch_time_slot VALUES(74, 4, 'Tuesday', '16:00:00', '17:00:00');
INSERT INTO pitch_time_slot VALUES(75, 4, 'Tuesday', '17:00:00', '18:00:00');
INSERT INTO pitch_time_slot VALUES(76, 4, 'Tuesday', '18:00:00', '19:00:00');
INSERT INTO pitch_time_slot VALUES(77, 4, 'Tuesday', '19:00:00', '20:00:00');
INSERT INTO pitch_time_slot VALUES(78, 4, 'Tuesday', '20:00:00', '21:00:00');
INSERT INTO pitch_time_slot VALUES(79, 4, 'Tuesday', '21:00:00', '22:00:00');

INSERT INTO pitch_time_slot VALUES(80, 4, 'Wednesday', '13:00:00', '14:00:00');
INSERT INTO pitch_time_slot VALUES(81, 4, 'Wednesday', '14:00:00', '15:00:00');
INSERT INTO pitch_time_slot VALUES(82, 4, 'Wednesday', '15:00:00', '16:00:00');
INSERT INTO pitch_time_slot VALUES(83, 4, 'Wednesday', '16:00:00', '17:00:00');
INSERT INTO pitch_time_slot VALUES(84, 4, 'Wednesday', '17:00:00', '18:00:00');
INSERT INTO pitch_time_slot VALUES(85, 4, 'Wednesday', '18:00:00', '19:00:00');
INSERT INTO pitch_time_slot VALUES(86, 4, 'Wednesday', '19:00:00', '20:00:00');
INSERT INTO pitch_time_slot VALUES(87, 4, 'Wednesday', '20:00:00', '21:00:00');
INSERT INTO pitch_time_slot VALUES(88, 4, 'Wednesday', '21:00:00', '22:00:00');

INSERT INTO pitch_time_slot VALUES(89, 4, 'Thursday', '13:00:00', '14:00:00');
INSERT INTO pitch_time_slot VALUES(90, 4, 'Thursday', '14:00:00', '15:00:00');
INSERT INTO pitch_time_slot VALUES(91, 4, 'Thursday', '15:00:00', '16:00:00');
INSERT INTO pitch_time_slot VALUES(92, 4, 'Thursday', '16:00:00', '17:00:00');
INSERT INTO pitch_time_slot VALUES(93, 4, 'Thursday', '17:00:00', '18:00:00');
INSERT INTO pitch_time_slot VALUES(94, 4, 'Thursday', '18:00:00', '19:00:00');
INSERT INTO pitch_time_slot VALUES(95, 4, 'Thursday', '19:00:00', '20:00:00');
INSERT INTO pitch_time_slot VALUES(96, 4, 'Thursday', '20:00:00', '21:00:00');
INSERT INTO pitch_time_slot VALUES(97, 4, 'Thursday', '21:00:00', '22:00:00');

INSERT INTO pitch_time_slot VALUES(98, 4, 'Friday', '13:00:00', '14:00:00');
INSERT INTO pitch_time_slot VALUES(99, 4, 'Friday', '14:00:00', '15:00:00');
INSERT INTO pitch_time_slot VALUES(100, 4, 'Friday', '15:00:00', '16:00:00');
INSERT INTO pitch_time_slot VALUES(101, 4, 'Friday', '16:00:00', '17:00:00');
INSERT INTO pitch_time_slot VALUES(102, 4, 'Friday', '17:00:00', '18:00:00');
INSERT INTO pitch_time_slot VALUES(103, 4, 'Friday', '18:00:00', '19:00:00');
INSERT INTO pitch_time_slot VALUES(104, 4, 'Friday', '19:00:00', '20:00:00');
INSERT INTO pitch_time_slot VALUES(105, 4, 'Friday', '20:00:00', '21:00:00');
INSERT INTO pitch_time_slot VALUES(106, 4, 'Friday', '21:00:00', '22:00:00');

INSERT INTO pitch_time_slot VALUES(107, 4, 'Saturday', '09:00:00', '10:00:00');
INSERT INTO pitch_time_slot VALUES(108, 4, 'Saturday', '10:00:00', '11:00:00');
INSERT INTO pitch_time_slot VALUES(109, 4, 'Saturday', '11:00:00', '12:00:00');
INSERT INTO pitch_time_slot VALUES(110, 4, 'Saturday', '12:00:00', '13:00:00');
INSERT INTO pitch_time_slot VALUES(111, 4, 'Saturday', '13:00:00', '14:00:00');
INSERT INTO pitch_time_slot VALUES(112, 4, 'Saturday', '14:00:00', '15:00:00');
INSERT INTO pitch_time_slot VALUES(113, 4, 'Saturday', '15:00:00', '16:00:00');
INSERT INTO pitch_time_slot VALUES(114, 4, 'Saturday', '16:00:00', '17:00:00');

INSERT INTO pitch_time_slot VALUES(115, 4, 'Sunday', '09:00:00', '10:00:00');
INSERT INTO pitch_time_slot VALUES(116, 4, 'Sunday', '10:00:00', '11:00:00');
INSERT INTO pitch_time_slot VALUES(117, 4, 'Sunday', '11:00:00', '12:00:00');
INSERT INTO pitch_time_slot VALUES(118, 4, 'Sunday', '12:00:00', '13:00:00');
INSERT INTO pitch_time_slot VALUES(119, 4, 'Sunday', '13:00:00', '14:00:00');
INSERT INTO pitch_time_slot VALUES(120, 4, 'Sunday', '14:00:00', '15:00:00');
INSERT INTO pitch_time_slot VALUES(121, 4, 'Sunday', '15:00:00', '16:00:00');
INSERT INTO pitch_time_slot VALUES(122, 4, 'Sunday', '16:00:00', '17:00:00');


-- Init Squads
INSERT INTO squad VALUES(1, 'Leos 5 a side', 'Belfast');
INSERT INTO squad VALUES(2, 'CR7 5 a side', 'Belfast');

-- Init Squad Players
INSERT INTO squad_player VALUES(1, 1, 7, 'player', 1);
INSERT INTO squad_player VALUES(2, 1, 11, 'player', 1);
INSERT INTO squad_player VALUES(3, 1, 5, 'player', 1);
INSERT INTO squad_player VALUES(4, 1, 15, 'player', 1);
INSERT INTO squad_player VALUES(5, 1, 1, 'player', 1);
INSERT INTO squad_player VALUES(6, 1, 14, 'player', 1);
INSERT INTO squad_player VALUES(7, 1, 10, 'player', 1);
INSERT INTO squad_player VALUES(8, 1, 17, 'player', 1);
INSERT INTO squad_player VALUES(9, 1, 8, 'player', 1);
INSERT INTO squad_player VALUES(10, 1, 6, 'admin', 1);

INSERT INTO squad_player VALUES(11, 2, 9, 'player', 1);
INSERT INTO squad_player VALUES(12, 2, 3, 'player', 1);
INSERT INTO squad_player VALUES(13, 2, 12, 'player', 1);
INSERT INTO squad_player VALUES(14, 2, 15, 'player', 1);
INSERT INTO squad_player VALUES(15, 2, 16, 'player', 1);
INSERT INTO squad_player VALUES(16, 2, 17, 'player', 1);
INSERT INTO squad_player VALUES(17, 2, 16, 'player', 1);
INSERT INTO squad_player VALUES(18, 2, 17, 'admin', 1);
INSERT INTO squad_player VALUES(19, 2, 13, 'player', 1);

-- Init Bookings
INSERT INTO booking VALUES(1, 1, 17, current_timestamp, current_timestamp);
INSERT INTO booking VALUES(2, 4, 6, current_timestamp, current_timestamp);

-- Init Pitch Slots
INSERT INTO pitch_slot VALUES(1, 1, 70, DATE '2022-04-20', 1);
INSERT INTO pitch_slot VALUES(2, 1, 70, DATE '2022-04-27', 1);
INSERT INTO pitch_slot VALUES(3, 1, 70, DATE '2022-05-04', 1);
INSERT INTO pitch_slot VALUES(4, 2, 52, DATE '2022-04-25', 4);

-- Init Match
INSERT INTO match VALUES(1, 1, 2, 1, DATE '2022-04-20', 30.0, false, current_timestamp, current_timestamp);
INSERT INTO match VALUES(2, 1, 2, 4, DATE '2022-04-27', 30.0, false, current_timestamp, current_timestamp);
INSERT INTO match VALUES(3, 1, 2, 4, DATE '2022-05-04', 30.0, false, current_timestamp, current_timestamp);

-- Init Match Player
INSERT INTO match_player VALUES(1, 1, 4, 3.0, 1);
INSERT INTO match_player VALUES(2, 1, 8, 3.0, 1);
INSERT INTO match_player VALUES(3, 1, 3, 3.0, 1);
INSERT INTO match_player VALUES(4, 1, 15, 3.0, 2);
INSERT INTO match_player VALUES(5, 1, 1, 3.0, 1);
INSERT INTO match_player VALUES(6, 1, 11, 3.0, 1);
INSERT INTO match_player VALUES(7, 1, 7, 3.0, 2);
INSERT INTO match_player VALUES(8, 1, 17, 3.0, 2);
INSERT INTO match_player VALUES(9, 1, 5, 3.0, 1);
INSERT INTO match_player VALUES(10, 1, 10, 3.0, 1);