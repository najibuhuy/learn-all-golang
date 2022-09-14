package service

import (
	"testing"
	"unit-test/entity"
	"unit-test/repository"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

//initiate value Mock
var categoryRepository = &repository.CategoryRepositoryMock{Mock: mock.Mock{}}

//
var categoryService = CategoryService{Repository: categoryRepository}

func TestCategoryService_Get(t *testing.T) {
	//define response when FindById id="1" mock
	categoryRepository.Mock.On("FindById", "1").Return(nil)

	//call response from mock
	category, err := categoryService.Get("1")

	//if Not nill get error
	assert.NotNil(t, err)

	//if nill get success
	assert.Nil(t, category)
}

func TestCategoryService_GetFound(t *testing.T) {
	category := entity.Category{
		Id:   "2",
		Name: "Pasar",
	}
	categoryRepository.Mock.On("FindById", "2").Return(category)

	result, err := categoryService.Get("2")
	assert.Nil(t, err)
	//check result is not nil or nil
	assert.NotNil(t, result)
	assert.Equal(t, category.Id, result.Id)
	assert.Equal(t, category.Name, result.Name)

}
