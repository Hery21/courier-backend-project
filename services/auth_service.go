package services

import (
	" hery-ciaputra/final-project-backend/config"
	" hery-ciaputra/final-project-backend/dto"
	" hery-ciaputra/final-project-backend/httperror"
	" hery-ciaputra/final-project-backend/models"
	" hery-ciaputra/final-project-backend/repositories"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"
)

type AuthService interface {
	LogIn(*dto.LogInReq) (*dto.TokenResponse, error)
	Profile(int) (*dto.ProfileRes, error)
	UpdateProfile(int, *dto.ProfileReq) (*dto.ProfileRes, error)
	TopUp(id int, req *dto.TopUpReq) (*dto.ProfileRes, error)
	Address(id int) ([]*dto.AddressRes, error)
	CreateAddress(id int, req *dto.AddressReq) (*dto.AddressRes, error)
	EditAddress(id int, req *dto.AddressReq) (*dto.AddressRes, error)
	CreateShipping(req *dto.ShippingReq) (*dto.ShippingRes, error)
	ShippingList(id int) ([]*dto.ShippingRes, error)
	Payment(id int, req *dto.PaymentReq) (*dto.PaymentRes, error)
	AddPromo(req *dto.PromoReq) (*dto.PromoRes, error)
	PromoList() ([]*dto.PromoRes, error)
	EditPromo(id int, req *dto.PromoReq) (*dto.PromoRes, error)
	SizeList() ([]*dto.SizeRes, error)
	CategoryList() ([]*dto.CategoryRes, error)
	AddOnList() ([]*dto.AddOnRes, error)
}

type authService struct {
	userRepository repositories.UserRepository
	appConfig      config.AppConfig
}

type AuthSConfig struct {
	UserRepository repositories.UserRepository
	AppConfig      config.AppConfig
}

func NewAuthService(c *AuthSConfig) AuthService {
	return &authService{
		userRepository: c.UserRepository,
		appConfig:      c.AppConfig,
	}
}

type idTokenClaims struct {
	jwt.RegisteredClaims
	User *models.User `json:"user"`
}

func (a *authService) generateJWTToken(user *models.User) (*dto.TokenResponse, error) {
	var idExp = a.appConfig.JWTExpireInMinutes * 60
	unixTime := time.Now().Unix()
	tokenExp := unixTime + idExp
	timeExpire := jwt.NumericDate{Time: time.Unix(tokenExp, 0)}
	timeNow := jwt.NumericDate{Time: time.Now()}

	claims := &idTokenClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    a.appConfig.AppName,
			IssuedAt:  &timeNow,
			ExpiresAt: &timeExpire,
		},
		User: user,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(a.appConfig.JWTSecret)

	if err != nil {
		return new(dto.TokenResponse), httperror.BadRequestError("BAD_REQUEST", "")
	}
	return &dto.TokenResponse{IDToken: tokenString}, nil
}

func (a *authService) LogIn(req *dto.LogInReq) (*dto.TokenResponse, error) {
	user, err := a.userRepository.MatchingCredential(req.Email)

	errNotMatch := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))

	if errNotMatch != nil || user == nil {
		return nil, httperror.AppError{
			StatusCode: http.StatusUnauthorized,
			Code:       "UNAUTHORIZED",
			Message:    "Unauthorized",
		}
	}
	token, err := a.generateJWTToken(user)
	return token, err
}

func (a *authService) Profile(id int) (*dto.ProfileRes, error) {
	user, err := a.userRepository.Profile(id)
	if err != nil {
		return new(dto.ProfileRes), err
	}

	return new(dto.ProfileRes).FromUser(user), nil
}

func (a *authService) UpdateProfile(id int, pr *dto.ProfileReq) (*dto.ProfileRes, error) {
	updateInfo := &models.User{
		Name:  pr.Name,
		Email: pr.Email,
		Phone: pr.Phone,
		Photo: pr.Photo,
	}

	user, err := a.userRepository.UpdateProfile(id, updateInfo)
	if err != nil {
		return new(dto.ProfileRes), err
	}

	return new(dto.ProfileRes).FromUser(user), nil
}

func (a *authService) TopUp(id int, req *dto.TopUpReq) (*dto.ProfileRes, error) {
	topUpInfo := &models.Transaction{
		ID:     id,
		Amount: req.Amount,
	}

	user, err := a.userRepository.TopUp(topUpInfo)
	if err != nil {
		return new(dto.ProfileRes), err
	}

	return new(dto.ProfileRes).FromUser(user), nil
}

func (a *authService) Address(id int) ([]*dto.AddressRes, error) {
	addresses, err := a.userRepository.Address(id)

	// todo handle error

	var addressList []*dto.AddressRes

	for i := range addresses {
		addressList = append(addressList, new(dto.AddressRes).FromAddress(addresses[i]))
	}

	return addressList, err
}

