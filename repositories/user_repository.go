package repositories

import (
	"errors"
	" hery-ciaputra/final-project-backend/httperror"
	" hery-ciaputra/final-project-backend/models"
	"gorm.io/gorm"
	"math"
	"time"
)

type UserRepository interface {
	MatchingCredential(email string) (*models.User, error)
	Profile(id int) (*models.User, error)
	UpdateProfile(id int, info *models.User) (*models.User, error)
	TopUp(info *models.Transaction) (*models.User, error)
	Address(id int) ([]*models.Address, error)
	CreateAddress(info *models.Address) (*models.Address, error)
	EditAddress(info *models.Address) (*models.Address, error)
	CreateShipping(info *models.Shipping) (*models.Shipping, error)
	ShippingList(id int) ([]*models.Shipping, error)
	Payment(info *models.Payment) (*models.Payment, error)
	AddPromo(info *models.Promo) (*models.Promo, error)
	PromoList() ([]*models.Promo, error)
	EditPromo(info *models.Promo) (*models.Promo, error)
	SizeList() ([]*models.Size, error)
	CategoryList() ([]*models.Category, error)
	AddOnList() ([]*models.AddOn, error)
}

type userRepository struct {
	db *gorm.DB
}

type URConfig struct {
	DB *gorm.DB
}

func NewUserRepository(c *URConfig) *userRepository {
	return &userRepository{db: c.DB}
}

func (u *userRepository) MatchingCredential(email string) (*models.User, error) {
	var user *models.User

	res := u.db.Where("email = ?", email).First(&user)

	isNotFound := errors.Is(res.Error, gorm.ErrRecordNotFound)
	if isNotFound {
		return nil, httperror.BadRequestError("Invalid email or password", "INVALID_LOGIN")
	}
	return user, nil
}

func (u *userRepository) Profile(id int) (*models.User, error) {
	var user *models.User

	res := u.db.Where("id = ?", id).First(&user)

	if res.Error != nil {
		return nil, res.Error
	}

	return user, nil
}

func (u *userRepository) UpdateProfile(id int, info *models.User) (*models.User, error) {
	var user *models.User

	err := u.db.Where("ID = ?", id).Updates(&info)
	if err.Error != nil {
		return nil, err.Error
	}

	err = u.db.Where("ID = ?", id).First(&user)
	if err.Error != nil {
		return nil, err.Error
	}

	return user, nil
}

func (u *userRepository) TopUp(info *models.Transaction) (*models.User, error) {
	var user *models.User

	res := u.db.Table("users").Where("id = ? ", info.ID).UpdateColumn("balance", gorm.Expr("balance + ?", info.Amount))
	if res.Error != nil {
		return nil, res.Error
	}

	u.db.Where("id = ? ", info.ID).First(&user)

	return user, nil
}

func (u *userRepository) Address(id int) ([]*models.Address, error) {
	var address []*models.Address
	var role *models.User

	res := u.db.Table("users").Select("role").Where("id = ?", id).First(&role)

	if role.Role == "user" {
		res = u.db.Table("addresses").Where("user_id = ?", id).Find(&address)
	} else {
		res = u.db.Table("addresses").Find(&address)
	}

	if res.Error != nil {
		return nil, res.Error
	}

	return address, nil
}

func (u *userRepository) CreateAddress(info *models.Address) (*models.Address, error) {
	res := u.db.Table("addresses").Create(&info)
	if err := res.Error; err != nil {
		return &models.Address{}, httperror.BadRequestError("Failed to create", "FAILED_TO_CREATE")
	}

	return info, res.Error
}

func (u *userRepository) EditAddress(info *models.Address) (*models.Address, error) {
	var address *models.Address

	err := u.db.Table("addresses").Where("ID = ? AND user_id = ?", info.ID, info.UserID).Updates(&info)
	if err.Error != nil {
		return nil, err.Error
	}

	u.db.Table("addresses").Where("ID = ? AND user_id = ?", info.ID, info.UserID).First(&address)

	return address, nil
}
func (u *userRepository) CreateShipping(info *models.Shipping) (*models.Shipping, error) {
	res := u.db.Preload("Size").Preload("Category").Preload("AddOn").Preload("Address").Table("shippings").Create(&info).First(&info)
	if res.Error != nil {
		return nil, res.Error
	}

	return info, nil
}

