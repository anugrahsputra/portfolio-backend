-- name: CreateExperience :one
insert into experiences (
    profile_id, company, position, description, start_date, end_date
) values (
    $1, $2, $3, $4, $5, $6
) returning *;


-- name: GetExperiences :many
select * from experiences
where profile_id = $1
order by start_date desc;

-- name: UpdateExperience :one
update experiences
set
    company = $2,
    position = $3,
    description = $4,
    start_date = $5,
    end_date = $6
where id = $1
returning *;

-- name: DeleteExperience :exec
delete from experiences
where id = $1;
