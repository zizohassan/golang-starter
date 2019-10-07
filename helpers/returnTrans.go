package helpers

import (
	"github.com/bykovme/gotrans"
	"github.com/gin-gonic/gin"
)

func DoneUpdate(g *gin.Context) string {
	return gotrans.Tr(GetCurrentLang(g), "done_update_item")
}

func DoneDelete(g *gin.Context) string {
	return gotrans.Tr(GetCurrentLang(g), "done_delete_item")
}

func DoneGetItem(g *gin.Context) string {
	return gotrans.Tr(GetCurrentLang(g), "done_get_item")
}

func DoneCreateItem(g *gin.Context) string {
	return gotrans.Tr(GetCurrentLang(g), "done_created_item")
}

func DoneGetAllItems(g *gin.Context) string {
	return gotrans.Tr(GetCurrentLang(g), "get_all_items")
}

func ItemNotFound(g *gin.Context) string {
	return gotrans.Tr(GetCurrentLang(g), "item_not_found")
}
