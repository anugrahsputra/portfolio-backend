-- +goose Up
alter table projects add column if not exists is_coming_soon boolean not null default false;

-- +goose Down
alter table projects drop column if exists is_coming_soon;
