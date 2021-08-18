package storage

import (
	"api_persons_echo/model"
	"fmt"
)

type Memory struct {
	currentID int
	Persons   map[int]model.Person
}

func NewMemory() Memory {
	persons := make(map[int]model.Person)
	return Memory{
		currentID: 0,
		Persons:   persons,
	}
}

func (m *Memory) Create(person *model.Person) error {
	if person == nil {
		return model.ErrPersonCanNotBeNil
	}
	m.currentID++
	m.Persons[m.currentID] = *person

	return nil
}

//Recibimos el ID a actualizar en la base de datos y el modelo de persona nuevo que vamos a introducir.
func (m *Memory) Update(id int, person *model.Person) error {
	if person == nil {
		return model.ErrPersonCanNotBeNil
	}
	//Busca si hay una persona con ese ID en el mapa.
	if _, ok := m.Persons[int(id)]; !ok {
		return fmt.Errorf("ID: %d: %w", id, model.ErrIDPersonDoesNotExists)
	}

	//le decimos que en ese ID, ponga a la nueva persona.
	m.Persons[int(id)] = *person
	return nil
}

func (m *Memory) Delete(id int) error {
	//chequeamos si la persona con el ID recibido existe:
	if _, ok := m.Persons[int(id)]; !ok {
		return fmt.Errorf("ID: %d: %w", id, model.ErrIDPersonDoesNotExists)
	}

	//si existe, lo borramos.
	delete(m.Persons, int(id))
	return nil
}

func (m *Memory) GetById(id int) (model.Person, error) {
	//buscamos en el mapa el ID. Si no lo encuentra, devolvemos un error informando que esa persona no existe.
	person, ok := m.Persons[id]
	if !ok {
		return person, fmt.Errorf("ID: %d, %w", id, model.ErrIDPersonDoesNotExists)
	}
	return person, nil
}

func (m *Memory) GetAll() ([]model.Person, error) {
	var persons []model.Person
	for _, person := range m.Persons {
		persons = append(persons, person)
	}
	return persons, nil
}
