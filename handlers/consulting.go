package handlers

import (
	consultingdto "halocorona/dto/consulting"
	dto "halocorona/dto/result"
	"halocorona/models"
	"halocorona/repositories"
	"net/http"
	"strconv"
	"time"

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

	bornDate, err := time.Parse("2006-01-02", c.FormValue("bornDate"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: "err3"})
	}

	liveConsul, err := time.Parse("2006-01-04", c.FormValue("liveConsul"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: "err4"})
	}

	meminta := consultingdto.CreateConsultingRequest{
		UserId:         int(userId),
		BornDate:       bornDate,
		Age:            convAge,
		Height:         convHeight,
		Weight:         convWeight,
		Subject:        c.FormValue("subject"),
		LiveConsulting: liveConsul,
		Description:    c.FormValue("description"),
	}

	validation := validator.New()
	err = validation.Struct(meminta)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: "err1"})
	}

	consultation := models.Consulting{
		UserId:         meminta.UserId,
		BornDate:       meminta.BornDate.Format("2006-01-02"),
		Age:            meminta.Age,
		Height:         meminta.Height,
		Weight:         meminta.Weight,
		Subject:        meminta.Subject,
		LiveConsulting: meminta.LiveConsulting.Format("2006-01-04"),
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

func (h *handlerConsulting) SuccesConsulting(c echo.Context) error {
	var err error

	id, _ := strconv.Atoi(c.Param("id"))

	consultation, err := h.ConsultingRepository.DapatConsulting(uint(id))

	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	consultation.Status = "success"

	data, err := h.ConsultingRepository.EditConsulting(consultation)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccesResult{Code: http.StatusOK, Data: data})
}

func (h *handlerConsulting) RejectConsulting(c echo.Context) error {
	var err error

	id, _ := strconv.Atoi(c.Param("id"))

	consultation, err := h.ConsultingRepository.DapatConsulting(uint(id))

	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	consultation.Status = "reject"

	data, err := h.ConsultingRepository.EditConsulting(consultation)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}
	return c.JSON(http.StatusOK, dto.SuccesResult{Code: http.StatusOK, Data: data})
}
