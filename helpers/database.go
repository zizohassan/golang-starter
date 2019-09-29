package helpers

import "golang-starter/config"

/***
* truncate tables
*/
func DbTruncate(tableName ...string) {
	for _, table := range tableName {
		config.DB.Exec("TRUNCATE " + table)
	}
}
