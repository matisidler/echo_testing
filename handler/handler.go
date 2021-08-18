package handler

import (
	"api_persons_echo/model"
	
)

//le vamos a decir que todos aquellos sistemas de almacenamiento que quieran trabajar con nuestro handler deben cumplir con los siguientes métodos:
type Storage interface {
	Create(person *model.Person) error
	Update(id int, person *model.Person) error
	Delete(id int) error
	GetById(id int) (model.Person, error)
	GetAll() ([]model.Person, error)
}
