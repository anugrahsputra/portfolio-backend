# API Documentation

This document provides a summary of all API endpoints available in the Portfolio Backend. All endpoints are prefixed with `/api/v1`.

## Profiles

| HTTP Method | Path | Description | Handler Method |
|-------------|------|-------------|----------------|
| POST | `/api/v1/profile` | Create a new profile | `CreateProfile` |
| GET | `/api/v1/profile/:id` | Get a profile by ID | `GetProfile` |
| PUT | `/api/v1/profile/:id` | Update a profile by ID | `UpdateProfile` |
| DELETE | `/api/v1/profile/:id` | Delete a profile by ID | `DeleteProfile` |

## Educations

| HTTP Method | Path | Description | Handler Method |
|-------------|------|-------------|----------------|
| POST | `/api/v1/education` | Create a new education entry | `CreateEducation` |
| GET | `/api/v1/education/:profile_id` | Get education entries by profile ID | `GetEducation` |
| PUT | `/api/v1/education/:education_id` | Update an education entry by ID | `UpdateEducation` |
| DELETE | `/api/v1/education/:education_id` | Delete an education entry by ID | `DeleteEducation` |

## Experiences

| HTTP Method | Path | Description | Handler Method |
|-------------|------|-------------|----------------|
| POST | `/api/v1/experience` | Create a new experience entry | `CreateExperience` |
| GET | `/api/v1/experience/:profile_id` | Get experience entries by profile ID | `GetExperiences` |
| PUT | `/api/v1/experience/:experience_id` | Update an experience entry by ID | `UpdateExperience` |
| DELETE | `/api/v1/experience/:experience_id` | Delete an experience entry by ID | `DeleteExperience` |

## Languages

| HTTP Method | Path | Description | Handler Method |
|-------------|------|-------------|----------------|
| POST | `/api/v1/language` | Create a new language entry | `CreateLanguage` |
| GET | `/api/v1/language/:profile_id` | Get language entries by profile ID | `GetLanguages` |
| PUT | `/api/v1/language/:language_id` | Update a language entry by ID | `UpdateLanguage` |
| DELETE | `/api/v1/language/:language_id` | Delete a language entry by ID | `DeleteLanguage` |

## Profile URLs

| HTTP Method | Path | Description | Handler Method |
|-------------|------|-------------|----------------|
| POST | `/api/v1/profile-url` | Create a new profile URL | `CreateProfileUrl` |
| GET | `/api/v1/profile-url/:profile_url_id` | Get a profile URL entry by ID | `GetProfileUrl` |
| PUT | `/api/v1/profile-url/:profile_url_id` | Update a profile URL entry by ID | `UpdateProfileUrl` |
| DELETE | `/api/v1/profile-url/:profile_url_id` | Delete a profile URL entry by ID | `DeleteProfileUrl` |

## Skills

| HTTP Method | Path | Description | Handler Method |
|-------------|------|-------------|----------------|
| POST | `/api/v1/skill` | Create a new skill entry | `CreateSkill` |
| GET | `/api/v1/skill/:profile_id` | Get skill entries by profile ID | `GetSkills` |
| PUT | `/api/v1/skill/:skill_id` | Update a skill entry by ID | `UpdateSkill` |
| DELETE | `/api/v1/skill/:skill_id` | Delete a skill entry by ID | `DeleteSkill` |
