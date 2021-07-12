//go:generate mockgen -destination mock_handlers/product_repo.go . ProductRepo

package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/roman-wb/crud-products/internal/models"
	"github.com/roman-wb/crud-products/pkg/utils"
	"go.uber.org/zap"
)

type ProductRepo interface {
	All(ctx context.Context) (*[]models.Product, error)
	Find(ctx context.Context, id int) (*models.Product, error)
	Create(ctx context.Context, product *models.Product) error
	Update(ctx context.Context, product *models.Product) error
	Destroy(ctx context.Context, id int) error
}

type ProductHandler struct {
	logger      *zap.Logger
	productRepo ProductRepo
}

func NewProductHandler(logger *zap.Logger, productRepo ProductRepo) *ProductHandler {
	return &ProductHandler{
		logger:      logger,
		productRepo: productRepo,
	}
}

func (p ProductHandler) IndexHandler(res http.ResponseWriter, req *http.Request) {
	// Get all products
	products, err := p.productRepo.All(context.Background())
	if err != nil {
		p.logger.Sugar().Error(err)
		utils.ResponseInternalError(res)
		return
	}

	utils.ResponseOK(res, products)
}

func (p ProductHandler) ShowHandler(res http.ResponseWriter, req *http.Request) {
	// Load product
	product, err := p.loadProduct(req)
	if err != nil {
		p.logger.Sugar().Error(err)
		utils.ResponseNotFound(res)
		return
	}

	utils.ResponseOK(res, product)
}

func (p ProductHandler) CreateHandler(res http.ResponseWriter, req *http.Request) {
	// Read JSON params to safe anonymous struct (mass assignment)
	var params struct {
		Name  string  `json:"name"`
		Price float64 `json:"price"`
	}
	decoder := json.NewDecoder(req.Body)
	if err := decoder.Decode(&params); err != nil {
		p.logger.Sugar().Error(err)
		utils.ResponseInternalError(res)
		return
	}

	// Fill and validate model
	product := models.Product{}
	product.Fill(params)
	if messages := product.Validate(); len(messages) > 0 {
		utils.ResponseInvalid(res, messages)
		return
	}

	// Create product in repo
	err := p.productRepo.Create(context.Background(), &product)
	if err != nil {
		p.logger.Sugar().Error(err)
		utils.ResponseInternalError(res)
		return
	}

	utils.ResponseCreate(res, product)
}

func (p ProductHandler) UpdateHandler(res http.ResponseWriter, req *http.Request) {
	// Load product
	product, err := p.loadProduct(req)
	if err != nil {
		p.logger.Sugar().Error(err)
		utils.ResponseNotFound(res)
		return
	}

	// Read JSON params to safe anonymous struct (mass assignment)
	var params struct {
		Name  string  `json:"name"`
		Price float64 `json:"price"`
	}
	decoder := json.NewDecoder(req.Body)
	if err := decoder.Decode(&params); err != nil {
		p.logger.Sugar().Error(err)
		utils.ResponseInternalError(res)
		return
	}

	// Fill and validate model
	product.Fill(params)
	if messages := product.Validate(); len(messages) > 0 {
		utils.ResponseInvalid(res, messages)
		return
	}

	// Update product in repo
	err = p.productRepo.Update(context.Background(), product)
	if err != nil {
		p.logger.Sugar().Error(err)
		utils.ResponseInternalError(res)
		return
	}

	utils.ResponseOK(res, product)
}

func (p ProductHandler) DestroyHandler(res http.ResponseWriter, req *http.Request) {
	// Load product
	product, err := p.loadProduct(req)
	if err != nil {
		p.logger.Sugar().Error(err)
		utils.ResponseNotFound(res)
		return
	}

	// Destroy product in repo
	err = p.productRepo.Destroy(context.Background(), product.Id)
	if err != nil {
		p.logger.Sugar().Error(err)
		utils.ResponseInternalError(res)
		return
	}

	utils.ResponseNoContent(res)
}

func (p ProductHandler) loadProduct(req *http.Request) (*models.Product, error) {
	// Parse query
	query := mux.Vars(req)
	id, err := strconv.Atoi(query["id"])
	if err != nil {
		return nil, err
	}

	// Find product
	product, err := p.productRepo.Find(context.Background(), id)
	if err != nil {
		return nil, err
	}

	return product, nil
}
