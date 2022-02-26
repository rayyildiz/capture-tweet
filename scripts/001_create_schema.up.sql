CREATE TABLE users
(
    id                text      not null primary key,
    username          text      not null,
    screen_name       text      not null,
    bio               text,
    profile_image_url text,
    registered_at     timestamp not null,
    created_at        timestamp default now(),
    updated_at        timestamp default now()
);

create unique index users_username on users (username);

CREATE TABLE tweets
(
    id                text      not null primary key,
    full_text         text      not null,
    capture_url       text,
    capture_thumb_url text,
    lang              text      not null,
    favorite_count    int       not null default 0,
    retweet_count     int       not null default 0,
    resources         jsonb,
    author_id         text      not null references users (id),
    posted_at         timestamp not null,
    created_at        timestamp          default now(),
    updated_at        timestamp          default now()
);

create index tweets_author on tweets (author_id);


CREATE TABLE contact_us
(
    id         text primary key,
    email      text,
    full_name  text,
    message    text,
    created_at timestamp
);