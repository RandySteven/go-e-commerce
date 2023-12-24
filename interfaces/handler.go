package interfaces

import "net/http"

type (
	UserHandler interface {
		RegisterUser(res http.ResponseWriter, req *http.Request)
		LoginUser(res http.ResponseWriter, req *http.Request)
		TestGetUser(res http.ResponseWriter, req *http.Request)
	}

	ShopHandler interface {
		RegisterShop(res http.ResponseWriter, req *http.Request)
		LoginShop(res http.ResponseWriter, req *http.Request)
		ShopDetail(res http.ResponseWriter, req *http.Request)
	}

	ProductHandler interface {
		CreateProduct(res http.ResponseWriter, req *http.Request)
		GetAllProducts(res http.ResponseWriter, req *http.Request)
	}

	CategoryHandler interface {
		AddCategory(res http.ResponseWriter, req *http.Request)
		GetAllCategories(res http.ResponseWriter, req *http.Request)
		GetCategoryById(res http.ResponseWriter, req *http.Request)
	}
)
