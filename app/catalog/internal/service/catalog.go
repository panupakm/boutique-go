package service

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	pb "github.com/panupakm/boutique-go/api/catalog"
	spb "github.com/panupakm/boutique-go/api/shared"
	"github.com/panupakm/boutique-go/app/catalog/internal/biz"
	pkgproduct "github.com/panupakm/boutique-go/pkg/product"
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

func (s *CatalogService) ListProducts(ctx context.Context, req *pb.ListProductsRequest) (*pb.ListProductsResponse, error) {
	ps, err := s.uc.ListProducts(ctx, int(req.GetPageSize()), req.GetPageToken())
	if err != nil {
		return nil, err
	}

	products := make([]*spb.Product, 0, len(ps))
	for _, p := range ps {
		products = append(products, &spb.Product{
			Id:          p.Id,
			Name:        p.Name,
			Description: p.Description,
			Picture:     p.Picture,
			Categories:  p.Categories,
		})
	}

	return &pb.ListProductsResponse{
		Products: products,
	}, nil
}

func (s *CatalogService) GetProduct(ctx context.Context, req *pb.GetProductRequest) (*spb.Product, error) {
	p, err := s.uc.GetProduct(ctx, req.GetId())
	if err != nil {
		return nil, err
	}
	var res spb.Product
	pkgproduct.ToProto(&p, &res)

	return &res, nil
}

func (s *CatalogService) SearchProducts(ctx context.Context, req *pb.SearchProductsRequest) (*pb.SearchProductsResponse, error) {
	products, err := s.uc.SearchProducts(ctx, req.GetQuery(), int(req.GetPageSize()), req.GetPageToken())
	if err != nil {
		return nil, err
	}

	if len(products) == 0 {
		return &pb.SearchProductsResponse{}, nil
	}
	resProducts := make([]*spb.Product, 0, len(products))
	for _, p := range products {
		var res spb.Product
		pkgproduct.ToProto(&p, &res)
		resProducts = append(resProducts, &res)
	}
	return &pb.SearchProductsResponse{
		Results: resProducts,
	}, nil
}
