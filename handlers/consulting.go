package handlers

import (
	consultingdto "halocorona/dto/consulting"
	dto "halocorona/dto/result"
	"halocorona/models"
	"halocorona/repositories"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

type handlerConsulting struct {
	ConsultingRepository repositories.ConsultingRepository
}

func HandlerConsulting(ConsultingRepository repositories.ConsultingRepository) *handlerConsulting {
	return &handlerConsulting{ConsultingRepository}
}

func (h *handlerConsulting) DapatConsulting(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	article, err := h.ConsultingRepository.DapatConsulting(uint(id))
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccesResult{Code: http.StatusOK, Data: article})
}

func (h *handlerConsulting) CariConsultingKu(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	consultations, err := h.ConsultingRepository.CariConsultingKu(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccesResult{Code: http.StatusOK, Data: consultations})
}

func (h *handlerConsulting) DapatConsul(c echo.Context) error {

	consultations, err := h.ConsultingRepository.DapatConsul()
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}
	return c.JSON(http.StatusOK, dto.SuccesResult{Code: http.StatusOK, Data: consultations})
}

func (h *handlerConsulting) MembuatConsulting(c echo.Context) error {
	userLogin := c.Get("userLogin")

	userId := userLogin.(jwt.MapClaims)["id"].(float64)

	convAge, _ := strconv.Atoi(c.FormValue("age"))
	convHeight, _ := strconv.Atoi(c.FormValue("height"))
	convWeight, _ := strconv.Atoi(c.FormValue("weight"))
	meminta := consultingdto.CreateConsultingRequest{
		UserId:         int(userId),
		BornDate:       c.FormValue("bornDate"),
		Age:            convAge,
		Height:         convHeight,
		Weight:         convWeight,
		Gender:         c.FormValue("gender"),
		Subject:        c.FormValue("subject"),
		LiveConsulting: c.FormValue("liveConsul"),
		Description:    c.FormValue("description"),
	}

	validation := validator.New()
	err := validation.Struct(meminta)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: "err1"})
	}

	consultation := models.Consulting{
		UserId:         meminta.UserId,
		BornDate:       meminta.BornDate,
		Age:            meminta.Age,
		Height:         meminta.Height,
		Weight:         meminta.Weight,
		Gender:         meminta.Gender,
		Subject:        meminta.Subject,
		LiveConsulting: meminta.LiveConsulting,
		Description:    meminta.Description,
		Status:         "waiting",
	}

	consultation, err = h.ConsultingRepository.MembuatConsulting(consultation)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: "err2"})
	}

	consultation, _ = h.ConsultingRepository.DapatConsulting(consultation.Id)

	return c.JSON(http.StatusOK, dto.SuccesResult{Code: http.StatusOK, Data: consultation})
}
