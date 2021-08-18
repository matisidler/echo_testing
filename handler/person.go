package handler

import (
	"api_persons_echo/model"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

type person struct {
	storage Storage
}

func newPerson(storage Storage) person {
	return person{storage}
}

func (p *person) create(e echo.Context) error {
	data := model.Person{}
	err := e.Bind(&data)
	if err != nil {
		response := newResponse(Error, "person's struct isn't correct", nil)
		return e.JSON(http.StatusBadRequest, response)
	}
	err = p.storage.Create(&data)
	if err != nil {
		response := newResponse(Error, "a problem occurred when creating person", nil)
		return e.JSON(http.StatusInternalServerError, response)
	}
	response := newResponse(Message, "person created succesfuly", data)
	return e.JSON(http.StatusCreated, response)
}

func (p *person) getAll(e echo.Context) error {
	data, err := p.storage.GetAll()
	if err != nil {
		response := newResponse(Error, "a problem occurred while obtaining all persons", nil)

		return e.JSON(http.StatusInternalServerError, response)
	}
	response := newResponse(Message, "ok", data)
	return e.JSON(http.StatusOK, response)
}

func (p *person) update(e echo.Context) error {
	id, err := strconv.Atoi(e.Param("id"))
	if err != nil {
		response := newResponse(Error, "ID must be a number", nil)
		return e.JSON(http.StatusBadRequest, response)
	}
	data := model.Person{}
	err = e.Bind(&data)
	if err != nil {
		response := newResponse(Error, "person's struct isn't correct", nil)
		return e.JSON(http.StatusBadRequest, response)
	}
	err = p.storage.Update(id, &data)
	if err != nil {
		response := newResponse(Error, "a problem occurred while updating person", nil)
		return e.JSON(http.StatusInternalServerError, response)
	}

	response := newResponse(Message, "person updated succesfully", data)
	return e.JSON(http.StatusOK, response)
}

func (p *person) delete(e echo.Context) error {
	ID, err := strconv.Atoi(e.Param("id"))
	if err != nil {
		response := newResponse(Error, "ID not valid", nil)
		return e.JSON(http.StatusBadRequest, response)
	}

	err = p.storage.Delete(ID)
	if err != nil {
		response := newResponse(Error, "a problem occurred while deleting person", nil)
		return e.JSON(http.StatusInternalServerError, response)
	}
	response := newResponse(Message, "person deleted succesfully", nil)
	return e.JSON(http.StatusOK, response)
}

func (p *person) getById(e echo.Context) error {
	id, err := strconv.Atoi(e.Param("id"))
	if err != nil {
		response := newResponse(Error, "ID not valid", nil)
		return e.JSON(http.StatusBadRequest, response)
	}
	data, err := p.storage.GetById(id)
	if err != nil {
		response := newResponse(Error, "a problem occurred while getting person", nil)
		return e.JSON(http.StatusInternalServerError, response)
	}
	response := newResponse(Message, "person obtained succesfully", data)
	return e.JSON(http.StatusOK, response)
}
