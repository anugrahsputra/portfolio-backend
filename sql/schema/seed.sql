-- Mock Data for Portfolio Backend

-- 1. Profile
INSERT INTO profiles (name, title, about, address, email, phone)
VALUES (
    'Anugrah Surya Putra', 
    'Senior Software Engineer & Cloud Architect', 
    'Experienced software engineer with a strong focus on building scalable backend systems and cloud infrastructure. Passionate about clean code, performance optimization, and mentoring other developers.', 
    'Jakarta, Indonesia', 
    'anugrah.putra@example.com', 
    '+6281234567890'
) ON CONFLICT (email) DO UPDATE SET
    name = EXCLUDED.name,
    title = EXCLUDED.title,
    about = EXCLUDED.about,
    address = EXCLUDED.address,
    phone = EXCLUDED.phone;

-- 2. Profile URLs
INSERT INTO profile_urls (profile_id, label, url)
VALUES 
    ('550e8400-e29b-41d4-a716-446655440000', 'GitHub', 'https://github.com/anugrahsputra'),
    ('550e8400-e29b-41d4-a716-446655440000', 'LinkedIn', 'https://linkedin.com/in/anugrahsputra'),
    ('550e8400-e29b-41d4-a716-446655440000', 'Portfolio', 'https://downormal.dev')
ON CONFLICT (profile_id, label) DO UPDATE SET url = EXCLUDED.url;

-- 3. Educations
INSERT INTO educations (profile_id, school, degree, field_of_study, gpa, start_date, graduation_date, is_present)
VALUES 
    ('550e8400-e29b-41d4-a716-446655440000', 'University of Indonesia', 'Bachelor of Science', 'Computer Science', 3.85, '2014-09-01', '2018-06-30', false),
    ('550e8400-e29b-41d4-a716-446655440000', 'Tech Institute of Technology', 'Master of Science', 'Artificial Intelligence', 3.95, '2019-09-01', '2021-06-30', false)
ON CONFLICT (profile_id, school, degree, start_date) DO UPDATE SET
    field_of_study = EXCLUDED.field_of_study,
    gpa = EXCLUDED.gpa,
    graduation_date = EXCLUDED.graduation_date,
    is_present = EXCLUDED.is_present;

-- 4. Experiences (Requested: more than 1)
INSERT INTO experiences (profile_id, company, position, description, location, start_date, end_date, is_present)
VALUES 
    (
        '550e8400-e29b-41d4-a716-446655440000', 
        'Global Tech Solutions', 
        'Lead Backend Engineer', 
        ARRAY[
            'Architected and implemented microservices using Go and gRPC',
            'Reduced system latency by 40% through strategic caching and database optimization',
            'Led a team of 5 developers in delivering high-impact features for a fintech platform'
        ], 
        'Remote', 
        '2021-07-01', 
        NULL, 
        true
    ),
    (
        '550e8400-e29b-41d4-a716-446655440000', 
        'Innovate Startups', 
        'Software Engineer', 
        ARRAY[
            'Developed real-time data processing pipelines using Python and Kafka',
            'Integrated third-party APIs for payment processing and identity verification',
            'Automated CI/CD pipelines using GitHub Actions and Kubernetes'
        ], 
        'Jakarta, Indonesia', 
        '2018-07-01', 
        '2021-06-30', 
        false
    )
ON CONFLICT (profile_id, company, position, start_date) DO UPDATE SET
    description = EXCLUDED.description,
    location = EXCLUDED.location,
    end_date = EXCLUDED.end_date,
    is_present = EXCLUDED.is_present;

-- 5. Languages (Requested: more than 1)
INSERT INTO languages (profile_id, language, proficiency)
VALUES 
    ('550e8400-e29b-41d4-a716-446655440000', 'Indonesian', 'native'),
    ('550e8400-e29b-41d4-a716-446655440000', 'English', 'professional'),
    ('550e8400-e29b-41d4-a716-446655440000', 'Japanese', 'basic')
ON CONFLICT (profile_id, language) DO UPDATE SET proficiency = EXCLUDED.proficiency;

-- 6. Skills (Requested: more than 1, interpreted as multiple items in arrays)
INSERT INTO skills (profile_id, tools, technologies, hard_skills, soft_skills)
VALUES 
    (
        '550e8400-e29b-41d4-a716-446655440000', 
        ARRAY['Docker', 'Kubernetes', 'Git', 'Postman', 'VS Code', 'Jira'], 
        ARRAY['Go', 'PostgreSQL', 'Redis', 'Kafka', 'AWS', 'Google Cloud Platform', 'React'], 
        ARRAY['Backend Development', 'System Architecture', 'Cloud Infrastructure', 'API Design', 'Microservices'], 
        ARRAY['Team Leadership', 'Problem Solving', 'Public Speaking', 'Mentoring']
    )
ON CONFLICT (profile_id) DO UPDATE SET
    tools = EXCLUDED.tools,
    technologies = EXCLUDED.technologies,
    hard_skills = EXCLUDED.hard_skills,
    soft_skills = EXCLUDED.soft_skills;

-- 7. Projects (Requested: more than 1)
INSERT INTO projects (profile_id, title, description, tech_stacks, live_demo_url, github_repo_url, is_live, is_nda, is_featured, image_url, company, start_date, end_date, is_present, location)
VALUES 
    (
        '550e8400-e29b-41d4-a716-446655440000', 
        'Scalable Portfolio API', 
        ARRAY[
            'A high-performance RESTful API built with Go and PostgreSQL',
            'Implements Clean Architecture and SOLID principles',
            'Fully containerized and ready for cloud deployment'
        ], 
        ARRAY['Go', 'PostgreSQL', 'sqlc', 'Gin', 'Docker'], 
        'https://api.downormal.dev', 
        'https://github.com/anugrahsputra/portfolio-backend', 
        true, 
        false, 
        true, 
        'https://images.example.com/project1.png', 
        NULL, 
        '2024-01-01', 
        NULL, 
        true, 
        'GitHub'
    ),
    (
        '550e8400-e29b-41d4-a716-446655440000', 
        'Real-time Analytics Dashboard', 
        ARRAY[
            'Interactive dashboard for visualizing real-time data streams',
            'Utilizes WebSockets for low-latency updates',
            'Secure authentication and authorization using OAuth2'
        ], 
        ARRAY['React', 'TypeScript', 'Node.js', 'Redis', 'Socket.io'], 
        'https://analytics.example.com', 
        'https://github.com/anugrahsputra/analytics-dashboard', 
        true, 
        false, 
        true, 
        'https://images.example.com/project2.png', 
        'Innovate Startups', 
        '2020-01-01', 
        '2021-06-30', 
        false, 
        'Jakarta, Indonesia'
    )
ON CONFLICT (profile_id, title) DO UPDATE SET
    description = EXCLUDED.description,
    tech_stacks = EXCLUDED.tech_stacks,
    live_demo_url = EXCLUDED.live_demo_url,
    github_repo_url = EXCLUDED.github_repo_url,
    is_live = EXCLUDED.is_live,
    is_nda = EXCLUDED.is_nda,
    is_featured = EXCLUDED.is_featured,
    image_url = EXCLUDED.image_url,
    company = EXCLUDED.company,
    start_date = EXCLUDED.start_date,
    end_date = EXCLUDED.end_date,
    is_present = EXCLUDED.is_present,
    location = EXCLUDED.location;
