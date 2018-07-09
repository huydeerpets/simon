package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["simon/controllers:AdminArticleController"] = append(beego.GlobalControllerRouter["simon/controllers:AdminArticleController"],
		beego.ControllerComments{
			Method: "Index",
			Router: `/admin`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["simon/controllers:AdminArticleController"] = append(beego.GlobalControllerRouter["simon/controllers:AdminArticleController"],
		beego.ControllerComments{
			Method: "ListArticles",
			Router: `/admin/article`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["simon/controllers:AdminArticleController"] = append(beego.GlobalControllerRouter["simon/controllers:AdminArticleController"],
		beego.ControllerComments{
			Method: "UpdateArticle",
			Router: `/admin/article/:id`,
			AllowHTTPMethods: []string{"get","post","delete"},
			Params: nil})

	beego.GlobalControllerRouter["simon/controllers:AdminArticleController"] = append(beego.GlobalControllerRouter["simon/controllers:AdminArticleController"],
		beego.ControllerComments{
			Method: "AddArticle",
			Router: `/admin/article/create`,
			AllowHTTPMethods: []string{"get","post"},
			Params: nil})

	beego.GlobalControllerRouter["simon/controllers:AdminCategoryController"] = append(beego.GlobalControllerRouter["simon/controllers:AdminCategoryController"],
		beego.ControllerComments{
			Method: "ListCategory",
			Router: `/admin/category`,
			AllowHTTPMethods: []string{"get","post"},
			Params: nil})

	beego.GlobalControllerRouter["simon/controllers:AdminCategoryController"] = append(beego.GlobalControllerRouter["simon/controllers:AdminCategoryController"],
		beego.ControllerComments{
			Method: "UpdateCategory",
			Router: `/admin/category/:id`,
			AllowHTTPMethods: []string{"post","put","delete"},
			Params: nil})

	beego.GlobalControllerRouter["simon/controllers:AdminConfigController"] = append(beego.GlobalControllerRouter["simon/controllers:AdminConfigController"],
		beego.ControllerComments{
			Method: "ConfigList",
			Router: `/admin/config`,
			AllowHTTPMethods: []string{"get","post"},
			Params: nil})

	beego.GlobalControllerRouter["simon/controllers:AdminConfigController"] = append(beego.GlobalControllerRouter["simon/controllers:AdminConfigController"],
		beego.ControllerComments{
			Method: "ConfigUpdate",
			Router: `/admin/config/:id`,
			AllowHTTPMethods: []string{"post","put","delete"},
			Params: nil})

	beego.GlobalControllerRouter["simon/controllers:AdminPictureController"] = append(beego.GlobalControllerRouter["simon/controllers:AdminPictureController"],
		beego.ControllerComments{
			Method: "ListPictures",
			Router: `/admin/picture`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["simon/controllers:AdminPictureController"] = append(beego.GlobalControllerRouter["simon/controllers:AdminPictureController"],
		beego.ControllerComments{
			Method: "EditPicture",
			Router: `/admin/picture/edit`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["simon/controllers:AdminPictureController"] = append(beego.GlobalControllerRouter["simon/controllers:AdminPictureController"],
		beego.ControllerComments{
			Method: "UpdatePicture",
			Router: `/admin/picture/:id`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["simon/controllers:AdminPictureController"] = append(beego.GlobalControllerRouter["simon/controllers:AdminPictureController"],
		beego.ControllerComments{
			Method: "UploadPicture",
			Router: `/admin/picture/upload`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["simon/controllers:AdminPictureController"] = append(beego.GlobalControllerRouter["simon/controllers:AdminPictureController"],
		beego.ControllerComments{
			Method: "UploadMarkdownPicture",
			Router: `/admin/markdown/upload`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["simon/controllers:AdminUserController"] = append(beego.GlobalControllerRouter["simon/controllers:AdminUserController"],
		beego.ControllerComments{
			Method: "Login",
			Router: `/admin/login`,
			AllowHTTPMethods: []string{"get","post"},
			Params: nil})

	beego.GlobalControllerRouter["simon/controllers:AdminUserController"] = append(beego.GlobalControllerRouter["simon/controllers:AdminUserController"],
		beego.ControllerComments{
			Method: "Logout",
			Router: `/admin/logout`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["simon/controllers:AdminUserController"] = append(beego.GlobalControllerRouter["simon/controllers:AdminUserController"],
		beego.ControllerComments{
			Method: "ListUsers",
			Router: `/admin/user`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["simon/controllers:AdminUserController"] = append(beego.GlobalControllerRouter["simon/controllers:AdminUserController"],
		beego.ControllerComments{
			Method: "UpdateUser",
			Router: `/admin/user/:id`,
			AllowHTTPMethods: []string{"get","post"},
			Params: nil})

	beego.GlobalControllerRouter["simon/controllers:ApiController"] = append(beego.GlobalControllerRouter["simon/controllers:ApiController"],
		beego.ControllerComments{
			Method: "GetCategoryList",
			Router: `/category/list`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["simon/controllers:ApiController"] = append(beego.GlobalControllerRouter["simon/controllers:ApiController"],
		beego.ControllerComments{
			Method: "GetArticleList",
			Router: `/article/list`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["simon/controllers:ArticleController"] = append(beego.GlobalControllerRouter["simon/controllers:ArticleController"],
		beego.ControllerComments{
			Method: "GetInfo",
			Router: `/article/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["simon/controllers:AvController"] = append(beego.GlobalControllerRouter["simon/controllers:AvController"],
		beego.ControllerComments{
			Method: "GetAv",
			Router: `/av/getav`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["simon/controllers:CategoryController"] = append(beego.GlobalControllerRouter["simon/controllers:CategoryController"],
		beego.ControllerComments{
			Method: "ListArticle",
			Router: `/category/:title`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["simon/controllers:MainController"] = append(beego.GlobalControllerRouter["simon/controllers:MainController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["simon/controllers:SearchController"] = append(beego.GlobalControllerRouter["simon/controllers:SearchController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/search`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["simon/controllers:WorksController"] = append(beego.GlobalControllerRouter["simon/controllers:WorksController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/works`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

}
