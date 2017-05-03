package handlers

import (
	"net/http"
	"strconv"

	"github.com/byuoitav/pulse-eight-neo-microservice/helpers"
	"github.com/labstack/echo"
)

func PowerOn(context echo.Context) error {
	return context.JSON(http.StatusOK, nil)
}

func Standby(context echo.Context) error {
	return context.JSON(http.StatusOK, nil)
}

func SwitchInput(context echo.Context) error {

	input := context.Param("input")
	output := context.Param("output")
	address := context.Param("address")

	err := helpers.SwitchInput(address, input, output)

	if err != nil {
		return context.JSON(http.StatusInternalServerError, err.Error())
	}

	return context.JSON(http.StatusOK, "Success")
}

func GetPower(context echo.Context) error {
	return context.JSON(http.StatusBadRequest, "Error")
}

func GetCurrentInput(context echo.Context) error {

	address := context.Param("address")
	inputs, err := helpers.GetCurrentInputs(address)
	if err != nil {
		return context.JSON(http.StatusInternalServerError, err.Error())
	}

	return context.JSON(http.StatusOK, inputs)
}

func GetInputByPort(context echo.Context) error {

	address := context.Param("address")
	port, err := strconv.Atoi(context.Param("port"))
	if err != nil {
		return context.JSON(http.StatusBadRequest, err.Error())
	}

	input, err := helpers.GetInputByOutputPort(address, port)
	if err != nil {
		return context.JSON(http.StatusInternalServerError, err.Error())
	}

	return context.JSON(http.StatusOK, input)
}
