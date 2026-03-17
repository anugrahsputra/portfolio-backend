create type proficiency_level as enum (
    'basic',
    'intermediate',
    'advanced',
    'professional',
    'native'
);

create table if not exists profiles (
    id uuid primary key default gen_random_uuid(),
    name text not null,
    about text not null,
    address text not null,
    email text not null unique,
    phone text not null
);

create table if not exists profile_urls (
    id uuid primary key default gen_random_uuid(),
    profile_id uuid not null references profiles(id) on delete cascade,
    label text not null,
    url text not null,
    unique (profile_id, label)
);

create index if not exists idx_profile_urls_profile_id on profile_urls(profile_id);

create table if not exists educations (
    id uuid primary key default gen_random_uuid(),
    profile_id uuid not null references profiles(id) on delete cascade,
    school text not null,
    degree text not null,
    field_of_study text not null,
    gpa numeric(3,2) not null check (gpa >= 0.0 and gpa <= 4.0),
    start_date date not null,
    graduation_date date,
    check (graduation_date is null or graduation_date >= start_date),
    unique (profile_id, school, degree, start_date)
);

create index if not exists idx_educations_profile_id on educations(profile_id);

create table if not exists experiences (
    id uuid primary key default gen_random_uuid(),
    profile_id uuid not null references profiles(id) on delete cascade,
    company text not null,
    position text not null,
    description text[],
    start_date date not null,
    end_date date,
    check (end_date is null or end_date >= start_date),
    unique (profile_id, company, position, start_date)
);

create index if not exists idx_experiences_profile_id on experiences(profile_id);

create table if not exists languages (
    id uuid primary key default gen_random_uuid(),
    profile_id uuid not null references profiles(id) on delete cascade,
    language text not null,
    proficiency proficiency_level not null default 'basic',
    unique (profile_id, language)
);

create index if not exists idx_languages_profile_id on languages(profile_id);

create table if not exists skills (
    id uuid primary key default gen_random_uuid(),
    profile_id uuid not null unique references profiles(id) on delete cascade,
    tools text[],
    technologies text[],
    hard_skills text[],
    soft_skills text[]
);

create index if not exists idx_skills_profile_id on skills(profile_id);

create table if not exists projects(
    id uuid primary key default gen_random_uuid(),
    profile_id uuid not null references profiles(id) on delete cascade,
    title text not null,
    description text not null,
    tech_stacks text[],
    live_demo_url text,
    github_repo_url text,
    is_live boolean default false,
    is_nda boolean default false,
    is_featured boolean default false,
    image_url text,
    company text,
    period text not null,
    location text not null,
    unique (profile_id, title)
);

create index if not exists idx_projects_profile_id on projects(profile_id);
