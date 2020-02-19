package providers

import (
	"github.com/vardius/gocontainer"
	"golang-starter/app/models"
	"golang-starter/config"
)

var Container gocontainer.Container

func StartContainer() {
	Container = gocontainer.New()
	/// register all models
	Container.Register("user", models.User{})
	Container.Register("category", models.Category{})
	Container.Register("answer", models.Answer{})
	Container.Register("faq", models.Faq{})
	Container.Register("page", models.Page{})
	Container.Register("pageImage", models.PageImage{})
	Container.Register("setting", models.Setting{})
	Container.Register("translation", models.Translation{})
	Container.Register("db", config.DB)
	/// register any thing you want on container
}
