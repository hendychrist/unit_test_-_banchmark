package service

import (
	"belajar-golang-unit-test/entity"
	"belajar-golang-unit-test/repository"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

var categoryRepository = &repository.CategoryRepositoryMock{Mock: mock.Mock{}}
var categoryService = CategoryService{Repository: categoryRepository}

func TestCategoryService_Get(t *testing.T) {
	// Program mock
	categoryRepository.Mock.On("FindById", "1").Return(nil)

	category, err := categoryService.Get("1")
	assert.Nil(t, category)
	assert.NotNil(t, err)
}

func TestCategoryService_GetSuccess(t *testing.T) {
	category := entity.Category{
		Id:   "1",
		Name: "Laptop",
	}

	// Jadi memanggil method FindById dengan parameter 2 > expected returnya ya category diatas
	categoryRepository.Mock.On("FindById", "2").Return(category)

	// Lalu di test di bawah ini apakah sesuai dengan expected return dari kita
	result, err := categoryService.Get("2")
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, category.Id, result.Id)
	assert.Equal(t, category.Name, result.Name)

	/*
		Jadi dengan Mock kita bisa testing dengna sistem yang lain tanpa run sistem lain nya.
		contohnya sistem payment gateway. apalagi kalau bukan wewenang kita.
		Service -> Payment gateway / sms gateway
		jadi kita bisa bikin MOCK OBJECTNYA
	*/

}
