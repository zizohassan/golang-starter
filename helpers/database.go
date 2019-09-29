package helpers

import "starter/config"

/***
* truncate tables
*/
func DbTruncate(tableName ...string) {
	for _, table := range tableName {
		config.DB.Exec("TRUNCATE " + table)
	}
}
