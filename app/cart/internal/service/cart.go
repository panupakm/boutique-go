package service

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"

	pb "github.com/panupakm/boutique-go/api/cart"
	spb "github.com/panupakm/boutique-go/api/shared"
	"github.com/panupakm/boutique-go/app/cart/internal/biz"
)

type CartServiceService struct {
	pb.UnimplementedCartServiceServer

	uc  *biz.CartUsecase
	log *log.Helper
}

func NewCartServiceService(uc *biz.CartUsecase, logger log.Logger) *CartServiceService {
	return &CartServiceService{
		uc:  uc,
		log: log.NewHelper(log.With(logger, "module", "cart/service")),
	}
}

func (s *CartServiceService) AddItem(ctx context.Context, req *pb.AddItemRequest) (*spb.Empty, error) {
	err := s.uc.AddItem(ctx, req.UserId, &biz.CartItem{
		ProductId: req.Item.GetProductId(),
		Quantity:  req.Item.GetQuantity(),
	})
	if err != nil {
		s.log.Errorf("AddItem: %w", err)
	}
	return &spb.Empty{}, err
}

func (s *CartServiceService) GetCart(ctx context.Context, req *pb.GetCartRequest) (*spb.Cart, error) {
	cart, err := s.uc.GetCart(ctx, req.UserId)
	if err != nil {
		s.log.Errorf("GetCart: %w", err)
	}

	var returnCart = spb.Cart{
		Items:  make([]*spb.CartItem, 0),
		UserId: cart.UserId,
	}

	for _, item := range cart.Items {
		returnCart.Items = append(returnCart.Items, &spb.CartItem{
			ProductId: item.ProductId,
			Quantity:  item.Quantity,
		})
	}

	return &returnCart, nil
}

func (s *CartServiceService) EmptyCart(ctx context.Context, req *pb.EmptyCartRequest) (*spb.Empty, error) {
	err := s.uc.Empty(ctx, req.UserId)
	return &spb.Empty{}, err
}
