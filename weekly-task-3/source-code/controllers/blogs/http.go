package blogs

import (
	"net/http"
	"regexp"
	"simple-blog/businesses/blogs"
	controller "simple-blog/controllers"
	"simple-blog/controllers/blogs/request"
	"simple-blog/controllers/blogs/response"

	"github.com/labstack/echo/v4"
)

type BlogController struct {
	blogUseCase blogs.Usecase
}

func NewBlogController(blogUC blogs.Usecase) *BlogController {
	return &BlogController{
		blogUseCase: blogUC,
	}
}

// get all and get blog with name
func (bc *BlogController) GetAll(c echo.Context) error {
	blogs := []response.Blog{}

	title := c.QueryParam("keyword")
	regexProcessing := regexp.MustCompile(`[-]`)
	blogTitleQuery := regexProcessing.ReplaceAllString(title, " ") 

	blogsData := bc.blogUseCase.GetAll(blogTitleQuery)

	for _, blog := range blogsData {
		blogs = append(blogs, response.FromDomain(blog))
	}
	
	if title != "" {
		return controller.NewResponse(c, http.StatusOK, "success", "blog found", blogs)
	}

	return controller.NewResponse(c, http.StatusOK, "success", "all blogs", blogs)
}

func (bc *BlogController) GetByID(c echo.Context) error {
	id := c.Param("id")
	blog := bc.blogUseCase.GetByID(id)

	if blog.ID == 0 {
		return controller.NewResponse(c, http.StatusNotFound, "failed", "blog not found", "")
	}

	return controller.NewResponse(c, http.StatusOK, "success", "blog found", response.FromDomain(blog))
}

func (bc *BlogController) GetByCategoryID(c echo.Context) error {
	categoryId := c.Param("category_id")
	blogs := []response.Blog{}
	blogsData := bc.blogUseCase.GetByCategoryID(categoryId)

	if categoryId == "" {
		return controller.NewResponse(c, http.StatusNotFound, "failed", "blog not found", "")
	}

	for _, blog := range blogsData {
		blogs = append(blogs, response.FromDomain(blog))
	}

	return controller.NewResponse(c, http.StatusOK, "success", "blogs found", blogs)
}

func (bc *BlogController) Create(c echo.Context) error {
	input := request.Blog{}

	if err := c.Bind(&input); err != nil {
		return controller.NewResponse(c, http.StatusBadRequest, "failed", "validation failed", "")
	}

	err := input.Validate()

	if err != nil {
		return controller.NewResponse(c, http.StatusBadRequest, "failed", "validation failed", "")
	}

	blog := bc.blogUseCase.Create(input.ToDomain())

	return controller.NewResponse(c, http.StatusCreated, "success", "blog created", response.FromDomain(blog))
}

func (bc *BlogController) Update(c echo.Context) error {
	input := request.Blog{}

	if err := c.Bind(&input); err != nil {
		return controller.NewResponse(c, http.StatusBadRequest, "failed", "validation failed", "")
	}
	
	var blogId string = c.Param("id")
	err := input.Validate()

	if err != nil {
		return controller.NewResponse(c, http.StatusBadRequest, "failed", "validation failed", "")
	}

	blog := bc.blogUseCase.Update(blogId, input.ToDomain())

	if blog.ID == 0 {
		return controller.NewResponse(c, http.StatusNotFound, "failed", "blog not found", "")
	}

	return controller.NewResponse(c, http.StatusOK, "success", "blog updated", response.FromDomain(blog))
}

func (bc *BlogController) Delete(c echo.Context) error {
	blogId := c.Param("id")
	isSuccess := bc.blogUseCase.Delete(blogId)

	if !isSuccess {
		return controller.NewResponse(c, http.StatusNotFound, "failed", "blog not found", "")
	}

	return controller.NewResponse(c, http.StatusOK, "success", "blog deleted", "")
}