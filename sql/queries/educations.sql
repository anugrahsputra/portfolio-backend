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

-- name: GetEducationByID :one
select * from educations
where id = $1;

-- name: UpdateEducation :one
update educations
set
    school = COALESCE($2, school),
    degree = COALESCE($3, degree),
    field_of_study = COALESCE($4, field_of_study),
    gpa = COALESCE($5, gpa),
    start_date = COALESCE($6, start_date),
    graduation_date = COALESCE($7, graduation_date)
where id = $1
returning *;

-- name: DeleteEducation :exec
delete from educations
where id = $1;
