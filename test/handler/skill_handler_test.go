package handler_test

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/anugrahsputra/portfolio-backend/internal/delivery/dto"
	"github.com/anugrahsputra/portfolio-backend/internal/delivery/handler"
	"github.com/anugrahsputra/portfolio-backend/internal/domain"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockSkillRepository struct {
	mock.Mock
}

func (m *MockSkillRepository) CreateSkill(ctx context.Context, s domain.SkillInput) (domain.Skill, error) {
	args := m.Called(ctx, s)
	return args.Get(0).(domain.Skill), args.Error(1)
}

func (m *MockSkillRepository) GetSkills(ctx context.Context, profileID string) (domain.Skill, error) {
	args := m.Called(ctx, profileID)
	return args.Get(0).(domain.Skill), args.Error(1)
}

func (m *MockSkillRepository) UpdateSkill(ctx context.Context, id string, s domain.SkillUpdateInput) error {
	args := m.Called(ctx, id, s)
	return args.Error(0)
}

func (m *MockSkillRepository) DeleteSkill(ctx context.Context, id string) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func TestSkillHandler_CreateSkill(t *testing.T) {
	gin.SetMode(gin.TestMode)
	t.Run("success", func(t *testing.T) {
		mockRepo := new(MockSkillRepository)
		handlerObj := handler.NewSkillHandler(mockRepo)
		r := gin.New()
		r.POST("/skills", handlerObj.CreateSkill)

		input := dto.SkillReq{
			ProfileID:    "1",
			Technologies: []string{"Go", "React"},
		}
		body, _ := json.Marshal(input)
		req, _ := http.NewRequest(http.MethodPost, "/skills", bytes.NewBuffer(body))
		w := httptest.NewRecorder()

		expectedSkill := domain.Skill{ID: "1", ProfileID: "1", Technologies: []string{"Go", "React"}}
		mockRepo.On("CreateSkill", mock.Anything, mock.Anything).Return(expectedSkill, nil)

		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusCreated, w.Code)
		mockRepo.AssertExpectations(t)
	})

	t.Run("bad request - invalid json", func(t *testing.T) {
		mockRepo := new(MockSkillRepository)
		handlerObj := handler.NewSkillHandler(mockRepo)
		r := gin.New()
		r.POST("/skills", handlerObj.CreateSkill)

		req, _ := http.NewRequest(http.MethodPost, "/skills", bytes.NewBufferString("invalid json"))
		w := httptest.NewRecorder()

		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("repository error", func(t *testing.T) {
		mockRepo := new(MockSkillRepository)
		handlerObj := handler.NewSkillHandler(mockRepo)
		r := gin.New()
		r.POST("/skills", handlerObj.CreateSkill)

		input := dto.SkillReq{
			ProfileID: "1",
		}
		body, _ := json.Marshal(input)
		req, _ := http.NewRequest(http.MethodPost, "/skills", bytes.NewBuffer(body))
		w := httptest.NewRecorder()

		mockRepo.On("CreateSkill", mock.Anything, mock.Anything).Return(domain.Skill{}, errors.New("internal error"))

		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})
}

func TestSkillHandler_GetSkills(t *testing.T) {
	gin.SetMode(gin.TestMode)
	t.Run("success", func(t *testing.T) {
		mockRepo := new(MockSkillRepository)
		handlerObj := handler.NewSkillHandler(mockRepo)
		r := gin.New()
		r.GET("/profiles/:profile_id/skills", handlerObj.GetSkills)

		req, _ := http.NewRequest(http.MethodGet, "/profiles/1/skills", nil)
		w := httptest.NewRecorder()

		expectedSkill := domain.Skill{ID: "1", ProfileID: "1", Technologies: []string{"Go"}}
		mockRepo.On("GetSkills", mock.Anything, "1").Return(expectedSkill, nil)

		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		mockRepo.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockRepo := new(MockSkillRepository)
		handlerObj := handler.NewSkillHandler(mockRepo)
		r := gin.New()
		r.GET("/profiles/:profile_id/skills", handlerObj.GetSkills)

		req, _ := http.NewRequest(http.MethodGet, "/profiles/1/skills", nil)
		w := httptest.NewRecorder()

		mockRepo.On("GetSkills", mock.Anything, "1").Return(domain.Skill{}, errors.New("not found"))

		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})
}

func TestSkillHandler_UpdateSkill(t *testing.T) {
	gin.SetMode(gin.TestMode)
	t.Run("success", func(t *testing.T) {
		mockRepo := new(MockSkillRepository)
		handlerObj := handler.NewSkillHandler(mockRepo)
		r := gin.New()
		r.PUT("/skills/:skill_id", handlerObj.UpdateSkill)

		tools := []string{"Updated Skill"}
		input := dto.SkillUpdateReq{Tools: tools}
		body, _ := json.Marshal(input)
		req, _ := http.NewRequest(http.MethodPut, "/skills/1", bytes.NewBuffer(body))
		w := httptest.NewRecorder()

		mockRepo.On("UpdateSkill", mock.Anything, "1", mock.Anything).Return(nil)

		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		mockRepo.AssertExpectations(t)
	})
}

func TestSkillHandler_DeleteSkill(t *testing.T) {
	gin.SetMode(gin.TestMode)
	t.Run("success", func(t *testing.T) {
		mockRepo := new(MockSkillRepository)
		handlerObj := handler.NewSkillHandler(mockRepo)
		r := gin.New()
		r.DELETE("/skills/:skill_id", handlerObj.DeleteSkill)

		req, _ := http.NewRequest(http.MethodDelete, "/skills/1", nil)
		w := httptest.NewRecorder()

		mockRepo.On("DeleteSkill", mock.Anything, "1").Return(nil)

		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		mockRepo.AssertExpectations(t)
	})
}
