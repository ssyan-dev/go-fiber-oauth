package response

import "github.com/gofiber/fiber/v2"

func SuccessResponse(ctx *fiber.Ctx, code int, message string, data interface{}) error {
	return sendResponse(ctx, code, message, data, true)
}

func ErrorResponse(ctx *fiber.Ctx, code int, message string) error {
	return sendResponse(ctx, code, message, nil, false)
}

func sendResponse(ctx *fiber.Ctx, code int, message string, data interface{}, isSuccess bool) error {
	return ctx.Status(code).JSON(fiber.Map{
		"success": isSuccess,
		"message": message,
		"data":    data,
	})
}
