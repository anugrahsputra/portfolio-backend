-- name: CreateSkill :one
insert into skills (
    profile_id, tools, technologies, hard_skills, soft_skills
) values (
    $1, $2, $3, $4, $5
) 
on conflict (profile_id) do update set
    tools = excluded.tools,
    technologies = excluded.technologies,
    hard_skills = excluded.hard_skills,
    soft_skills = excluded.soft_skills
returning *;

-- name: GetSkillsByProfile :one
select * from skills
where profile_id = $1 limit 1;

-- name: UpdateSkill :one
update skills
set
    tools = $2,
    technologies = $3,
    hard_skills = $4,
    soft_skills = $5
where id = $1
returning *;

-- name: DeleteSkill :exec
delete from skills
where id = $1;
