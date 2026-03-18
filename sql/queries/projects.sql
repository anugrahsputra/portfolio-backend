-- name: CreateProject :one
insert into projects (
    profile_id,
    title,
    description,
    tech_stacks,
    live_demo_url,
    github_repo_url,
    is_live,
    is_nda,
    is_featured,
    image_url,
    company,
    start_date,
    end_date,
    is_present,
    location
) values (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6,
    $7,
    $8,
    $9,
    $10,
    $11,
    $12,
    $13,
    $14,
    $15
)
on conflict (profile_id, title)
do update set
    description = excluded.description,
    tech_stacks = excluded.tech_stacks,
    live_demo_url = excluded.live_demo_url,
    github_repo_url = excluded.github_repo_url,
    is_live = excluded.is_live,
    is_nda = excluded.is_nda,
    is_featured = excluded.is_featured,
    image_url = excluded.image_url,
    company = excluded.company,
    start_date = excluded.start_date,
    end_date = excluded.end_date,
    is_present = excluded.is_present,
    location = excluded.location
returning *;

-- name: GetProjects :many
select * from projects
where profile_id = $1
order by start_date desc;

-- name: GetProjectByID :one
select *  from projects
where id = $1;

-- name: UpdateProject :one
update projects
set 
    title = COALESCE($2, title),
    description = COALESCE($3, description),
    tech_stacks = COALESCE($4, tech_stacks),
    live_demo_url = COALESCE($5, live_demo_url),
    github_repo_url = COALESCE($6, github_repo_url),
    is_live = COALESCE($7, is_live),
    is_nda = COALESCE($8, is_nda),
    is_featured = COALESCE($9, is_featured),
    image_url = COALESCE($10, image_url),
    company = COALESCE($11, company),
    start_date = COALESCE($12, start_date),
    end_date = COALESCE($13, end_date),
    is_present = COALESCE($14, is_present),
    location = COALESCE($15, location)
where id = $1
returning *;

-- name: DeleteProject :exec
delete from projects
where id = $1;
