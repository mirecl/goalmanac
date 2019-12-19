ALTER TABLE almanac ADD notify char(1);
create index notify_idx on almanac (starttime,notify);