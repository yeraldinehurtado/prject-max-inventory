package repository

import (
	context "context"
	entity "inventory/internal/entity"
)

const (
	qryInsertProduct = `
		insert into PRODUCTS (name, description, price, created_by) values (?, ?, ?, ?);
	`
	qryGetAllProducts = `
		select
			id,
			name,
			description,
			price,
			created_by
		from PRODUCTS
		where id = ?;
	`

	qryGetProductByID = `
	select
		id,
		name,
		description,
		price,
		created_by
	from PRODUCTS;
`
)

func (r *repo) SaveProduct(ctx context.Context, name, description string, price float32, createdBy int64) error {
	_, err := r.db.ExecContext(ctx, qryInsertProduct, name, description, price, createdBy)
	return err
}

func (r *repo) GetProducts(ctx context.Context) ([]entity.Product, error) {

	pp := []entity.Product{} // variables asi indica que es un arreglo

	err := r.db.SelectContext(ctx, &pp, qryGetAllProducts) // select es para tener multiples filas (rows)
	if err != nil {
		return nil, err
	}

	return pp, nil
}

func (r *repo) GetProduct(ctx context.Context, id int64) (*entity.Product, error) {
	p := &entity.Product{}

	err := r.db.GetContext(ctx, p, qryGetProductByID, id) // get es para obtener un report
	if err != nil {
		return nil, err
	}

	return p, nil
}
