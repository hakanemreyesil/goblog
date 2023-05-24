package config

import (
	admin "goblog/admin/controllers"
	site "goblog/site/controllers"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func Routes() *httprouter.Router {
	r := httprouter.New()
	//Admin
	r.GET("/admin", admin.Dashboard{}.Index)

	//Blog Posts
	r.GET("/admin/new-add", admin.Dashboard{}.NewItem)
	r.POST("/admin/add", admin.Dashboard{}.Add)
	r.GET("/admin/delete/:id", admin.Dashboard{}.Delete)
	r.GET("/admin/edit/:id", admin.Dashboard{}.Edit)
	r.POST("/admin/update/:id", admin.Dashboard{}.Update)

	//Categoriler
	r.GET("/admin/categories", admin.Categories{}.Index)
	r.POST("/admin/categories/add", admin.Categories{}.Add)
	r.GET("/admin/categories/delete/:id", admin.Categories{}.Delete)

	//Userops
	r.GET("/admin/login", admin.Userops{}.Index)
	r.POST("/admin/do_login", admin.Userops{}.Login)
	r.GET("/admin/logout", admin.Userops{}.Logout)

	//Site
	r.GET("/", site.Homepage{}.Index)
	r.GET("/blogs/:slug", site.Homepage{}.Detail)

	//Serve Files
	r.ServeFiles("/admin/assets/*filepath", http.Dir("admin/assets"))
	r.ServeFiles("/assets/*filepath", http.Dir("site/assets"))
	r.ServeFiles("/uploads/*filepath", http.Dir("uploads"))
	return r
}
