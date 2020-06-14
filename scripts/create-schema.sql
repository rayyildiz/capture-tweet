create table if not exists users
(
    id          varchar(36) not null
        constraint users_pkey
            primary key,
    created_at  timestamp with time zone,
    updated_at  timestamp with time zone,
    deleted_at  timestamp with time zone,
    user_name   text        not null,
    screen_name text
);

create index if not exists idx_users_deleted_at on users (deleted_at);



create table if not exists tweets
(
    id                varchar(36) not null primary key,
    created_at        timestamp with time zone,
    updated_at        timestamp with time zone,
    deleted_at        timestamp with time zone,
    full_text         text        not null,
    capture_url       text        not null,
    capture_thumb_url text        not null,
    lang              text,
    favorite_count    integer,
    retweet_count     integer,
    user_id           varchar(36) references users (id)
);

create index if not exists idx_tweets_deleted_at on tweets (deleted_at);

create table if not exists resources
(
    id         varchar(36) not null primary key,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    url        text        not null,
    width      integer,
    height     integer,
    media_type text,
    tweet_id   varchar(36) references tweets (id)
);

create index if not exists idx_resources_deleted_at on resources (deleted_at);
