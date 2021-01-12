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

func (s *Server) loadRoute(ctx *fiber.Ctx) error {
	link := ctx.Query("link")
	if link == "" {
		return sendError(ctx, errors.New("incorrect resource link"), 404)
	}
	r, err := s.service.Load(ctx.Context(), link)
	if err != nil {
		return sendError(ctx, err, 404)
	}

	b, err := json.Marshal(r)
	if err != nil {
		return sendError(ctx, err, 500)
	}

	return ctx.Send(b)
}

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

	b, err := json.Marshal(r)
	if err != nil {
		return sendError(ctx, err, 500)
	}

	return ctx.Send(b)
}
