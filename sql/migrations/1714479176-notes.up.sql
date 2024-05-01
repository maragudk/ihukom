create table notes (
  id text primary key default ('n_' || lower(hex(randomblob(16)))),
  created text not null default (strftime('%Y-%m-%dT%H:%M:%fZ')),
  updated text not null default (strftime('%Y-%m-%dT%H:%M:%fZ')),
  content text not null default ''
) strict;

create trigger notes_updated_timestamp after update on notes begin
  update notes set updated = strftime('%Y-%m-%dT%H:%M:%fZ') where id = old.id;
end;

create table speakers (
  id text primary key default ('s_' || lower(hex(randomblob(16)))),
  created text not null default (strftime('%Y-%m-%dT%H:%M:%fZ')),
  updated text not null default (strftime('%Y-%m-%dT%H:%M:%fZ')),
  name text not null default ''
) strict;

create trigger speakers_updated_timestamp after update on speakers begin
  update speakers set updated = strftime('%Y-%m-%dT%H:%M:%fZ') where id = old.id;
end;

create table conversations (
  id text primary key default ('c_' || lower(hex(randomblob(16)))),
  created text not null default (strftime('%Y-%m-%dT%H:%M:%fZ')),
  updated text not null default (strftime('%Y-%m-%dT%H:%M:%fZ'))
) strict;

create trigger conversations_updated_timestamp after update on conversations begin
  update conversations set updated = strftime('%Y-%m-%dT%H:%M:%fZ') where id = old.id;
end;

create table turns (
  id text primary key default ('t_' || lower(hex(randomblob(16)))),
  created text not null default (strftime('%Y-%m-%dT%H:%M:%fZ')),
  updated text not null default (strftime('%Y-%m-%dT%H:%M:%fZ')),
  conversationID text not null references conversations (id) on delete cascade,
  speakerID text not null references speakers (id),
  content text not null default ''
) strict;

create trigger turns_updated_timestamp after update on turns begin
  update turns set updated = strftime('%Y-%m-%dT%H:%M:%fZ') where id = old.id;
end;
