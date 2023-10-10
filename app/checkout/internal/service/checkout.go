package service

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	pb "github.com/panupakm/boutique-go/api/checkout"
	spb "github.com/panupakm/boutique-go/api/shared"
	"github.com/panupakm/boutique-go/app/checkout/internal/biz"
	"github.com/panupakm/boutique-go/pkg/order"
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

	outOrder := spb.OrderResult{}
	order.ToProtoResult(&orderResult, &outOrder)
	return &pb.PlaceOrderResponse{
		Order: &outOrder,
	}, nil
}
