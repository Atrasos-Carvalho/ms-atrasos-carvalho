package http

import (
	"ms-atrasos-carvalho/internal/domain"
	"ms-atrasos-carvalho/internal/usecase"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type RegistroHandler struct {
	usecase *usecase.RegistroUseCase
}

func NewRegistroHandler(uc *usecase.RegistroUseCase) *RegistroHandler {
	return &RegistroHandler{usecase: uc}
}

func (h *RegistroHandler) Create(c *gin.Context) {
	var input struct {
		TipoID int `json:"tipo_id"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid body"})
		return
	}

	data := domain.Date{Time: time.Now()}

	existing, err := h.usecase.FindByDate(data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if len(existing) > 0 {
		c.JSON(http.StatusConflict, gin.H{"error": "registro para hoje já existe"})
		return
	}

	reg := &domain.Registro{
		Data: data,
		TipoID: input.TipoID,
	}

	if err := h.usecase.Create(reg); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "criado"})
}

func (h *RegistroHandler) FindAll(c *gin.Context) {
	registros, err := h.usecase.FindAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, registros)
}

func (h *RegistroHandler) FindCount(c *gin.Context) {
	count, err := h.usecase.FindCount()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, count)
}