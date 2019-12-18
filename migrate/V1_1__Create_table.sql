create table almanac (
 id text,
"user" text,
 title text,
 body text,
 starttime timestamp,
 endtime timestamp
);
create index user_idx on almanac ("user");