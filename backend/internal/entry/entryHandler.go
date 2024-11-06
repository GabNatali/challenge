package entry

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type EntryHandler interface {
	CreateEntry(c *gin.Context)
	UpdateEntry(c *gin.Context)
	DeleteEntry(c *gin.Context)
	GetEntry(c *gin.Context)
	GetEntriesbyUser(c *gin.Context)
}

type entryHandler struct {
	entryUsesCases EntryUsesCases
}

func NewEntryHandler(cases EntryUsesCases) EntryHandler {
	return &entryHandler{
		entryUsesCases: cases,
	}
}

// AddEntry implements EntryHandler.
func (e *entryHandler) CreateEntry(c *gin.Context) {

	var dtoEntry AddEntryDto

	if err := c.ShouldBindJSON(&dtoEntry); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	entry, err := e.entryUsesCases.Create(dtoEntry)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Entry added successfully",
		"data":    entry,
	})
}

// DeleteEntry implements EntryHandler.
func (e *entryHandler) DeleteEntry(c *gin.Context) {

	id := c.Param("id")

	idn, err := strconv.Atoi(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	_, err = e.entryUsesCases.Delete(uint(idn))

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Entry deleted successfully",
	})
}

// GetEntries implements EntryHandler.
func (e *entryHandler) GetEntriesbyUser(c *gin.Context) {

	id := c.Param("id")
	limit := c.Query("limit")
	offset := c.Query("offset")

	idn, err := strconv.Atoi(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	limitInt, err := strconv.Atoi(limit)

	if err != nil {
		limitInt = 10
	}
	offsetInt, err := strconv.Atoi(offset)

	if err != nil {
		offsetInt = 0
	}

	entries, err := e.entryUsesCases.GetAllEntriesByUserId(uint(idn), limitInt, offsetInt)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Entries retrieved successfully",
		"data":    entries,
	})
}

// GetEntry implements EntryHandler.
func (e *entryHandler) GetEntry(c *gin.Context) {
	id := c.Param("id")

	idn, err := strconv.Atoi(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	entry, err := e.entryUsesCases.GetEntryById(uint(idn))

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Entry retrieved successfully",
		"data":    entry,
	})
}

// UpdateEntry implements EntryHandler.
func (e *entryHandler) UpdateEntry(c *gin.Context) {

	id := c.Param("id")

	idn, err := strconv.Atoi(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	var dtoEntry UpdateEntryDto

	if err := c.ShouldBindJSON(&dtoEntry); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	entry, err := e.entryUsesCases.Update(uint(idn), dtoEntry)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Entry updated successfully",
		"data":    entry,
	})

}