func (u *userRepository) ShippingList(id int) ([]*models.Shipping, error) {
	var shipping []*models.Shipping
	var role *models.User

	res := u.db.Table("users").Select("role").Where("id = ?", id).First(&role)

	if role.Role == "user" {
		res = u.db.Preload("Size").Preload("Category").Preload("AddOn").Preload("Address").Table("shippings").Joins("JOIN addresses ON addresses.id = shippings.address_id").Where("addresses.user_id = ?", id).Find(&shipping)
	} else {
		res = u.db.Preload("Size").Preload("Category").Preload("AddOn").Preload("Address").Table("shippings").Find(&shipping)
	}

	if res.Error != nil {
		return nil, res.Error
	}

	// todo handle error

	return shipping, nil
}

func (u *userRepository) Payment(info *models.Payment) (*models.Payment, error) {
	var shipping *models.Shipping
	var promo *models.Promo
	var amount float64
	var paid *models.Payment

	res := u.db.Table("payments").Where("shipping_id = ?", info.ShippingID).First(&paid)
	if paid.Status == "Paid" {
		return nil, httperror.BadRequestError("Item paid", "ITEM_PAID")
	}

	res = u.db.Preload("Size").Preload("Category").Preload("AddOn").Preload("Address").Table("shippings").Where("id = ?", info.ShippingID).First(&shipping)

	res = u.db.Table("promos").Where("id = ?", info.PromoID).First(&promo).UpdateColumn("quota", gorm.Expr("quota - 1"))

	amount = amount + float64(shipping.Size.Price) + float64(shipping.Category.Price) + float64(shipping.AddOn.Price)
	if promo.ExpiryDate.Before(time.Now()) || amount >= float64(promo.MinimumOrder) {
		amount = math.Max(amount-float64(promo.MaximumDiscount), amount*float64(1-promo.Discount))
	}

	info.Amount = int(amount)
	info.Status = "Paid"

	res = u.db.Table("payments").Create(&info)
	if res.Error != nil {
		return nil, res.Error
	}

	shipping.PaymentID = info.ID
	shipping.Status = "Packing"
	res = u.db.Table("shippings").Updates(&shipping)

	return info, nil
}

func (u *userRepository) AddPromo(info *models.Promo) (*models.Promo, error) {
	res := u.db.Table("promos").Create(&info)
	if res.Error != nil {
		return nil, res.Error
	}

	// todo handle error

	return info, nil
}

func (u *userRepository) PromoList() ([]*models.Promo, error) {
	var promo []*models.Promo

	res := u.db.Find(&promo)

	if res.Error != nil {
		return nil, res.Error
	}

	return promo, nil
}

func (u *userRepository) EditPromo(info *models.Promo) (*models.Promo, error) {
	var promo *models.Promo

	err := u.db.Table("promos").Where("ID = ?", info.ID).Updates(&info)
	if err.Error != nil {
		return nil, err.Error
	}

	u.db.Table("promos").Where("ID = ?", info.ID).First(&promo)

	return promo, nil
}

func (u *userRepository) SizeList() ([]*models.Size, error) {
	var size []*models.Size

	res := u.db.Find(&size)

	if res.Error != nil {
		return nil, res.Error
	}

	return size, nil
}

func (u *userRepository) CategoryList() ([]*models.Category, error) {
	var category []*models.Category

	res := u.db.Find(&category)

	if res.Error != nil {
		return nil, res.Error
	}

	return category, nil
}

func (u *userRepository) AddOnList() ([]*models.AddOn, error) {
	var addOn []*models.AddOn

	res := u.db.Find(&addOn)

	if res.Error != nil {
		return nil, res.Error
	}

	return addOn, nil
}
