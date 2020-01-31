package models

func MigrateAllTable(path string) {
	AnswerMigrate()
	CategoryMigrate()
	FaqMigrate()
	PageMigrate()
	PageImageMigrate()
	SettingMigrate()
	TranslationMigrate()
	UserMigrate()
}
