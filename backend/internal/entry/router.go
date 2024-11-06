package entry

import (
	"github.com/codeableorg/weekend-challenge-13-GabNatali/api"
	"github.com/gin-gonic/gin"
)

func AddEntryRoutes(router *gin.Engine, cases EntryUsesCases) {

	handler := NewEntryHandler(cases)

	router.Use(api.AuthenticateSession())

	router.POST("/entries", handler.CreateEntry)
	router.GET("/entries/user/:id", handler.GetEntriesbyUser)
	router.GET("/entries/:id", handler.GetEntry)
	router.DELETE("/entries/:id", handler.DeleteEntry)
	router.PATCH("/entries/:id", handler.UpdateEntry)

}
