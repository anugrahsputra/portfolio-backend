-- name: CreateProfileURL :one
insert into profile_urls (
    profile_id, label, url
) values (
    $1, $2, $3
) 
on conflict (profile_id, label) do update set
    url = excluded.url
returning *;

-- name: GetProfileURL :one
select * from profile_urls
where id = $1 limit 1;

-- name: ListProfileURLs :many
select * from profile_urls
where profile_id = $1
order by label;

-- name: UpdateProfileURL :one
update profile_urls
set
    label = $2,
    url = $3
where id = $1
returning *;

-- name: DeleteProfileURL :exec
delete from profile_urls
where id = $1;
