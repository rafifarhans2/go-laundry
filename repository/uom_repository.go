package repository

import (
	"database/sql"
	"enigma-laundry/models"
)

type UomRepository interface {
	Save(uom models.Uom) error              // INSERT
	FindById(id string) (models.Uom, error) // SELECT by id
	FindAll() ([]models.Uom, error)         // SELECT *
	Update(uom models.Uom) error            // UPDATE
	DeleteById(id string)                   // DELETE by id
}

type uomRepository struct {
	db *sql.DB
}

// DeleteById implements UomRepository.
func (u *uomRepository) DeleteById(id string) {
	panic("unimplemented")
}

// FindAll implements UomRepository.
func (u *uomRepository) FindAll() ([]models.Uom, error) {
	panic("unimplemented")
}

// FindById implements UomRepository.
func (u *uomRepository) FindById(id string) (models.Uom, error) {
	panic("unimplemented")
}

// Save implements UomRepository.
func (u *uomRepository) Save(uom models.Uom) error {
	_, err := u.db.Exec("INSERT INTO uom VALUES ($1, $2)", uom.Id, uom.Name)
	if err != nil {
		return err
	}
	return nil
}

// Update implements UomRepository.
func (u *uomRepository) Update(uom models.Uom) error {
	panic("unimplemented")
}

// Constructor
func NewUomRepository(db *sql.DB) UomRepository {
	return &uomRepository{db: db}
}
