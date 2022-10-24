package main

import (
	_middlewares "simple-blog/app/middlewares"
	
	_driverFactory "simple-blog/drivers"
	_dbDriver "simple-blog/drivers/mysql"

	_blogUseCase "simple-blog/businesses/blogs"
	_categoryUseCase "simple-blog/businesses/categories"
	_userUseCase "simple-blog/businesses/users"
	
	_blogController "simple-blog/controllers/blogs"
	_categoryController "simple-blog/controllers/categories"
	_userController "simple-blog/controllers/users"
	
	
	"simple-blog/util"
	_routes "simple-blog/app/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	configDB := _dbDriver.ConfigDB{
		DB_USERNAME: util.GetConfig("DB_USERNAME"),
		DB_PASSWORD: util.GetConfig("DB_PASSWORD"),
		DB_HOST: util.GetConfig("DB_HOST"),
		DB_PORT: util.GetConfig("DB_PORT"),
		DB_NAME: util.GetConfig("DB_NAME"),
	}

	db := configDB.InitDB()

	_dbDriver.DBMigrate(db)

	configJWT := _middlewares.ConfigJWT{
		SecretJWT: util.GetConfig("JWT_SECRET_KEY"),
		ExpiresDuration: 1,
	}

	configLogger := _middlewares.ConfigLogger{
		Format: "[${time_rfc3339}] ${status} ${method} ${host} ${path} ${latency_human}" + "\n",
	}

	app := echo.New()

	categoryRepo := _driverFactory.NewCategoryRepository(db)
	categoryUseCase := _categoryUseCase.NewCategoryUsecase(categoryRepo)
	categoryController := _categoryController.NewCategoryController(categoryUseCase)

	blogRepo := _driverFactory.NewBlogRepository(db)
	blogUseCase := _blogUseCase.NewBlogUsecase(blogRepo)
	blogController := _blogController.NewBlogController(blogUseCase)

	userRepo := _driverFactory.NewUserRepository(db)
	userUseCase := _userUseCase.NewUserUsecase(userRepo, &configJWT)
	userController := _userController.NewAuthController(userUseCase)

	routesInit := _routes.ControllerList{
		LoggerMiddleware: configLogger.Init(),
		JWTMiddleware: configJWT.Init(),
		CategoryController: *categoryController,
		BlogController: *blogController,
		AuthController: *userController,
	}

	routesInit.RouteRegister(app)

	app.Logger.Fatal(app.Start(":8000"))
}