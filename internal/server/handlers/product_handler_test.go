package handlers

import (
	"bytes"
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/roman-wb/crud-products/internal/models"
	"github.com/roman-wb/crud-products/internal/repos"
	"github.com/roman-wb/crud-products/internal/server/handlers/mock_handlers"
	"github.com/roman-wb/crud-products/pkg/utils"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"go.uber.org/zap/zaptest"
)

func Test_NewProductHandler(t *testing.T) {
	logger := &zap.Logger{}
	repo := repos.NewProductRepo(&pgxpool.Pool{})

	handler := NewProductHandler(logger, repo)

	assert.Equal(t, logger, handler.logger)
	assert.Equal(t, repo, handler.productRepo)
}

func Test_Product_IndexHandler_Case1_ExecError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mock_handlers.NewMockProductRepo(ctrl)
	handler := NewProductHandler(zaptest.NewLogger(t), mock)

	mock.
		EXPECT().
		All(context.Background()).
		Return(nil, errors.New("Some error..."))

	res := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)

	handler.IndexHandler(res, req)

	assert.Equal(t, utils.ContentTypeJSON, res.Header().Values(utils.HeaderContentType)[0])
	assert.Equal(t, http.StatusInternalServerError, res.Result().StatusCode)
	assert.Equal(t, utils.DataToJson(utils.ResponseMessage{
		Message: utils.MessageInternalError,
	}), utils.BodyToString(res.Body))
}

func Test_Product_IndexHandler_Case2_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mock_handlers.NewMockProductRepo(ctrl)
	handler := NewProductHandler(zaptest.NewLogger(t), mock)

	mock.
		EXPECT().
		All(context.Background()).
		Return(&[]models.Product{
			{
				Id:    1,
				Name:  "Name 1",
				Price: 100.00,
			},
			{
				Id:    2,
				Name:  "Name 2",
				Price: 200.99,
			},
		}, nil)

	res := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)

	handler.IndexHandler(res, req)

	assert.Equal(t, utils.ContentTypeJSON, res.Header().Values(utils.HeaderContentType)[0])
	assert.Equal(t, http.StatusOK, res.Result().StatusCode)
	assert.Equal(t, utils.DataToJson([]models.Product{
		{Id: 1, Name: "Name 1", Price: 100.00},
		{Id: 2, Name: "Name 2", Price: 200.99},
	}), utils.BodyToString(res.Body))
}

func Test_Product_IndexHandler_Case3_Blank(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mock_handlers.NewMockProductRepo(ctrl)
	handler := NewProductHandler(zaptest.NewLogger(t), mock)

	mock.
		EXPECT().
		All(context.Background()).
		Return(&[]models.Product{}, nil)

	res := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)

	handler.IndexHandler(res, req)

	assert.Equal(t, utils.ContentTypeJSON, res.Header().Values(utils.HeaderContentType)[0])
	assert.Equal(t, http.StatusOK, res.Result().StatusCode)
	assert.Equal(t, utils.DataToJson([]models.Product{}), utils.BodyToString(res.Body))
}

func Test_Product_ShowHandler_Case1_ParseQueryError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mock_handlers.NewMockProductRepo(ctrl)
	handler := NewProductHandler(zaptest.NewLogger(t), mock)

	res := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	req = mux.SetURLVars(req, map[string]string{"id": "abc"})

	handler.ShowHandler(res, req)

	assert.Equal(t, utils.ContentTypeJSON, res.Header().Values(utils.HeaderContentType)[0])
	assert.Equal(t, http.StatusNotFound, res.Result().StatusCode)
	assert.Equal(t, utils.DataToJson(utils.ResponseMessage{
		Message: utils.MessageNotFound,
	}), utils.BodyToString(res.Body))
}

func Test_Product_ShowHandler_Case2_FindError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mock_handlers.NewMockProductRepo(ctrl)
	handler := NewProductHandler(zaptest.NewLogger(t), mock)

	mock.
		EXPECT().
		Find(context.Background(), 1).
		Return(nil, errors.New("Some error..."))

	res := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	req = mux.SetURLVars(req, map[string]string{"id": "1"})

	handler.ShowHandler(res, req)

	assert.Equal(t, utils.ContentTypeJSON, res.Header().Values(utils.HeaderContentType)[0])
	assert.Equal(t, http.StatusNotFound, res.Result().StatusCode)
	assert.Equal(t, utils.DataToJson(utils.ResponseMessage{
		Message: utils.MessageNotFound,
	}), utils.BodyToString(res.Body))
}

