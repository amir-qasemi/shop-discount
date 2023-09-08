package discount

import (
	"database/sql"
)

type Repository interface {
	GetDiscountByCode(string) (Discount, error)
	GetDiscountByUsageId(string) (Discount, error)
	Save(Discount) error
	Delete(Discount) error
}

func NewDummyRepository(db *sql.DB) DummyRepository {
	// Populate based on config
	return DummyRepository{}
}

type DummyRepository struct {
	// db *gorm.Db

}

func (r *DummyRepository) GetDiscountByCode(code string) (Discount, error) {
	return nil, nil
}

func (r *DummyRepository) GetDiscountByUsageId(code string) (Discount, error) {
	return nil, nil
}

func (r *DummyRepository) Save(discount Discount) error {
	return nil
}

func (r *DummyRepository) Delete(discount Discount) error {
	return nil
}
