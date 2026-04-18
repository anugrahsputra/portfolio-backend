-- name: CreateProfile :one
insert into profiles (
    name, title, about, address, email, phone
) values (
    $1, $2, $3, $4, $5, $6
) returning *;

-- name: GetProfile :one
SELECT
    p.id,
    p.name,
    p.title,
    p.about,
    p.address,
    p.email,
    p.phone,
    (
        SELECT COALESCE(json_agg(urls), '[]'::json)
        FROM (
            SELECT id AS "ID", profile_id AS "ProfileID", label AS "Label", url AS "Url"
            FROM profile_urls
            WHERE profile_id = p.id
        ) urls
    ) AS urls
FROM profiles p
WHERE p.id = $1;

-- name: UpdateProfile :exec
update profiles
set
    name = COALESCE($2, name),
    title = COALESCE($3, title),
    about = COALESCE($4, about),
    address = COALESCE($5, address),
    email = COALESCE($6, email),
    phone = COALESCE($7, phone)
where id = $1;

-- name: DeleteProfile :exec
delete from profiles
where id = $1;
