package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/felipefabricio/golang-rest-api/internal/dto"
	"github.com/felipefabricio/golang-rest-api/internal/entity"
	"github.com/felipefabricio/golang-rest-api/internal/infra/database"
	entityPkg "github.com/felipefabricio/golang-rest-api/pkg/entity"
	"github.com/go-chi/chi"
)

type ProductHandler struct {
	ProductDB database.ProductInterface
}

func NewProductHandler(db database.ProductInterface) *ProductHandler {
	return &ProductHandler{
		ProductDB: db,
	}
}

func (h *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product dto.CreateProductInput
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	p, err := entity.NewProduct(product.Name, product.Description, product.Price)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = h.ProductDB.Create(p)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (h *ProductHandler) GetAllProducts(w http.ResponseWriter, r *http.Request) {
	//Pegar os parametros do request
	page, err := strconv.Atoi(chi.URLParam(r, "page"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	// sort := chi.URLParam(r, "sort")

	//Chamar o FindAll no Repository
	products, err := h.ProductDB.FindAllProducts(page, 10, "name")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	//Encodar a lista de produtos retornada
	productsDecoded, err := json.Marshal(&products)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	//Retornar o resultado
	w.Write(productsDecoded)
}

func (h *ProductHandler) GetProductById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	product, err := h.ProductDB.FindById(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(product)
}

func (h *ProductHandler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		log.Fatal("Não foi informado o Id")
		return
	}

	var product entity.Product
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Fatal(err)
		return
	}

	product.ID, err = entityPkg.ParseID(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Fatal(err)
		return
	}

	_, err = h.ProductDB.FindById(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	err = h.ProductDB.Update(&product)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Fatal(err)
		return
	}
	w.WriteHeader(http.StatusOK)
}
