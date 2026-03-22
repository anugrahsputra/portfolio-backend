-- name: CreateExperience :one
insert into experiences (
    profile_id, company, position, description, start_date, end_date, is_present
) values (
    $1, $2, $3, $4, $5, $6, $7
) 
on conflict (profile_id, company, position, start_date) do update set
    description = excluded.description,
    end_date = excluded.end_date,
    is_present = excluded.is_present
returning *;


-- name: GetExperiences :many
select * from experiences
where profile_id = $1
order by start_date desc;

-- name: GetExperienceByID :one
select * from experiences
where id = $1;

-- name: UpdateExperience :one
update experiences
set
    company = $2,
    position = $3,
    description = $4,
    start_date = $5,
    end_date = $6,
    is_present = $7
where id = $1
returning *;

-- name: DeleteExperience :exec
delete from experiences
where id = $1;
