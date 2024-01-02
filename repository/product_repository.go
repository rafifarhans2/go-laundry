package repository

import (
	"database/sql"
	"enigma-laundry/models"
)

type ProductRepository interface {
	Save(product models.Product) error
	FindById(id string) (models.Product, error)
	FindAll() ([]models.Product, error)
	FindByName(name string) ([]models.Product, error)
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
	row, err := p.db.Query(`SELECT p.id, p.name, p.price, u.id, u.name FROM product p JOIN uom u ON u.id = p.uom_id`)
	if err != nil {
		return nil, err
	}
	var products []models.Product
	for row.Next() {
		product := models.Product{}
		err := row.Scan(
			&product.Id,
			&product.Name,
			&product.Price,
			&product.Uom.Id,
			&product.Uom.Name,
		)
		if err != nil {
			return nil, err
		}
	}
	return products, nil
}

// FindById implements ProductRepository.
func (p *productRepository) FindById(id string) (models.Product, error) {
	row := p.db.QueryRow(`SELECT p.id, p.name, p.price, u.id, u.name FROM product p JOIN uom u ON u.id = p.uom_id WHERE id = $1`, id)
	product := models.Product{}
	err := row.Scan(
		&product.Id,
		&product.Name,
		&product.Price,
		&product.Uom.Id,
		&product.Uom.Name,
	)
	if err != nil {
		return models.Product{}, err
	}
	return product, nil
}

// FindByName implements ProductRepository.
func (p *productRepository) FindByName(name string) ([]models.Product, error) {
	row, err := p.db.Query(`SELECT p.id, p.name, p.price, u.id, u.name FROM product p JOIN uom u ON u.id = p.uom_id WHERE name ILIKE '%$1%'`, name)
	if err != nil {
		return nil, err
	}
	var products []models.Product
	for row.Next() {
		product := models.Product{}
		err := row.Scan(
			&product.Id,
			&product.Name,
			&product.Price,
			&product.Uom.Id,
			&product.Uom.Name,
		)
		if err != nil {
			return nil, err
		}
	}
	return products, nil
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
	_, err := p.db.Exec("UPDATE product SET name = $2, price = $3, uom_id = $4 WHERE id = $1", product.Id, product.Name, product.Price, product.Uom.Id)
	if err != nil {
		return err
	}
	return nil
}

func NewProductRepository(db *sql.DB) ProductRepository {
	return &productRepository{db: db}
}
