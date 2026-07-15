package response

import "github.com/gin-gonic/gin"

type SuccessResponse struct {
	Status  string      `json:"status" example:"success"`
	Message string      `json:"message" example:"Success to collect data"`
	Data    interface{} `json:"data,omitempty"`
}

type ErrorResponse struct {
	Status  string `json:"status" example:"error"`
	Message string `json:"message" example:"There is an error"`
	Error   string `json:"error,omitempty" example:"Detail error in here"`
}

type ListResponse struct {
	Status  string      `json:"status" example:"success"`
	Message string      `json:"message" example:"Success to collect data"`
	Data    interface{} `json:"data"`
}

func OK(c *gin.Context, message string, data interface{}) {
	c.JSON(200, SuccessResponse{
		Status:  "success",
		Message: message,
		Data:    data,
	})
}

func Created(c *gin.Context, message string, data interface{}) {
	c.JSON(201, SuccessResponse{
		Status:  "success",
		Message: message,
		Data:    data,
	})
}

func List(c *gin.Context, message string, data interface{}) {
	c.JSON(200, ListResponse{
		Status:  "success",
		Message: message,
		Data:    data,
	})
}

func BadRequest(c *gin.Context, message string, err error) {
	resp := ErrorResponse{
		Status:  "error",
		Message: message,
	}
	if err != nil {
		resp.Error = err.Error()
	}
	c.JSON(400, resp)
}

func NotFound(c *gin.Context, message string) {
	c.JSON(404, ErrorResponse{
		Status:  "error",
		Message: message,
	})
}

func InternalError(c *gin.Context, message string, err error) {
	resp := ErrorResponse{
		Status:  "error",
		Message: message,
	}
	if err != nil {
		resp.Error = err.Error()
	}
	c.JSON(500, resp)
}
