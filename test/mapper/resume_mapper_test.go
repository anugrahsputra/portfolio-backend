package mapper_test

import (
	"testing"
	"time"

	"github.com/anugrahsputra/portfolio-backend/internal/db"
	"github.com/anugrahsputra/portfolio-backend/internal/mapper"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestToResumeDomain(t *testing.T) {
	id := uuid.New()

	t.Run("with other fields", func(t *testing.T) {
		// Mock data
		urls := []map[string]any{
			{
				"ID":    "url-id",
				"Label": "GitHub",
				"Url":   "https://github.com/user",
			},
		}

		skills := []map[string]any{
			{
				"ID":           "skill-id",
				"ProfileID":    id,
				"Tools":        []string{"tool1", "tool2"},
				"Technologies": []string{"tech1", "tech2"},
				"HardSkills":   []string{"hard1", "hard2"},
				"SoftSkills":   []string{"soft1", "soft2"},
			},
		}
		languages := []map[string]any{
			{
				"ID":          "language-id",
				"ProfileID":   id,
				"Language":    "english",
				"Proficiency": "beginner",
			},
		}

		experiences := []map[string]any{
			{
				"ID":        "experience-id",
				"ProfileID": id,
				"Company":   "the company",
				"Position":  "the positiion",
				"Description": []string{
					"desc1",
					"desc2",
				},
				"Location":  "the location",
				"StartDate": "2006-01-02T15:04:05Z",
				"EndDate":   "2027-05-08T00:00:00Z",
				"IsPresent": true,
			},
		}

		educations := []map[string]any{
			{
				"ID":             "education-id",
				"ProfileID":      id,
				"School":         "the school",
				"Degree":         "the degree",
				"FieldOfStudy":   "the field of study",
				"GPA":            3.8,
				"StartDate":      "2006-01-02T15:04:05Z",
				"GraduationDate": "2027-05-08T00:00:00Z",
				"IsPresent":      true,
			},
		}

		projects := []map[string]any{
			{
				"ID":            "project-id",
				"ProfileID":     id,
				"Title":         "the project name",
				"Description":   []string{"the project description"},
				"TechStacks":    []string{"tech1", "tech2"},
				"LiveDemoUrl":   "https://live.com",
				"GithubRepoUrl": "https://github.com/user/project",
				"IsLive":        true,
				"IsNda":         false,
				"IsFeatured":    true,
				"ImageUrl":      "https://image.com",
				"Company":       "the company",
				"StartDate":     "2006-01-02T15:04:05Z",
				"EndDate":       "2027-05-08T00:00:00Z",
				"IsPresent":     true,
				"Location":      "remote",
			},
		}

		dbResume := db.GetResumeRow{
			ID:          id,
			Name:        "John Doe",
			Title:       "Developer",
			About:       "I am a software developer",
			Address:     "123 Street",
			Email:       "john@example.com",
			Phone:       "123456789",
			Urls:        urls,
			Skills:      skills,
			Languages:   languages,
			Experiences: experiences,
			Educations:  educations,
			Projects:    projects,
		}

		domainResume, err := mapper.ToResumeDomain(dbResume)
		if err != nil {
			t.Fatalf("ToResumeDomain failed: %v", err)
		}
		assert.Equal(t, id.String(), domainResume.ID)
		assert.Equal(t, "John Doe", domainResume.Name)
		assert.Equal(t, "Developer", domainResume.Title)
		assert.Equal(t, "I am a software developer", domainResume.About)
		assert.Equal(t, "123 Street", domainResume.Address)
		assert.Equal(t, "john@example.com", domainResume.Email)
		assert.Equal(t, "123456789", domainResume.Phone)

		assert.Len(t, domainResume.Url, 1)
		assert.Equal(t, "GitHub", domainResume.Url[0].Label)
		assert.Equal(t, "https://github.com/user", domainResume.Url[0].Url)

		assert.Len(t, domainResume.Skills, 1)
		assert.Equal(t, "skill-id", domainResume.Skills[0].ID)
		assert.Equal(t, id.String(), domainResume.Skills[0].ProfileID)
		assert.Equal(t, []string{"tool1", "tool2"}, domainResume.Skills[0].Tools)

		assert.Len(t, domainResume.Languages, 1)
		assert.Equal(t, "language-id", domainResume.Languages[0].ID)
		assert.Equal(t, "english", domainResume.Languages[0].Language)

		assert.Len(t, domainResume.Experiences, 1)
		assert.Equal(t, "experience-id", domainResume.Experiences[0].ID)
		assert.Equal(t, "the company", domainResume.Experiences[0].Company)
		assert.Equal(t, true, domainResume.Experiences[0].IsPresent)
		assert.Equal(t, "2006-01-02T15:04:05Z", domainResume.Experiences[0].StartDate.Format(time.RFC3339))

		assert.Len(t, domainResume.Educations, 1)
		assert.Equal(t, "education-id", domainResume.Educations[0].ID)
		assert.Equal(t, "the school", domainResume.Educations[0].School)
		assert.Equal(t, 3.8, domainResume.Educations[0].Gpa)

		assert.Len(t, domainResume.Projects, 1)
		assert.Equal(t, "project-id", domainResume.Projects[0].ID)
		assert.Equal(t, "the project name", domainResume.Projects[0].Title)
	})
}
