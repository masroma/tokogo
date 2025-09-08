package services

import (
	"errors"
	"tokogo/config"
	"tokogo/models"
	"tokogo/repositories"
	"tokogo/requests"
	"tokogo/responses"
)

type CartService struct {
	cartRepo    *repositories.CartRepository
	productRepo *repositories.ProductRepository
}

func NewCartService() *CartService {
	return &CartService{
		cartRepo:    repositories.NewCartRepository(config.DB),
		productRepo: repositories.NewProductRepository(config.DB),
	}
}

func (s *CartService) AddToCart(userID uint, req requests.AddToCartRequest) (*responses.CartItemResponse, error) {
	// Check if product exists
	product, err := s.productRepo.GetByID(req.ProductID)
	if err != nil {
		return nil, errors.New("product not found")
	}

	// Check if product is in stock
	if product.Stock < req.Quantity {
		return nil, errors.New("insufficient stock")
	}

	// Check if item already exists in cart
	existingCart, err := s.cartRepo.GetByUserIDAndProductID(userID, req.ProductID)
	if err == nil {
		// Item exists, update quantity
		existingCart.Quantity += req.Quantity

		// Check stock again after adding
		if product.Stock < existingCart.Quantity {
			return nil, errors.New("insufficient stock")
		}

		if err := s.cartRepo.Update(existingCart); err != nil {
			return nil, errors.New("failed to update cart")
		}

		response := responses.ConvertCartToResponse(*existingCart)
		return &response, nil
	}

	// Create new cart item
	cart := &models.Cart{
		UserID:    userID,
		ProductID: req.ProductID,
		Quantity:  req.Quantity,
	}

	if err := s.cartRepo.Create(cart); err != nil {
		return nil, errors.New("failed to add to cart")
	}

	// Get the created cart with product details
	createdCart, err := s.cartRepo.GetByUserIDAndProductID(userID, req.ProductID)
	if err != nil {
		return nil, errors.New("failed to retrieve cart item")
	}

	response := responses.ConvertCartToResponse(*createdCart)
	return &response, nil
}

func (s *CartService) GetCart(userID uint) (*responses.CartResponse, error) {
	carts, err := s.cartRepo.GetByUserID(userID)
	if err != nil {
		return nil, errors.New("failed to get cart")
	}

	response := responses.CreateCartResponse(carts)
	return &response, nil
}

func (s *CartService) UpdateCartItem(userID uint, productID uint, req requests.UpdateCartItemRequest) (*responses.CartItemResponse, error) {
	// Check if product exists
	product, err := s.productRepo.GetByID(productID)
	if err != nil {
		return nil, errors.New("product not found")
	}

	// Check if product is in stock
	if product.Stock < req.Quantity {
		return nil, errors.New("insufficient stock")
	}

	// Get existing cart item
	cart, err := s.cartRepo.GetByUserIDAndProductID(userID, productID)
	if err != nil {
		return nil, errors.New("cart item not found")
	}

	// Update quantity
	cart.Quantity = req.Quantity

	if err := s.cartRepo.Update(cart); err != nil {
		return nil, errors.New("failed to update cart")
	}

	response := responses.ConvertCartToResponse(*cart)
	return &response, nil
}

func (s *CartService) RemoveFromCart(userID uint, productID uint) error {
	// Check if cart item exists
	_, err := s.cartRepo.GetByUserIDAndProductID(userID, productID)
	if err != nil {
		return errors.New("cart item not found")
	}

	if err := s.cartRepo.DeleteByUserIDAndProductID(userID, productID); err != nil {
		return errors.New("failed to remove from cart")
	}

	return nil
}

func (s *CartService) ClearCart(userID uint) error {
	if err := s.cartRepo.ClearCart(userID); err != nil {
		return errors.New("failed to clear cart")
	}

	return nil
}

func (s *CartService) GetCartItemCount(userID uint) (int64, error) {
	count, err := s.cartRepo.GetCartItemCount(userID)
	if err != nil {
		return 0, errors.New("failed to get cart count")
	}

	return count, nil
}
