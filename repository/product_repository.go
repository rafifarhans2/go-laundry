package repository

import (
	"database/sql"
	"enigma-laundry/models"
)

type ProductRepository interface {
	Save(product models.Product) error
	FindById(id string) (models.Product, error)
	FindAll() ([]models.Product, error)
	Update(product models.Product) error
	DeleteById(id string) error
}

type productRepository struct {
	db *sql.DB
}

// DeleteById implements ProductRepository.
func (p *productRepository) DeleteById(id string) error {
	if _, err := p.db.Exec("DELETE FROM product WHERE id = $1", id); err != nil {
		return err
	}
	return nil
}

// FindAll implements ProductRepository.
func (p *productRepository) FindAll() ([]models.Product, error) {
	rows, err := p.db.Query("SELECT * FROM product")

	var product []models.Product

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var products models.Product
		if err := rows.Scan(&products.Id, &products.Name, &products.Price, &products.Uom.Id); err != nil {
			return nil, err
		}
		product = append(product, products)
	}
	return product, nil

}

// FindById implements ProductRepository.
func (p *productRepository) FindById(id string) (models.Product, error) {
	row := p.db.QueryRow("SELECT id, name FROM product WHERE id = $1", id)

	var product models.Product
	if err := row.Scan(&product.Id, &product.Name); err != nil {
		return models.Product{}, err
	}
	return product, nil

}

// Save implements ProductRepository.
func (p *productRepository) Save(product models.Product) error {
	_, err := p.db.Exec("INSERT INTO product VALUES ($1, $2, $3, $4)", product.Id, product.Name, product.Price, product.Uom.Id)
	if err != nil {
		return err
	}
	return nil
}

// Update implements ProductRepository.
func (p *productRepository) Update(product models.Product) error {
	_, err := p.db.Exec("UPDATE product SET name = $2 WHERE id = $1", product.Id, product.Name)
	if err != nil {
		return err
	}
	return nil
}

func NewProductRepository(db *sql.DB) ProductRepository {
	return &productRepository{db: db}
}
