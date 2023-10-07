package service

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	pb "github.com/panupakm/boutique-go/api/catalog"
	spb "github.com/panupakm/boutique-go/api/shared"
	"github.com/panupakm/boutique-go/app/catalog/internal/biz"
)

type CatalogService struct {
	pb.UnimplementedCatalogServer

	uc  *biz.CatalogUsecase
	log *log.Helper
}

func NewCatalogService(uc *biz.CatalogUsecase, logger log.Logger) *CatalogService {
	return &CatalogService{
		uc:  uc,
		log: log.NewHelper(log.With(logger, "module", "catalog/service")),
	}
}

func (s *CatalogService) ListProducts(ctx context.Context, req *spb.Empty) (*pb.ListProductsResponse, error) {
	return &pb.ListProductsResponse{}, nil
}

func (s *CatalogService) GetProduct(ctx context.Context, req *pb.GetProductRequest) (*pb.Product, error) {
	p, err := s.uc.GetProduct(ctx, req.GetId())
	if err != nil {
		return nil, err
	}

	return &pb.Product{
		Id:          p.Id,
		Name:        p.Name,
		Description: p.Description,
		Picture:     p.Picture,
		PriceUsd: &spb.Money{
			CurrencyCode: p.PriceUsd.CurrencyCode,
			Units:        p.PriceUsd.Units,
			Nanos:        p.PriceUsd.Nanos,
		},
		Categories: p.Categories,
	}, nil
}

func (s *CatalogService) SearchProducts(ctx context.Context, req *pb.SearchProductsRequest) (*pb.SearchProductsResponse, error) {
	return &pb.SearchProductsResponse{}, nil
}
