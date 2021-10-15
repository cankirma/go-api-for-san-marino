package repositories

import (
	"github.com/cankirma/go-api-for-san-marino/app/entities"
	"github.com/jmoiron/sqlx"
)

type ProductRepository struct {
	*sqlx.DB

}

func (pr ProductRepository) GetAllProducts() ([]entities.Product, error) {

	var products []entities.Product

	query := `SELECT * FROM products`

	err := pr.Select(&products, query)
	if err != nil {

		return nil, err
	}

	return products, nil
}

func (pr ProductRepository) GetProductById(id int) (entities.Product, error) {
	product := entities.Product{}
	var query = `SELECT * FROM Products where product_id = $1`
	err := pr.Get(&product, query, id)
	if err != nil {
		return entities.Product{}, err
	}
	return product, nil
}

func (pr ProductRepository) InsertProduct(p *entities.Product) error {
	var query = `INSERT INTO public.products(
    product_id,                        
	 product_name,
    supplier_id,
    category_id, 
    quantity_per_unit,
    unit_price,
    units_in_stock,
    units_on_order,
    reorder_level,
    discontinued)
	VALUES ($1 ,$2, $3, $4,$5, $6, $7, $8,$9, $10);`
	_, err := pr.Exec(query,
		p.ProductName,
		p.SupplierId,
		p.CategoryId,
		p.QuantityPerUnit,
		p.UnitPrice,
		p.UnitsInStock,
		p.UnitsOnOrder,
		p.ReOrderLevel,
		p.Discontinued)
	if err != nil {
		return err
	}
	return nil
}

func (pr ProductRepository) UpdateProduct( model *entities.Product) error {

	var query = `UPDATE public.products
	SET product_name=$1,
	    supplier_id=$2,
	    category_id=$3,
	    quantity_per_unit=$4,
	    unit_price=$5,
	    units_in_stock=$6,
	    units_on_order=$7, 
	    reorder_level=$8, 
	    discontinued=$9
	WHERE product_id = $10;`
	_,err := pr.Exec(query,
		model.ProductName,
		model.SupplierId,
		model.CategoryId,
		model.QuantityPerUnit,
		model.UnitPrice,
		model.UnitsInStock,
		model.UnitsOnOrder,
		model.ReOrderLevel,
		model.Discontinued,
	)
	if err != nil{
		return err
	}
	return nil

}
func (pr ProductRepository) DeleteProduct(id int) error {
	query := `DELETE FROM 
            public.products
			WHERE product_id= $1;`
	_, err := pr.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}