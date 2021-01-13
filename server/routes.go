package server

import (
	"encoding/json"
	"errors"
	"github.com/gofiber/fiber/v2"
)

func sendError(ctx *fiber.Ctx, err error, code int) error {
	_ = ctx.SendStatus(code)
	return ctx.SendString(err.Error())
}

// @Summary Get image by filename
// @Description Get image by filename
// @Tags load
// @Accept json
// @Produce image/png, image/jpg
// @Param filename path string true "Image name"
// @Success 200 {string} string
// @Failure 404 {string} string
// @Failure 500 {string} string
// @Router /{filename} [get]
func (s *Server) loadServingRoot() string {
	return s.servingFolder
}

// @Summary Upload single image and get his name
// @Description Upload single image and get his name
// @Tags load
// @Accept json
// @Produce text/plain
// @Param link query string true "Image url for uploading"
// @Success 200 {string} string
// @Failure 404 {string} string
// @Failure 500 {string} string
// @Router /load [get]
func (s *Server) loadRoute(ctx *fiber.Ctx) error {
	link := ctx.Query("link")
	if link == "" {
		return sendError(ctx, errors.New("incorrect resource link"), 404)
	}
	r, err := s.service.Load(ctx.Context(), link)
	if err != nil {
		return sendError(ctx, err, 404)
	}

	return ctx.JSON(r)
}

// @Summary Loading multiple images and getting its name as a map
// @Description Loading multiple images and getting its name as a map
// @Tags load
// @Accept json
// @Produce json
// @Param images body []string true "Pass images urls"
// @Success 200 {object} map[string]string
// @Failure 404 {string} string
// @Failure 500 {string} string
// @Router /load/batch [post]
func (s *Server) loadBatchRoute(ctx *fiber.Ctx) error {
	var links []string

	err := json.Unmarshal(ctx.Body(), &links)
	if err != nil  {
		return sendError(ctx, err, 404)
	}

	r, err := s.service.LoadBatch(ctx.Context(), links)
	if err != nil {
		return sendError(ctx, err, 404)
	}

	return ctx.JSON(r)
}
