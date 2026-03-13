-- name: CreateEducation :one
insert into educations (
    profile_id, school, degree, field_of_study, gpa, start_date, graduation_date
) values (
    $1, $2, $3, $4, $5, $6, $7
) 
on conflict (profile_id, school, degree, start_date) do update set
    field_of_study = excluded.field_of_study,
    gpa = excluded.gpa,
    graduation_date = excluded.graduation_date
returning *;


-- name: GetEducations :many
select * from educations
where profile_id = $1
order by start_date desc;

-- name: UpdateEducation :one
update educations
set
    school = $2,
    degree = $3,
    field_of_study = $4,
    gpa = $5,
    start_date = $6,
    graduation_date = $7
where id = $1
returning *;

-- name: DeleteEducation :exec
delete from educations
where id = $1;
