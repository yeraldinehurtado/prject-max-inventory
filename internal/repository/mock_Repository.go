package repository

import (
	context "context"
	entity "inventory/internal/entity"

	mock "github.com/stretchr/testify/mock"
)

// MockRepository is an auto-generated mock type for the repository type
type MockRepository struct {
	mock.Mock
}

// GetUserByEmail provides a mock function with given fields ctx, email
/* func (_m *MockRepository) GetProduct(ctx context.Context, id int64) (*entity.Product, error) {
	ret := _m.Called(ctx, id)

	var r0 *entity.Product
	if rf, ok := ret.Get(0).(func(context.Context, int64) *entity.Product); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Product)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetProducts provides a mock function with given fields: ctx
func (_m *MockRepository) GetProducts(ctx context.Context) ([]entity.Product, error) {
	ret := _m.Called(ctx)

	var r0 []entity.Product
	if rf, ok := ret.Get(0).(func(context.Context) []entity.Product); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entity.Product)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
*/
// GetUserByEmail provides a mock function with given fields: ctx, email
func (_m *MockRepository) GetUserByEmail(ctx context.Context, email string) (*entity.User, error) {
	ret := _m.Called(ctx, email)

	var r0 *entity.User
	if rf, ok := ret.Get(0).(func(context.Context, string) *entity.User); ok {
		r0 = rf(ctx, email)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, email)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUserRoles provides a mock function with given fields: ctx, userID
func (_m *MockRepository) GetUserRoles(ctx context.Context, userID int64) ([]entity.UserRole, error) {
	ret := _m.Called(ctx, userID)

	var r0 []entity.UserRole
	if rf, ok := ret.Get(0).(func(context.Context, int64) []entity.UserRole); ok {
		r0 = rf(ctx, userID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entity.UserRole)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RemoveUserRole provides a mock function with given fields: ctx, userID, roleID
func (_m *MockRepository) RemoveUserRole(ctx context.Context, userID int64, roleID int64) error {
	ret := _m.Called(ctx, userID, roleID)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int64, int64) error); ok {
		r0 = rf(ctx, userID, roleID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

/*
// SaveProduct provides a mock function with given fields: ctx, name, description, price, createdBy
func (_m *MockRepository) SaveProduct(ctx context.Context, name string, description string, price float32, createdBy int64) error {
	ret := _m.Called(ctx, name, description, price, createdBy)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, float32, int64) error); ok {
		r0 = rf(ctx, name, description, price, createdBy)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
*/
// SaveUser provides a mock function with given fields: ctx, email, name, password
func (_m *MockRepository) SaveUser(ctx context.Context, email string, name string, password string) error {
	ret := _m.Called(ctx, email, name, password)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string) error); ok {
		r0 = rf(ctx, email, name, password)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SaveUserRole provides a mock function with given fields: ctx, userID, roleID
func (_m *MockRepository) SaveUserRole(ctx context.Context, userID int64, roleID int64) error {
	ret := _m.Called(ctx, userID, roleID)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int64, int64) error); ok {
		r0 = rf(ctx, userID, roleID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
