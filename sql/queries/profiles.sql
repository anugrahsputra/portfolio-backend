-- name: CreateProfile :one
insert into profiles (
    name, about, address, email, phone
) values (
    $1, $2, $3, $4, $5
) returning *;

-- name: GetProfile :one
SELECT
    p.id,
    p.name,
    p.about,
    p.address,
    p.email,
    p.phone,
    COALESCE(
        json_agg(
            json_build_object(
                'id', pu.id,
                'label', pu.label,
                'url', pu.url
            )
        ) FILTER (WHERE pu.id IS NOT NULL),
        '[]'
    ) AS urls
FROM profiles p
LEFT JOIN profile_urls pu ON pu.profile_id = p.id
WHERE p.id = $1
GROUP BY p.id;

-- name: UpdateProfile :exec
update profiles
set
    name = COALESCE($2, name),
    about = COALESCE($3, about),
    address = COALESCE($4, address),
    email = COALESCE($5, email),
    phone = COALESCE($6, phone)
where id = $1;

-- name: DeleteProfile :exec
delete from profiles
where id = $1;
