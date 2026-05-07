-- name: GetResume :one
SELECT
    p.id,
    p.name,
    p.title,
    p.about,
    p.address,
    p.email,
    p.phone,

    urls.urls,
    experiences.experiences,
    projects.projects,
    educations.educations,
    skills.skills,
    languages.languages

FROM profiles p

-- URLS
LEFT JOIN LATERAL (
    SELECT COALESCE(jsonb_agg(u), '[]'::jsonb) AS urls
    FROM (
        SELECT
            id AS "ID",
            profile_id AS "ProfileID",
            label AS "Label",
            url AS "Url"
        FROM profile_urls
        WHERE profile_id = p.id
    ) u
) urls ON true

-- EXPERIENCES
LEFT JOIN LATERAL (
    SELECT COALESCE(jsonb_agg(e), '[]'::jsonb) AS experiences
    FROM (
        SELECT
            id AS "ID",
            profile_id AS "ProfileID",
            company AS "Company",
            position AS "Position",
            description AS "Description",
            start_date::timestamptz AS "StartDate",
            end_date::timestamptz AS "EndDate",
            is_present AS "IsPresent",
            location AS "Location"
        FROM experiences
        WHERE profile_id = p.id
    ) e
) experiences ON true

-- PROJECTS
LEFT JOIN LATERAL (
    SELECT COALESCE(jsonb_agg(pr), '[]'::jsonb) AS projects
    FROM (
        SELECT
            id AS "ID",
            profile_id AS "ProfileID",
            title AS "Title",
            description AS "Description",
            tech_stacks AS "TechStacks",
            live_demo_url AS "LiveDemoUrl",
            github_repo_url AS "GithubRepoUrl",
            is_live AS "IsLive",
            is_nda AS "IsNda",
            is_featured AS "IsFeatured",
            image_url AS "ImageUrl",
            company AS "Company",
            start_date::timestamptz AS "StartDate",
            end_date::timestamptz AS "EndDate",
            is_present AS "IsPresent",
            location AS "Location"
        FROM projects
        WHERE profile_id = p.id
    ) pr
) projects ON true

-- EDUCATIONS
LEFT JOIN LATERAL (
    SELECT COALESCE(jsonb_agg(ed), '[]'::jsonb) AS educations
    FROM (
        SELECT
            id AS "ID",
            profile_id AS "ProfileID",
            school AS "School",
            degree AS "Degree",
            field_of_study AS "FieldOfStudy",
            gpa AS "GPA",
            start_date::timestamptz AS "StartDate",
            graduation_date::timestamptz AS "GraduationDate",
            is_present AS "IsPresent"
        FROM educations
        WHERE profile_id = p.id
    ) ed
) educations ON true

-- SKILLS
LEFT JOIN LATERAL (
    SELECT COALESCE(jsonb_agg(s), '[]'::jsonb) AS skills
    FROM (
        SELECT
            id AS "ID",
            profile_id AS "ProfileID",
            tools AS "Tools",
            technologies AS "Technologies",
            hard_skills AS "HardSkills",
            soft_skills AS "SoftSkills"
        FROM skills
        WHERE profile_id = p.id
    ) s
) skills ON true

-- LANGUAGES
LEFT JOIN LATERAL (
    SELECT COALESCE(jsonb_agg(l), '[]'::jsonb) AS languages
    FROM (
        SELECT
            id AS "ID",
            profile_id AS "ProfileID",
            language AS "Language",
            proficiency AS "Proficiency"
        FROM languages
        WHERE profile_id = p.id
    ) l
) languages ON true

WHERE p.id = $1;
