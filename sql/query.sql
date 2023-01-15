create table url (
  id              SERIAL primary key,
  code            varchar(10) unique,
  original        text,
  created_at      timestamp not null default current_timestamp
)
