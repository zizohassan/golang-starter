package seeders

import (
	"golang-starter/app/models"
	"golang-starter/config"
	"strings"
)

/***
*	Seed Function must Have the same file Name then Add Seeder key word
* 	Example :  file is user function must be UserSeeder
 */
func (s *Seeder) TranslationSeeder() {
	pages := pages()
	languages := languages()
	for _, page := range pages {
		for _, lang := range languages {
			newTranslation(page, lang)
		}
	}
}

/**
* fake data and create data base
 */
func newTranslation(pageName string, lang string) {
	data := models.Translation{
		Slug:  pageName,
		Page:  strings.Title(pageName),
		Lang:  lang,
		Value: strings.Title(pageName),
	}
	config.DB.Create(&data)
}

func pages() []string {
	return []string{
		"home",
		"about",
		"contact",
		"terms",
		"police",
	}
}

func languages() []string {
	return []string{
		"en",
		"ar",
	}
}
