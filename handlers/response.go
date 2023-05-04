package handlers

import (
	responsedto "halocorona/dto/response"
	dto "halocorona/dto/result"
	"halocorona/models"
	"halocorona/repositories"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

type handlerResponse struct {
	ResponseRepository repositories.ResponseRepository
}

func HandlerResponse(ResponseRepository repositories.ResponseRepository) *handlerResponse {
	return &handlerResponse{ResponseRepository}
}

func (h *handlerResponse) MembuatResponse(c echo.Context) error {
	userLogin := c.Get("userLogin")
	idUser := userLogin.(jwt.MapClaims)["id"].(float64)

	consultId, _ := strconv.Atoi(c.Param("id"))

	request := responsedto.ResponseRequest{
		ResponseText: c.FormValue("responseText"),
		ConsulLink:   c.FormValue("consulLink"),
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: "err1"})
	}

	response := models.Response{
		UserId:       int(idUser),
		ConsulId:     consultId,
		ResponseText: request.ResponseText,
		ConsulLink:   request.ConsulLink,
	}

	response, err = h.ResponseRepository.MembuatResponse(response)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: "err2"})
	}

	response, _ = h.ResponseRepository.DapatResponse(response.ID)

	return c.JSON(http.StatusOK, dto.SuccesResult{Code: http.StatusOK, Data: response})
}

func (h *handlerResponse) DapatResponse(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	response, err := h.ResponseRepository.DapatResponse(uint(id))
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccesResult{Code: http.StatusOK, Data: response})
}
