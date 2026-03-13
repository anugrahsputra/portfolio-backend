-- name: CreateLanguage :one
insert into languages (
    profile_id, language, proficiency
) values (
    $1, $2, $3
) 
on conflict (profile_id, language) do update set
    proficiency = excluded.proficiency
returning *;

-- name: GetLanguage :one
select * from languages
where id = $1 limit 1;

-- name: ListLanguages :many
select * from languages
where profile_id = $1
order by language;

-- name: UpdateLanguage :one
update languages
set
    language = $2,
    proficiency = $3
where id = $1
returning *;

-- name: DeleteLanguage :exec
delete from languages
where id = $1;
