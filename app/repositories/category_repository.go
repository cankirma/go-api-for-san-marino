package repositories

import (
	"database/sql"
	"github.com/cankirma/go-api-for-san-marino/app/dtos"
	"github.com/cankirma/go-api-for-san-marino/app/entities"
	"github.com/jmoiron/sqlx"
)

type CategoryRepository struct {
	*sqlx.DB
}

func (cr CategoryRepository) GetAllCategories() ([]entities.Category, error) {

	var category []entities.Category

	query := `SELECT * FROM categories`

	err := cr.Select(&category, query)
	if err != nil {

		return category, err
	}

	return category, nil
}

func (cr CategoryRepository) GetCategoryById(id int) (entities.Category, error) {
	category := entities.Category{}
	query := `SELECT * FROM categories where category_id = $1`

	err := cr.Get(&category, query, id)
	if err != nil {
		return entities.Category{}, err
	}
	return category, nil
}

func (cr CategoryRepository) InsertCategory(category *dtos.CreateCategoryModel) error {
	query := `INSERT INTO 
    public.categories(category_id,category_name, description)
	VALUES ($1,$2,$3);`
	_, err := cr.Exec(query, category.CategoryId, category.CategoryName, category.Description)
	if err != nil {
		return err
	}
	return nil
}

func (cr CategoryRepository) UpdateCategory(id int, model *dtos.UpdateCategoryModel) error {

	query := `UPDATE public.categories
    SET
	category_name= $2, description= $3
	WHERE category_id = $1;`
	//_,err := cr.Exec(query,id,model.CategoryName,model.Description)
	//if err != nil{
	//	return err
	//}
	//return nil

	stmt, err := cr.Prepare(query)
	if err != nil {
		return err
	}
	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {

		}
	}(stmt)
	stmt.Exec(id, model.CategoryName,model.Description)
	return nil
}
func (cr CategoryRepository) DeleteCategory(id int) error  {
	query := `DELETE FROM 
            public.categories
			WHERE category_id = $1;`
	 _ ,err :=cr.Exec(query,id)
	if err != nil{
		return err
	}
	return nil
}