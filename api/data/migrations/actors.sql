drop if exists actors;
create table if not exists actors (
	actor_id BIGSERIAL PRIMARY KEY,
	first_name VARCHAR(50) NOT NULL,
	last_name VARCHAR(50) NOT NULL,
	gender VARCHAR(50) NOT NULL,
	dob DATE NOT NULL,
	email VARCHAR(50) NOT NULL
);
insert into actors (first_name, last_name, gender, dob, email) values ('Christa', 'Hatchette', 'Female', '1963-06-28 23:43:31', 'chatchette0@geocities.jp');
insert into actors (first_name, last_name, gender, dob, email) values ('Anett', 'Cattle', 'Female', '1987-06-17 13:50:22', 'acattle1@vimeo.com');
insert into actors (first_name, last_name, gender, dob, email) values ('Robby', 'Nisuis', 'Male', '1980-06-02 08:29:27', 'rnisuis2@wikimedia.org');
insert into actors (first_name, last_name, gender, dob, email) values ('Chandal', 'Kunert', 'Female', '1989-04-03 19:49:42', 'ckunert3@nhs.uk');
insert into actors (first_name, last_name, gender, dob, email) values ('Nobe', 'Bicheno', 'Male', '1991-03-26 18:03:07', 'nbicheno4@theglobeandmail.com');
insert into actors (first_name, last_name, gender, dob, email) values ('Anson', 'Rollitt', 'Male', '1959-02-23 22:06:36', 'arollitt5@naver.com');
insert into actors (first_name, last_name, gender, dob, email) values ('Morrie', 'Gianninotti', 'Male', '1958-12-22 00:21:11', 'mgianninotti6@vk.com');
insert into actors (first_name, last_name, gender, dob, email) values ('Patsy', 'Tookill', 'Male', '1990-12-18 22:58:19', 'ptookill7@toplist.cz');
insert into actors (first_name, last_name, gender, dob, email) values ('Ted', 'Recke', 'Female', '1969-06-26 13:23:57', 'trecke8@sfgate.com');
insert into actors (first_name, last_name, gender, dob, email) values ('Burke', 'Castiglione', 'Male', '1966-04-05 12:17:52', 'bcastiglione9@netvibes.com');
insert into actors (first_name, last_name, gender, dob, email) values ('Tome', 'Pinel', 'Male', '1990-01-22 22:18:58', 'tpinela@prnewswire.com');
insert into actors (first_name, last_name, gender, dob, email) values ('Basilius', 'Dowbekin', 'Male', '1967-12-18 13:54:38', 'bdowbekinb@printfriendly.com');
insert into actors (first_name, last_name, gender, dob, email) values ('Alexa', 'Rubbert', 'Female', '1994-08-22 20:04:23', 'arubbertc@java.com');
insert into actors (first_name, last_name, gender, dob, email) values ('Rhianon', 'Webborn', 'Female', '1999-07-31 21:17:49', 'rwebbornd@jimdo.com');
insert into actors (first_name, last_name, gender, dob, email) values ('Huntlee', 'Howgego', 'Male', '1954-03-14 07:38:49', 'hhowgegoe@reference.com');
insert into actors (first_name, last_name, gender, dob, email) values ('Winnie', 'Windsor', 'Female', '1975-11-03 23:36:19', 'wwindsorf@bandcamp.com');
insert into actors (first_name, last_name, gender, dob, email) values ('Tate', 'Wedgwood', 'Male', '1996-11-30 04:27:37', 'twedgwoodg@fastcompany.com');
insert into actors (first_name, last_name, gender, dob, email) values ('Becka', 'Whimper', 'Female', '1998-01-14 02:58:27', 'bwhimperh@about.com');
insert into actors (first_name, last_name, gender, dob, email) values ('Perice', 'Triggel', 'Male', '1964-10-19 08:55:01', 'ptriggeli@mit.edu');
insert into actors (first_name, last_name, gender, dob, email) values ('Sibel', 'Muller', 'Female', '1981-01-11 19:39:02', 'smullerj@cbsnews.com');
insert into actors (first_name, last_name, gender, dob, email) values ('Verne', 'Woltman', 'Male', '1978-02-06 06:16:35', 'vwoltmank@shinystat.com');
insert into actors (first_name, last_name, gender, dob, email) values ('Tabina', 'Houdmont', 'Female', '1950-11-27 05:31:34', 'thoudmontl@google.com');
insert into actors (first_name, last_name, gender, dob, email) values ('Cleopatra', 'Bausmann', 'Female', '1950-08-15 02:50:04', 'cbausmannm@youtube.com');
insert into actors (first_name, last_name, gender, dob, email) values ('Charlton', 'Tiplady', 'Male', '1951-10-08 05:06:57', 'ctipladyn@cafepress.com');
insert into actors (first_name, last_name, gender, dob, email) values ('Norrie', 'Hartas', 'Male', '1963-05-03 19:11:23', 'nhartaso@flickr.com');
insert into actors (first_name, last_name, gender, dob, email) values ('Dee dee', 'Orrick', 'Female', '1977-03-20 08:30:31', 'dorrickp@gnu.org');
insert into actors (first_name, last_name, gender, dob, email) values ('Donnie', 'Seally', 'Female', '1973-01-03 15:50:27', 'dseallyq@oakley.com');
insert into actors (first_name, last_name, gender, dob, email) values ('Jim', 'Pundy', 'Male', '1973-10-09 08:38:44', 'jpundyr@stanford.edu');
insert into actors (first_name, last_name, gender, dob, email) values ('Gerrard', 'Firmin', 'Male', '1963-05-13 20:09:04', 'gfirmins@mysql.com');
insert into actors (first_name, last_name, gender, dob, email) values ('Gus', 'Feirn', 'Male', '1972-10-10 17:41:33', 'gfeirnt@skyrock.com');
insert into actors (first_name, last_name, gender, dob, email) values ('Ernestus', 'Casement', 'Male', '1982-12-25 06:01:01', 'ecasementu@samsung.com');
insert into actors (first_name, last_name, gender, dob, email) values ('Pen', 'Meharry', 'Male', '1978-05-10 09:50:09', 'pmeharryv@imgur.com');
insert into actors (first_name, last_name, gender, dob, email) values ('Vic', 'Oxbrow', 'Male', '1986-07-18 21:28:43', 'voxbroww@mapy.cz');
insert into actors (first_name, last_name, gender, dob, email) values ('Terri', 'Arnoldi', 'Male', '1989-05-12 11:05:02', 'tarnoldix@senate.gov');
insert into actors (first_name, last_name, gender, dob, email) values ('Pippo', 'Gledstane', 'Male', '1996-06-16 02:44:46', 'pgledstaney@mayoclinic.com');
insert into actors (first_name, last_name, gender, dob, email) values ('Marcy', 'Burtwell', 'Female', '1964-06-17 08:38:44', 'mburtwellz@europa.eu');
insert into actors (first_name, last_name, gender, dob, email) values ('Kristian', 'Forten', 'Male', '1980-03-28 00:49:51', 'kforten10@amazon.de');
insert into actors (first_name, last_name, gender, dob, email) values ('Henderson', 'Balfour', 'Male', '1958-09-22 10:17:21', 'hbalfour11@mapquest.com');
insert into actors (first_name, last_name, gender, dob, email) values ('See', 'Fouldes', 'Male', '1985-11-26 08:11:41', 'sfouldes12@4shared.com');
insert into actors (first_name, last_name, gender, dob, email) values ('Pippy', 'Quince', 'Female', '1977-12-28 22:12:43', 'pquince13@comcast.net');
insert into actors (first_name, last_name, gender, dob, email) values ('Weider', 'Astbury', 'Male', '1984-07-15 22:14:32', 'wastbury14@thetimes.co.uk');
insert into actors (first_name, last_name, gender, dob, email) values ('Hayden', 'Whitley', 'Male', '1964-03-19 15:30:59', 'hwhitley15@dagondesign.com');
insert into actors (first_name, last_name, gender, dob, email) values ('Jasper', 'Vautre', 'Male', '1997-03-11 01:43:48', 'jvautre16@lulu.com');
insert into actors (first_name, last_name, gender, dob, email) values ('Sonnnie', 'Grossier', 'Female', '1988-05-22 21:21:04', 'sgrossier17@wsj.com');
insert into actors (first_name, last_name, gender, dob, email) values ('Enrico', 'Circuit', 'Male', '1989-01-14 06:28:47', 'ecircuit18@un.org');
insert into actors (first_name, last_name, gender, dob, email) values ('Zacharias', 'Pedrick', 'Male', '1962-01-14 23:56:41', 'zpedrick19@blogtalkradio.com');
insert into actors (first_name, last_name, gender, dob, email) values ('Carrissa', 'Bamforth', 'Female', '1962-11-07 13:08:06', 'cbamforth1a@mtv.com');
insert into actors (first_name, last_name, gender, dob, email) values ('Heath', 'Manthorpe', 'Male', '1995-02-26 00:12:37', 'hmanthorpe1b@earthlink.net');
insert into actors (first_name, last_name, gender, dob, email) values ('Carmencita', 'Dinzey', 'Female', '1981-11-14 04:20:45', 'cdinzey1c@sfgate.com');
insert into actors (first_name, last_name, gender, dob, email) values ('Waldemar', 'Mander', 'Male', '1963-02-26 16:17:36', 'wmander1d@dot.gov');
insert into actors (first_name, last_name, gender, dob, email) values ('Jethro', 'Faiers', 'Male', '1997-01-12 21:38:18', 'jfaiers1e@redcross.org');
insert into actors (first_name, last_name, gender, dob, email) values ('Renaldo', 'Luis', 'Male', '1954-04-05 06:23:53', 'rluis1f@t-online.de');
insert into actors (first_name, last_name, gender, dob, email) values ('Lucita', 'MacAvddy', 'Female', '1971-01-08 05:14:05', 'lmacavddy1g@techcrunch.com');
insert into actors (first_name, last_name, gender, dob, email) values ('Darda', 'Skupinski', 'Female', '1999-06-09 08:11:50', 'dskupinski1h@timesonline.co.uk');
insert into actors (first_name, last_name, gender, dob, email) values ('Lennard', 'Shurrock', 'Male', '1992-03-15 00:09:31', 'lshurrock1i@liveinternet.ru');
insert into actors (first_name, last_name, gender, dob, email) values ('Shari', 'Gubbin', 'Female', '1991-01-17 20:53:14', 'sgubbin1j@bandcamp.com');
insert into actors (first_name, last_name, gender, dob, email) values ('Prissie', 'Dome', 'Female', '1993-09-02 06:27:59', 'pdome1k@imdb.com');
insert into actors (first_name, last_name, gender, dob, email) values ('Isidro', 'Pelham', 'Male', '1975-05-15 07:08:17', 'ipelham1l@rakuten.co.jp');
insert into actors (first_name, last_name, gender, dob, email) values ('Cammie', 'Di Boldi', 'Female', '1994-09-01 12:59:47', 'cdiboldi1m@wired.com');
insert into actors (first_name, last_name, gender, dob, email) values ('Kellsie', 'Craiker', 'Female', '1984-02-19 19:54:17', 'kcraiker1n@canalblog.com');
insert into actors (first_name, last_name, gender, dob, email) values ('Putnam', 'Dufty', 'Male', '1977-08-13 06:34:33', 'pdufty1o@weebly.com');
insert into actors (first_name, last_name, gender, dob, email) values ('Arthur', 'Rowbrey', 'Male', '1963-08-12 16:26:20', 'arowbrey1p@walmart.com');
insert into actors (first_name, last_name, gender, dob, email) values ('Garrott', 'Dedenham', 'Male', '1981-06-08 16:09:14', 'gdedenham1q@craigslist.org');
insert into actors (first_name, last_name, gender, dob, email) values ('Myron', 'Gumm', 'Male', '1954-12-29 00:43:07', 'mgumm1r@typepad.com');
insert into actors (first_name, last_name, gender, dob, email) values ('Mae', 'Stagge', 'Female', '1979-06-06 11:51:49', 'mstagge1s@php.net');
insert into actors (first_name, last_name, gender, dob, email) values ('Dom', 'Mynett', 'Male', '1953-04-06 04:39:53', 'dmynett1t@archive.org');
insert into actors (first_name, last_name, gender, dob, email) values ('Melita', 'Lawee', 'Female', '1968-03-14 08:05:58', 'mlawee1u@ocn.ne.jp');
insert into actors (first_name, last_name, gender, dob, email) values ('Daniel', 'Burnet', 'Male', '1966-03-28 17:24:22', 'dburnet1v@free.fr');
insert into actors (first_name, last_name, gender, dob, email) values ('Sandye', 'Piddlesden', 'Female', '1960-11-19 01:48:38', 'spiddlesden1w@freewebs.com');
insert into actors (first_name, last_name, gender, dob, email) values ('Darn', 'Christoffe', 'Male', '1954-06-28 15:07:21', 'dchristoffe1x@blog.com');
insert into actors (first_name, last_name, gender, dob, email) values ('Bucky', 'Twining', 'Male', '1951-02-24 20:40:33', 'btwining1y@free.fr');
insert into actors (first_name, last_name, gender, dob, email) values ('Leonora', 'Hainey`', 'Female', '1960-08-22 21:34:23', 'lhainey1z@weather.com');
insert into actors (first_name, last_name, gender, dob, email) values ('Thibaut', 'Fabler', 'Male', '1985-06-23 13:37:28', 'tfabler20@shop-pro.jp');
insert into actors (first_name, last_name, gender, dob, email) values ('Clair', 'Snewin', 'Male', '1996-04-27 14:13:34', 'csnewin21@independent.co.uk');
insert into actors (first_name, last_name, gender, dob, email) values ('Baird', 'Kingsmill', 'Male', '1991-03-03 08:23:11', 'bkingsmill22@newyorker.com');
insert into actors (first_name, last_name, gender, dob, email) values ('Maurene', 'Blincko', 'Female', '1960-02-17 07:26:57', 'mblincko23@mtv.com');
insert into actors (first_name, last_name, gender, dob, email) values ('Lilith', 'Baford', 'Female', '1959-03-06 07:11:47', 'lbaford24@skyrock.com');
insert into actors (first_name, last_name, gender, dob, email) values ('Kenton', 'Carwithim', 'Male', '1970-03-15 07:42:16', 'kcarwithim25@tripadvisor.com');
insert into actors (first_name, last_name, gender, dob, email) values ('Hyacinth', 'MacInherney', 'Female', '1970-10-22 11:49:47', 'hmacinherney26@wired.com');
insert into actors (first_name, last_name, gender, dob, email) values ('Jennee', 'Minet', 'Female', '1977-11-03 05:07:53', 'jminet27@edublogs.org');
insert into actors (first_name, last_name, gender, dob, email) values ('Melinda', 'Moy', 'Female', '1958-07-07 02:03:14', 'mmoy28@yolasite.com');
insert into actors (first_name, last_name, gender, dob, email) values ('Sonnie', 'Nussey', 'Male', '1972-06-18 00:18:39', 'snussey29@cmu.edu');
insert into actors (first_name, last_name, gender, dob, email) values ('Clemens', 'Hanes', 'Male', '1969-09-25 19:38:22', 'chanes2a@accuweather.com');
insert into actors (first_name, last_name, gender, dob, email) values ('Idalia', 'Newvell', 'Female', '1979-10-26 14:14:45', 'inewvell2b@businessweek.com');
insert into actors (first_name, last_name, gender, dob, email) values ('Delilah', 'Manwell', 'Female', '1998-08-31 06:05:22', 'dmanwell2c@baidu.com');
insert into actors (first_name, last_name, gender, dob, email) values ('Evangelia', 'Nunson', 'Female', '1977-04-11 17:51:37', 'enunson2d@msu.edu');
insert into actors (first_name, last_name, gender, dob, email) values ('Farlay', 'Demangel', 'Male', '2000-07-06 00:11:18', 'fdemangel2e@360.cn');
insert into actors (first_name, last_name, gender, dob, email) values ('Ario', 'Pipworth', 'Male', '1991-10-06 17:48:05', 'apipworth2f@buzzfeed.com');
insert into actors (first_name, last_name, gender, dob, email) values ('Gates', 'Cortnay', 'Female', '1993-11-30 01:30:43', 'gcortnay2g@blinklist.com');
insert into actors (first_name, last_name, gender, dob, email) values ('Nealson', 'Caillou', 'Male', '1972-10-15 14:08:26', 'ncaillou2h@sfgate.com');
insert into actors (first_name, last_name, gender, dob, email) values ('Clayborn', 'Elcock', 'Male', '1995-11-14 03:28:14', 'celcock2i@jugem.jp');
insert into actors (first_name, last_name, gender, dob, email) values ('Tadio', 'Fritchley', 'Male', '1963-04-05 08:51:32', 'tfritchley2j@yellowpages.com');
insert into actors (first_name, last_name, gender, dob, email) values ('Nyssa', 'MacDougal', 'Female', '1959-02-25 18:04:29', 'nmacdougal2k@seattletimes.com');
insert into actors (first_name, last_name, gender, dob, email) values ('Pru', 'Conkay', 'Female', '1988-12-03 11:01:01', 'pconkay2l@europa.eu');
insert into actors (first_name, last_name, gender, dob, email) values ('Nell', 'Jowle', 'Female', '1982-05-23 10:59:19', 'njowle2m@histats.com');
insert into actors (first_name, last_name, gender, dob, email) values ('Elie', 'Seden', 'Female', '1952-08-13 04:01:20', 'eseden2n@gravatar.com');
insert into actors (first_name, last_name, gender, dob, email) values ('Bendite', 'Lanon', 'Female', '1989-03-09 05:20:47', 'blanon2o@mail.ru');
insert into actors (first_name, last_name, gender, dob, email) values ('Greggory', 'Baughn', 'Male', '1983-01-22 21:16:23', 'gbaughn2p@ycombinator.com');
insert into actors (first_name, last_name, gender, dob, email) values ('Holly-anne', 'Batecok', 'Female', '1979-12-12 20:25:21', 'hbatecok2q@clickbank.net');
insert into actors (first_name, last_name, gender, dob, email) values ('Jenine', 'Ninotti', 'Female', '1967-09-02 21:50:52', 'jninotti2r@usa.gov');
