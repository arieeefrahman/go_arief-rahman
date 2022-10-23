package categories

import (
	"net/http"
	"simple-blog/businesses/categories"
	controller "simple-blog/controllers"
	"simple-blog/controllers/categories/request"
	"simple-blog/controllers/categories/response"

	"github.com/labstack/echo/v4"
)

type CategoryController struct {
	categoryUseCase categories.Usecase
}

func NewCategoryController(categoryUC categories.Usecase) *CategoryController {
	return &CategoryController{
		categoryUseCase: categoryUC,
	}
}

func (cc *CategoryController) GetAllCategories(c echo.Context) error {
	categories := []response.Category{}
	categoriesData := cc.categoryUseCase.GetAll()

	for _, category := range categoriesData {
		categories = append(categories, response.FromDomain(&category))
	}

	return controller.NewResponse(c, http.StatusOK, "success", "all categories", categories)
}

func (cc *CategoryController) CreateCategory(c echo.Context) error {
	input := request.Category{}

	if err := c.Bind(&input); err != nil {
		return controller.NewResponse(c, http.StatusBadRequest, "failed", "validation failed", "")
	}

	err := input.Validate()

	if err != nil {
		return controller.NewResponse(c, http.StatusBadRequest, "failed", "validation failed", "")
	}

	category := cc.categoryUseCase.Create(input.ToDomain())

	return controller.NewResponse(c, http.StatusCreated, "success", "category created", response.FromDomain(&category))
}

func (cc *CategoryController) UpdateCategory(c echo.Context) error {
	input := request.Category{}
	var id string = c.Param("id")

	if err := c.Bind(&input); err != nil {
		return controller.NewResponse(c, http.StatusBadRequest, "failed", "validation failed", "")
	}

	err := input.Validate()

	if err != nil {
		return controller.NewResponse(c, http.StatusBadRequest, "failed", "validation failed", "")
	}

	category := cc.categoryUseCase.Update(id, input.ToDomain())

	return controller.NewResponse(c, http.StatusOK, "success", "category updated", response.FromDomain(&category))
}

func (cc *CategoryController) DeleteCategory(c echo.Context) error {
	var id string = c.Param("id")

	isDeleted := cc.categoryUseCase.Delete(id)

	if !isDeleted {
		return controller.NewResponse(c, http.StatusNotFound, "failed", "category not found", "")
	}

	return controller.NewResponse(c, http.StatusOK, "success", "category deleted", "")
}