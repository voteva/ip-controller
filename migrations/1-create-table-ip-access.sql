create table ip_access
(
	ip varchar(60) not null PRIMARY KEY,
	datetime_first timestamp not null,
	datetime_last timestamp not null
);