func Test_Product_ShowHandler_Case3_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mock_handlers.NewMockProductRepo(ctrl)
	handler := NewProductHandler(zaptest.NewLogger(t), mock)

	mock.
		EXPECT().
		Find(context.Background(), 1).
		Return(&models.Product{
			Id:    1,
			Name:  "Name 1",
			Price: 100.00,
		}, nil)

	res := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	req = mux.SetURLVars(req, map[string]string{"id": "1"})

	handler.ShowHandler(res, req)

	assert.Equal(t, utils.ContentTypeJSON, res.Header().Values(utils.HeaderContentType)[0])
	assert.Equal(t, http.StatusOK, res.Result().StatusCode)
	assert.Equal(t, utils.DataToJson(models.Product{
		Id:    1,
		Name:  "Name 1",
		Price: 100.00,
	}), utils.BodyToString(res.Body))
}

func Test_Product_CreateHandler_Case1_ParseJsonError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mock_handlers.NewMockProductRepo(ctrl)
	handler := NewProductHandler(zaptest.NewLogger(t), mock)

	res := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", bytes.NewBufferString(``))

	handler.CreateHandler(res, req)

	assert.Equal(t, utils.ContentTypeJSON, res.Header().Values(utils.HeaderContentType)[0])
	assert.Equal(t, http.StatusInternalServerError, res.Result().StatusCode)
	assert.Equal(t, utils.DataToJson(utils.ResponseMessage{
		Message: utils.MessageInternalError,
	}), utils.BodyToString(res.Body))
}

