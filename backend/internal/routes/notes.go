package routes

import (
	"errors"
	"log/slog"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"

	"github.com/skilluv-community/starter-fullstack-go/backend/internal/models"
)

type NotesHandler struct {
	DB *sqlx.DB
}

func (h *NotesHandler) List(c *gin.Context) {
	notes := []models.Note{}
	if err := h.DB.SelectContext(c, &notes,
		`SELECT id, text, created_at FROM notes ORDER BY created_at DESC LIMIT 100`); err != nil {
		slog.Error("notes.list", "err", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, notes)
}

type createBody struct {
	Text string `json:"text" binding:"required,min=1,max=500"`
}

func (h *NotesHandler) Create(c *gin.Context) {
	var body createBody
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	text := strings.TrimSpace(body.Text)
	if text == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "text is required"})
		return
	}
	id := uuid.New()
	var out models.Note
	if err := h.DB.GetContext(c, &out,
		`INSERT INTO notes (id, text) VALUES ($1, $2) RETURNING id, text, created_at`, id, text); err != nil {
		slog.Error("notes.create", "err", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusCreated, out)
}

func (h *NotesHandler) Delete(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	res, err := h.DB.ExecContext(c, `DELETE FROM notes WHERE id = $1`, id)
	if err != nil {
		slog.Error("notes.delete", "err", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	rows, err := res.RowsAffected()
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, errors.New("rows affected"))
		return
	}
	if rows == 0 {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	c.Status(http.StatusNoContent)
}
