package service

import (
	context "context"
	"errors"
	"inventory/internal/models"
)

var validRolesToAddProduct []int64 = []int64{1, 2}
var ErrInvalidPermissions = errors.New("user does not have permission to add product")

func (s *serv) GetProducts(ctx context.Context) ([]models.Product, error) {

	pp, err := s.repo.GetProducts(ctx)
	if err != nil {
		return nil, err
	}

	products := []models.Product{}

	for _, p := range pp { // recorremos productos
		products = append(products, models.Product{
			ID:          p.ID,
			Name:        p.Name,
			Description: p.Description,
			Price:       p.Price,
		})

	}

	return products, nil
}

func (s *serv) GetProduct(ctx context.Context, id int64) (*models.Product, error) {

	p, err := s.repo.GetProduct(ctx, id)
	if err != nil {
		return nil, err
	}

	product := &models.Product{
		ID:          p.ID,
		Name:        p.Name,
		Description: p.Description,
		Price:       p.Price,
	}

	return product, nil
}

func (s *serv) AddProdcut(ctx context.Context, product models.Product, email string) error {

	u, err := s.repo.GetUserByEmail(ctx, email)
	if err != nil {
		return err
	} // obtenemos la info del usuario

	roles, err := s.repo.GetUserRoles(ctx, u.ID)
	if err != nil {
		return err
	} // obtenemos cuales son los roles por usuario

	userCanAdd := false

	for _, r := range roles {
		for _, vr := range validRolesToAddProduct {
			if vr == r.RoleID {
				userCanAdd = true

			}
		}
	} // verifica si al menos uno de los roles permite agregar productos

	if !userCanAdd {
		return ErrInvalidPermissions
	} // si no es ninguno de los roles que tiene permisos de agregar productos, retorna este error

	return s.repo.SaveProduct(
		ctx,
		product.Name,
		product.Description,
		product.Price,
		u.ID,
	)
}
