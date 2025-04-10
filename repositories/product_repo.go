package repositories

import (
	"errors"
	"inventory/db"
	"inventory/models"

	"github.com/jmoiron/sqlx"
)

type ProductRepo struct {
	db *sqlx.DB
}

func NewProductRepo() *ProductRepo {
	return &ProductRepo{db: db.GetConnection()}
}

func (r *ProductRepo) ListProducts() ([]models.Product, error) {

	var products []models.Product

	query :=
		`
    SELECT p.id, p.name, p.quantity, p.price, 
    c.id AS "category.id",
    c.name AS "category.name"
    FROM product p
    JOIN category c ON c.id=p.category_id
    `

	err := r.db.Select(&products, query)

	if err != nil {
		return nil, err
	}

	return products, nil
}

func (r *ProductRepo) GetProduct(id uint64) (*models.Product, error) {

	var product models.Product

	query :=
		`
    SELECT p.id, p.name, p.quantity, p.price,
    c.id AS "category.id",
    c.name AS "category.name"

    FROM product p
    JOIN category c ON c.id=p.category_id
    WHERE p.id = ?
    `

	err := r.db.Get(&product, query, id)

	if err != nil {
		return nil, err
	}

	return &product, nil

}

func (r *ProductRepo) SaveProduct(product *models.Product) (*models.Product, error) {

	query :=
		`
        INSERT INTO product (name,quantity,price,category_id)
        VALUES(:name,:quantity,:price,:category_id)
    `

	response, err := r.db.NamedExec(query, &product)

	if err != nil {
		return nil, err
	}

	lastId, err := response.LastInsertId()
	if err != nil {
		return nil, err
	}

	return r.GetProduct(uint64(lastId))
}

func (r *ProductRepo) UpdateProduct(id uint64, product *models.Product) (*models.Product, error) {

	query :=
		`
        UPDATE product SET 
        name=?,
        quantity=?,
        price=?,
        category_id=?
        WHERE id =?
    `

	response, err := r.db.Exec(query,
		product.Name,
		product.Quantity,
		product.Price,
		product.CategoryID,
		id,
	)

	if err != nil {
		return nil, err
	}

	affectedRow, err := response.RowsAffected()
	if err != nil {
		return nil, err
	}

	if affectedRow == 0 {
		return nil, errors.New("Nothing got updated, check your ID")

	}

	return r.GetProduct(id)
}

func (r *ProductRepo) DeleteProduct(id uint64) error {

	query :=
		`
        DELETE FROM product
        where id = ?
    `

	response, err := r.db.Exec(query, id)

	if err != nil {
		return err
	}

	affectedRow, err := response.RowsAffected()
	if err != nil {
		return err
	}

	if affectedRow == 0 {
		return errors.New("Nothing got deleted, check your ID")

	}

	return nil
}