func (a *authService) CreateAddress(id int, req *dto.AddressReq) (*dto.AddressRes, error) {
	addressInfo := &models.Address{
		UserID:    id,
		Recipient: req.Recipient,
		Address:   req.Address,
		Phone:     req.Phone,
	}

	res, err := a.userRepository.CreateAddress(addressInfo)
	if err != nil {
		return new(dto.AddressRes), err
	}

	return new(dto.AddressRes).FromAddress(res), nil
}

func (a *authService) EditAddress(id int, req *dto.AddressReq) (*dto.AddressRes, error) {
	addressInfo := &models.Address{
		UserID:    id,
		ID:        req.ID,
		Recipient: req.Recipient,
		Address:   req.Address,
		Phone:     req.Phone,
	}

	res, err := a.userRepository.EditAddress(addressInfo)
	if err != nil {
		return new(dto.AddressRes), err
	}

	return new(dto.AddressRes).FromAddress(res), nil
}

func (a *authService) CreateShipping(req *dto.ShippingReq) (*dto.ShippingRes, error) {
	shippingInfo := &models.Shipping{
		SizeID:     req.SizeID,
		CategoryID: req.CategoryID,
		AddOnID:    req.AddOnID,
		AddressID:  req.AddressID,
		Status:     "On Cart",
	}

	ship, err := a.userRepository.CreateShipping(shippingInfo)
	if err != nil {
		return new(dto.ShippingRes), err
	}

	return new(dto.ShippingRes).FromShipping(ship), nil
}

func (a *authService) ShippingList(id int) ([]*dto.ShippingRes, error) {
	shipping, err := a.userRepository.ShippingList(id)

	// todo handle error

	var shippingList []*dto.ShippingRes

	for i := range shipping {
		shippingList = append(shippingList, new(dto.ShippingRes).FromShipping(shipping[i]))
	}

	return shippingList, err
}

func (a *authService) Payment(id int, req *dto.PaymentReq) (*dto.PaymentRes, error) {
	paymentInfo := &models.Payment{
		ShippingID: id,
		PromoID:    req.PromoID,
	}

	payment, err := a.userRepository.Payment(paymentInfo)
	if err != nil {
		return new(dto.PaymentRes), err
	}

	return new(dto.PaymentRes).FromPayment(payment), nil
}

func (a *authService) AddPromo(req *dto.PromoReq) (*dto.PromoRes, error) {
	promoInfo := &models.Promo{
		Name:            req.Name,
		MinimumOrder:    req.MinimumOrder,
		Discount:        req.Discount,
		MaximumDiscount: req.MaximumDiscount,
		Quota:           req.Quota,
		ExpiryDate:      req.ExpiryDate,
	}

	promo, err := a.userRepository.AddPromo(promoInfo)
	if err != nil {
		return new(dto.PromoRes), err
	}

	return new(dto.PromoRes).FromPromo(promo), nil
}

func (a *authService) PromoList() ([]*dto.PromoRes, error) {
	promos, err := a.userRepository.PromoList()

	// todo handle error

	var promoList []*dto.PromoRes

	for i := range promos {
		promoList = append(promoList, new(dto.PromoRes).FromPromo(promos[i]))
	}

	return promoList, err
}

func (a *authService) EditPromo(id int, req *dto.PromoReq) (*dto.PromoRes, error) {
	promoInfo := &models.Promo{
		ID:              id,
		Name:            req.Name,
		MinimumOrder:    req.MinimumOrder,
		Discount:        req.Discount,
		MaximumDiscount: req.MaximumDiscount,
		Quota:           req.Quota,
		ExpiryDate:      req.ExpiryDate,
	}

	res, err := a.userRepository.EditPromo(promoInfo)
	if err != nil {
		return new(dto.PromoRes), err
	}

	return new(dto.PromoRes).FromPromo(res), nil
}

func (a *authService) SizeList() ([]*dto.SizeRes, error) {
	sizes, err := a.userRepository.SizeList()

	// todo handle error

	var sizeList []*dto.SizeRes

	for i := range sizes {
		sizeList = append(sizeList, new(dto.SizeRes).FromSize(sizes[i]))
	}

	return sizeList, err
}

func (a *authService) CategoryList() ([]*dto.CategoryRes, error) {
	categories, err := a.userRepository.CategoryList()

	// todo handle error

	var categoryList []*dto.CategoryRes

	for i := range categories {
		categoryList = append(categoryList, new(dto.CategoryRes).FromCategory(categories[i]))
	}

	return categoryList, err
}

func (a *authService) AddOnList() ([]*dto.AddOnRes, error) {
	addOns, err := a.userRepository.AddOnList()

	// todo handle error

	var addOnList []*dto.AddOnRes

	for i := range addOns {
		addOnList = append(addOnList, new(dto.AddOnRes).FromAddOn(addOns[i]))
	}

	return addOnList, err
}
