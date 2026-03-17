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

type MockSkillUsecase struct {
	mock.Mock
}

func (m *MockSkillUsecase) CreateSkill(ctx context.Context, s domain.SkillInput) (domain.Skill, error) {
	args := m.Called(ctx, s)
	return args.Get(0).(domain.Skill), args.Error(1)
}

func (m *MockSkillUsecase) GetSkills(ctx context.Context, profileID string) (domain.Skill, error) {
	args := m.Called(ctx, profileID)
	return args.Get(0).(domain.Skill), args.Error(1)
}

func (m *MockSkillUsecase) UpdateSkill(ctx context.Context, id string, s domain.SkillUpdateInput) error {
	args := m.Called(ctx, id, s)
	return args.Error(0)
}

func (m *MockSkillUsecase) DeleteSkill(ctx context.Context, id string) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func TestSkillHandler_CreateSkill(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("success", func(t *testing.T) {
		mockUsecase := new(MockSkillUsecase)
		handlerObj := handler.NewSkillHandler(mockUsecase)
		r := gin.Default()
		r.POST("/skills", handlerObj.CreateSkill)

		input := dto.SkillReq{
			ProfileID:    "1",
			Technologies: []string{"Go", "React"},
		}
		body, _ := json.Marshal(input)
		req, _ := http.NewRequest(http.MethodPost, "/skills", bytes.NewBuffer(body))
		w := httptest.NewRecorder()

		expectedSkill := domain.Skill{ID: "1", ProfileID: "1", Technologies: []string{"Go", "React"}}
		mockUsecase.On("CreateSkill", mock.Anything, mock.Anything).Return(expectedSkill, nil)

		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		mockUsecase.AssertExpectations(t)
	})

	t.Run("bad request - invalid json", func(t *testing.T) {
		mockUsecase := new(MockSkillUsecase)
		handlerObj := handler.NewSkillHandler(mockUsecase)
		r := gin.Default()
		r.POST("/skills", handlerObj.CreateSkill)

		req, _ := http.NewRequest(http.MethodPost, "/skills", bytes.NewBufferString("invalid json"))
		w := httptest.NewRecorder()

		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})

	t.Run("usecase error", func(t *testing.T) {
		mockUsecase := new(MockSkillUsecase)
		handlerObj := handler.NewSkillHandler(mockUsecase)
		r := gin.Default()
		r.POST("/skills", handlerObj.CreateSkill)

		input := dto.SkillReq{
			ProfileID: "1",
		}
		body, _ := json.Marshal(input)
		req, _ := http.NewRequest(http.MethodPost, "/skills", bytes.NewBuffer(body))
		w := httptest.NewRecorder()

		mockUsecase.On("CreateSkill", mock.Anything, mock.Anything).Return(domain.Skill{}, errors.New("internal error"))

		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})
}

func TestSkillHandler_GetSkills(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("success", func(t *testing.T) {
		mockUsecase := new(MockSkillUsecase)
		handlerObj := handler.NewSkillHandler(mockUsecase)
		r := gin.Default()
		r.GET("/profiles/:profile_id/skills", handlerObj.GetSkills)

		req, _ := http.NewRequest(http.MethodGet, "/profiles/1/skills", nil)
		w := httptest.NewRecorder()

		expectedSkill := domain.Skill{ID: "1", ProfileID: "1", Technologies: []string{"Go"}}
		mockUsecase.On("GetSkills", mock.Anything, "1").Return(expectedSkill, nil)

		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		mockUsecase.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockUsecase := new(MockSkillUsecase)
		handlerObj := handler.NewSkillHandler(mockUsecase)
		r := gin.Default()
		r.GET("/profiles/:profile_id/skills", handlerObj.GetSkills)

		req, _ := http.NewRequest(http.MethodGet, "/profiles/1/skills", nil)
		w := httptest.NewRecorder()

		mockUsecase.On("GetSkills", mock.Anything, "1").Return(domain.Skill{}, errors.New("not found"))

		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})
}

