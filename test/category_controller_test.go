package test

import (
	"context"
	"database/sql"
	"encoding/json"
	"go-crud/app"
	"go-crud/controller"
	"go-crud/helper"
	"go-crud/middleware"
	"go-crud/model/domain"
	"go-crud/repository"
	"go-crud/service"
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
)

func setupTestDB() *sql.DB {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/go_crud_test")
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}

func setupRouter(db *sql.DB) http.Handler {
	categoryRepository := repository.NewCategoryRepostory(db)
	categoryService := service.NewCategoryService(categoryRepository)
	categoryController := controller.NewCategoryController(categoryService)
	router := app.NewRouter(categoryController)

	return middleware.NewAuthMiddleware(router)
}

func truncateCategory(db *sql.DB) {
	db.Exec("TRUNCATE category")
}

func TestCreateCategorySuccess(t *testing.T) {
	db := setupTestDB()
	truncateCategory(db)

	router := setupRouter(db)
	requestBody := strings.NewReader(`
		{
			"name" : "Fashion"
		}
	`)

	request := httptest.NewRequest(http.MethodPost, "/api/categories", requestBody)
	request.Header.Add("content-type", "application/json")
	request.Header.Add("X-API-Key", "RAHASIA")

	recoder := httptest.NewRecorder()

	router.ServeHTTP(recoder, request)

	r := recoder.Result()
	assert.Equal(t, 200, r.StatusCode)
}

func TestCreateCategoryFailed(t *testing.T) {
	db := setupTestDB()
	truncateCategory(db)

	router := setupRouter(db)
	requestBody := strings.NewReader(`
		{
			"name" : ""
		}
	`)

	request := httptest.NewRequest(http.MethodPost, "/api/categories", requestBody)
	request.Header.Add("content-type", "application/json")
	request.Header.Add("X-API-Key", "RAHASIA")

	recoder := httptest.NewRecorder()

	router.ServeHTTP(recoder, request)

	r := recoder.Result()
	assert.Equal(t, 400, r.StatusCode)
}

func TestUpdateCategorySuccess(t *testing.T) {
	db := setupTestDB()
	truncateCategory(db)

	categoryRepository := repository.NewCategoryRepostory(db)
	category := categoryRepository.Save(context.Background(), domain.Category{
		Name: "Gadget",
	})

	router := setupRouter(db)
	requestBody := strings.NewReader(`
		{
			"name" : "Gadget update"
		}
	`)

	request := httptest.NewRequest(http.MethodPut, "/api/categories/"+strconv.Itoa(category.Id), requestBody)
	request.Header.Add("content-type", "application/json")
	request.Header.Add("X-API-Key", "RAHASIA")

	recoder := httptest.NewRecorder()

	router.ServeHTTP(recoder, request)

	r := recoder.Result()
	assert.Equal(t, 200, r.StatusCode)
}

func TestUpdateCategoryFailed(t *testing.T) {
	db := setupTestDB()
	truncateCategory(db)

	categoryRepository := repository.NewCategoryRepostory(db)
	category := categoryRepository.Save(context.Background(), domain.Category{
		Name: "Gadget",
	})

	router := setupRouter(db)
	requestBody := strings.NewReader(`
		{
			"name" : ""
		}
	`)

	request := httptest.NewRequest(http.MethodPut, "/api/categories/"+strconv.Itoa(category.Id), requestBody)
	request.Header.Add("content-type", "application/json")
	request.Header.Add("X-API-Key", "RAHASIA")

	recoder := httptest.NewRecorder()

	router.ServeHTTP(recoder, request)

	r := recoder.Result()
	assert.Equal(t, 400, r.StatusCode)
}

func TestGetCategorySuccess(t *testing.T) {
	db := setupTestDB()
	truncateCategory(db)

	categoryRepository := repository.NewCategoryRepostory(db)
	category := categoryRepository.Save(context.Background(), domain.Category{
		Name: "Gadget",
	})

	router := setupRouter(db)

	request := httptest.NewRequest(http.MethodGet, "/api/categories/"+strconv.Itoa(category.Id), nil)
	request.Header.Add("content-type", "application/json")
	request.Header.Add("X-API-Key", "RAHASIA")

	recoder := httptest.NewRecorder()

	router.ServeHTTP(recoder, request)

	r := recoder.Result()
	assert.Equal(t, 200, r.StatusCode)
}

func TestGetCategoryFailed(t *testing.T) {
	db := setupTestDB()
	truncateCategory(db)

	categoryRepository := repository.NewCategoryRepostory(db)
	categoryRepository.Save(context.Background(), domain.Category{
		Name: "Gadget",
	})

	router := setupRouter(db)

	request := httptest.NewRequest(http.MethodGet, "/api/categories/10000", nil)
	request.Header.Add("content-type", "application/json")
	request.Header.Add("X-API-Key", "RAHASIA")

	recoder := httptest.NewRecorder()

	router.ServeHTTP(recoder, request)

	r := recoder.Result()
	assert.Equal(t, 404, r.StatusCode)
}

func TestDeleteCategorySuccess(t *testing.T) {
	db := setupTestDB()
	truncateCategory(db)

	categoryRepository := repository.NewCategoryRepostory(db)
	category := categoryRepository.Save(context.Background(), domain.Category{
		Name: "Gadget",
	})

	router := setupRouter(db)

	request := httptest.NewRequest(http.MethodDelete, "/api/categories/"+strconv.Itoa(category.Id), nil)
	request.Header.Add("content-type", "application/json")
	request.Header.Add("X-API-Key", "RAHASIA")

	recoder := httptest.NewRecorder()

	router.ServeHTTP(recoder, request)

	r := recoder.Result()
	assert.Equal(t, 200, r.StatusCode)
}

func TestDeleteCategoryFailed(t *testing.T) {
	db := setupTestDB()
	truncateCategory(db)

	categoryRepository := repository.NewCategoryRepostory(db)
	categoryRepository.Save(context.Background(), domain.Category{
		Name: "Gadget",
	})

	router := setupRouter(db)

	request := httptest.NewRequest(http.MethodDelete, "/api/categories/100", nil)
	request.Header.Add("content-type", "application/json")
	request.Header.Add("X-API-Key", "RAHASIA")

	recoder := httptest.NewRecorder()

	router.ServeHTTP(recoder, request)

	r := recoder.Result()
	assert.Equal(t, 200, r.StatusCode)
}

func TestGetAllCategorySuccess(t *testing.T) {
	db := setupTestDB()
	truncateCategory(db)

	categoryRepository := repository.NewCategoryRepostory(db)
	category0 := categoryRepository.Save(context.Background(), domain.Category{
		Name: "Gadget",
	})
	category1 := categoryRepository.Save(context.Background(), domain.Category{
		Name: "Gadget",
	})

	router := setupRouter(db)

	request := httptest.NewRequest(http.MethodGet, "/api/categories", nil)
	request.Header.Add("content-type", "application/json")
	request.Header.Add("X-API-Key", "RAHASIA")

	recoder := httptest.NewRecorder()

	router.ServeHTTP(recoder, request)

	r := recoder.Result()
	assert.Equal(t, 200, r.StatusCode)

	b, _ := io.ReadAll(r.Body)
	var resBody map[string]interface{}
	json.Unmarshal(b, &resBody)

	var categories = resBody["data"].([]map[string]interface{})

	assert.Equal(t, category0.Id, categories[0]["id"])
	assert.Equal(t, category1.Id, categories[1]["id"])

}

func TestGetAllCategoryFailed(t *testing.T) {

}
