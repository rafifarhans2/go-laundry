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
	DeleteById(id string) error             // DELETE by id
}

type uomRepository struct {
	db *sql.DB
}

// DeleteById implements UomRepository.
func (u *uomRepository) DeleteById(id string) error {
	if _, err := u.db.Exec("DELETE FROM uom WHERE id = $1", id); err != nil {
		return err
	}
	return nil
}

// FindAll implements UomRepository.
func (u *uomRepository) FindAll() ([]models.Uom, error) {
	row, err := u.db.Query("SELECT * FROM uom")

	var uom []models.Uom

	if err != nil {
		return nil, err
	}
	for row.Next() {
		var uoms models.Uom
		if err := row.Scan(&uoms.Id, &uoms.Name); err != nil {
			return nil, err
		}
		uom = append(uom, uoms)
	}
	return uom, nil
}

// FindById implements UomRepository.
func (u *uomRepository) FindById(id string) (models.Uom, error) {
	row := u.db.QueryRow("SELECT id, name FROM uom WHERE id = $1", id)
	var uom models.Uom
	if err := row.Scan(&uom.Id, &uom.Name); err != nil {
		return models.Uom{}, err
	}
	return uom, nil
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
	_, err := u.db.Exec("UPDATE uom SET name = $2 WHERE id = $1;", uom.Id, uom.Name)
	if err != nil {
		return err
	}
	return nil
}

// Constructor
func NewUomRepository(db *sql.DB) UomRepository {
	return &uomRepository{db: db}
}
