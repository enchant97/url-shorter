CREATE TABLE shorts (
    id bigint primary key generated always as identity,
    created_at timestamptz not null default now(),
    updated_at timestamptz not null default now(),
    deleted_at timestamptz,
    slug text not null unique,
    target_url text not null
);
