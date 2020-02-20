package seeders

import (
	"golang-starter/app/models"
	"golang-starter/config"
)

/***
*	Seed Function must Have the same file Name then Add Seeder key word
* 	Example :  file is user function must be UserSeeder
 */
func ActionSeeder() {
	///// define default actions
	var defaultActions = make(map[string]interface{})
	/// define modules
	var modules = make(map[string]interface{})
	defaultActions["all"] = models.ALL
	////////////// start modules have all action only
	modules["settings"] = defaultActions
	///////////// end modules has all action only
	defaultActions["active"] =  models.ACTIVE
	////////////// start modules have all , active actions
	////////////// end modules have all , active  actions
	defaultActions["deactivate"] = models.DEACTIVATE
	////////////// start modules have all , active , deactivate actions
	////////////// end modules have all , active , deactivate  actions
	defaultActions["trashed"] = models.TRASH
	////////////// start modules have all , active , deactivate , trash actions
	modules["pages"] = defaultActions
	modules["faqs"] = defaultActions
	modules["categories"] = defaultActions
	////////////// end modules have all , active , deactivate , trash  actions
	defaultActions["blocked"] = models.BLOCK
	modules["users"] = defaultActions
	/// loop to create actions
	for module , actions := range modules {
		loop := actions.(map[string]interface{})
		for noun , verb := range loop {
			newAction(models.Action{
				Noun:       noun,
				Verb:       verb.(string),
				Count:      0,
				Slug:       verb.(string) + "_" + module,
				ModuleName: module,
			})
		}
	}
}

/**
* fake data and create data base
 */
func newAction(data models.Action) {
	config.DB.Create(&data)
}
