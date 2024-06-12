CREATE TABLE shorts (
    id bigint primary key generated always as identity,
    created_at timestamp not null,
    updated_at timestamp not null,
    deleted_at timestamp,
    slug varchar(128) not null,
    target_url text not null
);
