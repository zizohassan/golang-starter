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
	attr := globalAttrs()
	for index, page := range pages {
		for _, lang := range languages {
			for _, at := range attr {
				newTranslation(index+1, page, lang, at)
			}
		}
	}
}

/**
* fake data and create data base
 */
func newTranslation(pageId int, pageName string, lang string, slug string) {
	data := models.Translation{
		Slug:   slug,
		PageId: pageId,
		Lang:   lang,
		Value:  strings.Title(pageName),
	}
	config.DB.Create(&data)
}

func languages() []string {
	return []string{
		"en",
		"ar",
	}
}

func globalAttrs() []string {
	return []string{
		"title",
		"des",
	}
}
