package service

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	pb "github.com/panupakm/boutique-go/api/checkout"
	spb "github.com/panupakm/boutique-go/api/shared"
	"github.com/panupakm/boutique-go/app/checkout/internal/biz"
)

type CheckoutService struct {
	pb.UnimplementedCheckoutServer

	cuc *biz.CheckoutUseCase
	log *log.Helper
}

func NewCheckoutService(cuc *biz.CheckoutUseCase, logger log.Logger) *CheckoutService {
	return &CheckoutService{
		cuc: cuc,
		log: log.NewHelper(log.With(logger, "module", "checkout/service")),
	}
}

func (s *CheckoutService) PlaceOrder(ctx context.Context, req *pb.PlaceOrderRequest) (*pb.PlaceOrderResponse, error) {
	orderResult, err := s.cuc.PlaceOrder(ctx, req.GetUserId())
	if err != nil {
		s.log.Errorf("PlaceOrder: %w", err)
		return nil, err
	}

	resOrderResult := spb.OrderResult{
		OrderId:            orderResult.OrderId,
		ShippingTrackingId: orderResult.ShippingTrackingId,
		ShippingCost: &spb.Money{
			CurrencyCode: orderResult.ShippingCost.CurrencyCode,
			Units:        orderResult.ShippingCost.Units,
			Nanos:        orderResult.ShippingCost.Nanos,
		},
		ShippingAddress: &spb.Address{
			StreetAddress: orderResult.ShippingAddress.StreetAddress,
			City:          orderResult.ShippingAddress.City,
			State:         orderResult.ShippingAddress.State,
			Country:       orderResult.ShippingAddress.Country,
			ZipCode:       orderResult.ShippingAddress.ZipCode,
		},
	}

	return &pb.PlaceOrderResponse{
		Order: &resOrderResult,
	}, nil
}