func Test_Product_CreateHandler_Case2_InvalidParams(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mock_handlers.NewMockProductRepo(ctrl)
	handler := NewProductHandler(zaptest.NewLogger(t), mock)

	res := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", bytes.NewBufferString(`
		{
			"name": "",
			"price": -1
		}
	`))

	handler.CreateHandler(res, req)

	assert.Equal(t, utils.ContentTypeJSON, res.Header().Values(utils.HeaderContentType)[0])
	assert.Equal(t, http.StatusUnprocessableEntity, res.Result().StatusCode)
	assert.Equal(t, utils.DataToJson([]string{
		models.ProductValidationNameRequired,
		models.ProductValidationPriceGte,
	}), utils.BodyToString(res.Body))
}

func Test_Product_CreateHandler_Case3_ExecError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mock_handlers.NewMockProductRepo(ctrl)
	handler := NewProductHandler(zaptest.NewLogger(t), mock)

	mock.
		EXPECT().
		Create(context.Background(), &models.Product{
			Name:  "Name 1",
			Price: 100.00,
		}).
		Return(errors.New("Some error..."))

	res := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", bytes.NewBufferString(`
		{
			"name": "Name 1",
			"price": 100.00
		}
	`))

	handler.CreateHandler(res, req)

	assert.Equal(t, utils.ContentTypeJSON, res.Header().Values(utils.HeaderContentType)[0])
	assert.Equal(t, http.StatusInternalServerError, res.Result().StatusCode)
	assert.Equal(t, utils.DataToJson(utils.ResponseMessage{
		Message: utils.MessageInternalError,
	}), utils.BodyToString(res.Body))
}

func Test_Product_CreateHandler_Case4_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mock_handlers.NewMockProductRepo(ctrl)
	handler := NewProductHandler(zaptest.NewLogger(t), mock)

	mock.
		EXPECT().
		Create(context.Background(), &models.Product{
			Name:  "Name 1",
			Price: 100.00,
		}).
		Return(nil)

	res := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", bytes.NewBufferString(`
		{
			"name": "Name 1",
			"price": 100.00
		}
	`))

	handler.CreateHandler(res, req)

	assert.Equal(t, utils.ContentTypeJSON, res.Header().Values(utils.HeaderContentType)[0])
	assert.Equal(t, http.StatusCreated, res.Result().StatusCode)
	assert.Equal(t, utils.DataToJson(&models.Product{
		Name:  "Name 1",
		Price: 100.00,
	}), utils.BodyToString(res.Body))
}

func Test_Product_UpdateHandler_Case1_ParseQueryError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mock_handlers.NewMockProductRepo(ctrl)
	handler := NewProductHandler(zaptest.NewLogger(t), mock)

	res := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	req = mux.SetURLVars(req, map[string]string{"id": "abc"})

	handler.UpdateHandler(res, req)

	assert.Equal(t, utils.ContentTypeJSON, res.Header().Values(utils.HeaderContentType)[0])
	assert.Equal(t, http.StatusNotFound, res.Result().StatusCode)
	assert.Equal(t, utils.DataToJson(utils.ResponseMessage{
		Message: utils.MessageNotFound,
	}), utils.BodyToString(res.Body))
}

func Test_Product_UpdateHandler_Case2_FindError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mock_handlers.NewMockProductRepo(ctrl)
	handler := NewProductHandler(zaptest.NewLogger(t), mock)

	mock.
		EXPECT().
		Find(context.Background(), 1).
		Return(nil, errors.New("Some error..."))

	res := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	req = mux.SetURLVars(req, map[string]string{"id": "1"})

	handler.UpdateHandler(res, req)

	assert.Equal(t, utils.ContentTypeJSON, res.Header().Values(utils.HeaderContentType)[0])
	assert.Equal(t, http.StatusNotFound, res.Result().StatusCode)
	assert.Equal(t, utils.DataToJson(utils.ResponseMessage{
		Message: utils.MessageNotFound,
	}), utils.BodyToString(res.Body))
}

func Test_Product_UpdateHandler_Case3_ParseJsonError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mock_handlers.NewMockProductRepo(ctrl)
	handler := NewProductHandler(zaptest.NewLogger(t), mock)

	mock.
		EXPECT().
		Find(context.Background(), 1).
		Return(&models.Product{
			Id:    1,
			Name:  "Name 1",
			Price: 100.00,
		}, nil)

	res := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", bytes.NewBufferString(``))
	req = mux.SetURLVars(req, map[string]string{"id": "1"})

	handler.UpdateHandler(res, req)

	assert.Equal(t, utils.ContentTypeJSON, res.Header().Values(utils.HeaderContentType)[0])
	assert.Equal(t, http.StatusInternalServerError, res.Result().StatusCode)
	assert.Equal(t, utils.DataToJson(utils.ResponseMessage{
		Message: utils.MessageInternalError,
	}), utils.BodyToString(res.Body))
}

func Test_Product_UpdateHandler_Case4_InvalidParams(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mock_handlers.NewMockProductRepo(ctrl)
	handler := NewProductHandler(zaptest.NewLogger(t), mock)

	mock.
		EXPECT().
		Find(context.Background(), 1).
		Return(&models.Product{
			Id:    1,
			Name:  "Name 1",
			Price: 100.00,
		}, nil)

	res := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", bytes.NewBufferString(`
		{
			"name": "",
			"price": -1
		}
	`))
	req = mux.SetURLVars(req, map[string]string{"id": "1"})

	handler.UpdateHandler(res, req)

	assert.Equal(t, utils.ContentTypeJSON, res.Header().Values(utils.HeaderContentType)[0])
	assert.Equal(t, http.StatusUnprocessableEntity, res.Result().StatusCode)
	assert.Equal(t, utils.DataToJson([]string{
		models.ProductValidationNameRequired,
		models.ProductValidationPriceGte,
	}), utils.BodyToString(res.Body))
}

func Test_Product_UpdateHandler_Case5_ExecError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mock_handlers.NewMockProductRepo(ctrl)
	handler := NewProductHandler(zaptest.NewLogger(t), mock)

	mock.
		EXPECT().
		Find(context.Background(), 1).
		Return(&models.Product{
			Id:    1,
			Name:  "Name 1",
			Price: 100.00,
		}, nil)

	mock.
		EXPECT().
		Update(context.Background(), &models.Product{
			Id:    1,
			Name:  "Name 1 - update",
			Price: 999.00,
		}).
		Return(errors.New("Some error..."))

	res := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", bytes.NewBufferString(`
		{
			"name": "Name 1 - update",
			"price": 999.00
		}
	`))
	req = mux.SetURLVars(req, map[string]string{"id": "1"})

	handler.UpdateHandler(res, req)

	assert.Equal(t, utils.ContentTypeJSON, res.Header().Values(utils.HeaderContentType)[0])
	assert.Equal(t, http.StatusInternalServerError, res.Result().StatusCode)
	assert.Equal(t, utils.DataToJson(utils.ResponseMessage{
		Message: utils.MessageInternalError,
	}), utils.BodyToString(res.Body))
}

func Test_Product_UpdateHandler_Case6_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mock_handlers.NewMockProductRepo(ctrl)
	handler := NewProductHandler(zaptest.NewLogger(t), mock)

	mock.
		EXPECT().
		Find(context.Background(), 1).
		Return(&models.Product{
			Id:    1,
			Name:  "Name 1",
			Price: 100.00,
		}, nil)

	mock.
		EXPECT().
		Update(context.Background(), &models.Product{
			Id:    1,
			Name:  "Name 1 - update",
			Price: 999.00,
		}).
		Return(nil)

	res := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", bytes.NewBufferString(`
		{
			"name": "Name 1 - update",
			"price": 999.00
		}
	`))
	req = mux.SetURLVars(req, map[string]string{"id": "1"})

	handler.UpdateHandler(res, req)

	assert.Equal(t, utils.ContentTypeJSON, res.Header().Values(utils.HeaderContentType)[0])
	assert.Equal(t, http.StatusOK, res.Result().StatusCode)
	assert.Equal(t, utils.DataToJson(&models.Product{
		Id:    1,
		Name:  "Name 1 - update",
		Price: 999.00,
	}), utils.BodyToString(res.Body))
}

func Test_Product_DestroyHandler_Case1_ParseQueryError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mock_handlers.NewMockProductRepo(ctrl)
	handler := NewProductHandler(zaptest.NewLogger(t), mock)

	res := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	req = mux.SetURLVars(req, map[string]string{"id": "abc"})

	handler.DestroyHandler(res, req)

	assert.Equal(t, utils.ContentTypeJSON, res.Header().Values(utils.HeaderContentType)[0])
	assert.Equal(t, http.StatusNotFound, res.Result().StatusCode)
	assert.Equal(t, utils.DataToJson(utils.ResponseMessage{
		Message: utils.MessageNotFound,
	}), utils.BodyToString(res.Body))
}

func Test_Product_DestroyHandler_Case2_FindError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mock_handlers.NewMockProductRepo(ctrl)
	handler := NewProductHandler(zaptest.NewLogger(t), mock)

	mock.
		EXPECT().
		Find(context.Background(), 1).
		Return(nil, errors.New("Some error..."))

	res := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	req = mux.SetURLVars(req, map[string]string{"id": "1"})

	handler.DestroyHandler(res, req)

	assert.Equal(t, utils.ContentTypeJSON, res.Header().Values(utils.HeaderContentType)[0])
	assert.Equal(t, http.StatusNotFound, res.Result().StatusCode)
	assert.Equal(t, utils.DataToJson(utils.ResponseMessage{
		Message: utils.MessageNotFound,
	}), utils.BodyToString(res.Body))
}

func Test_Product_DestroyHandler_Case3_ExecError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mock_handlers.NewMockProductRepo(ctrl)
	handler := NewProductHandler(zaptest.NewLogger(t), mock)

	mock.
		EXPECT().
		Find(context.Background(), 1).
		Return(&models.Product{
			Id:    1,
			Name:  "Name 1",
			Price: 100.00,
		}, nil)

	mock.
		EXPECT().
		Destroy(context.Background(), 1).
		Return(errors.New("Some error..."))

	res := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	req = mux.SetURLVars(req, map[string]string{"id": "1"})

	handler.DestroyHandler(res, req)

	assert.Equal(t, utils.ContentTypeJSON, res.Header().Values(utils.HeaderContentType)[0])
	assert.Equal(t, http.StatusInternalServerError, res.Result().StatusCode)
	assert.Equal(t, utils.DataToJson(utils.ResponseMessage{
		Message: utils.MessageInternalError,
	}), utils.BodyToString(res.Body))
}

func Test_Product_DestroyHandler_Case4_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mock_handlers.NewMockProductRepo(ctrl)
	handler := NewProductHandler(zaptest.NewLogger(t), mock)

	mock.
		EXPECT().
		Find(context.Background(), 1).
		Return(&models.Product{
			Id:    1,
			Name:  "Name 1",
			Price: 100.00,
		}, nil)

	mock.
		EXPECT().
		Destroy(context.Background(), 1).
		Return(nil)

	res := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	req = mux.SetURLVars(req, map[string]string{"id": "1"})

	handler.DestroyHandler(res, req)

	assert.Equal(t, utils.ContentTypeJSON, res.Header().Values(utils.HeaderContentType)[0])
	assert.Equal(t, http.StatusNoContent, res.Result().StatusCode)
	assert.Equal(t, "", utils.BodyToString(res.Body))
}
