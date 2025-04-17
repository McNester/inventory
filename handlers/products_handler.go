package handlers

import (
	"context"
	"inventory/models"
	"inventory/services"

	pb "cloud_commons/inventory"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ProductHandler struct {
	service *services.ProductService
	pb.UnimplementedProductServiceServer
}

func NewProductHandler() *ProductHandler {
	return &ProductHandler{service: services.NewProductService()}
}

func (h *ProductHandler) Save(c context.Context, req *pb.Product) (*pb.Product, error) {

	product := models.Product{
		Id:         req.Id,
		Name:       req.Name,
		Quantity:   req.Quantity,
		Price:      req.Price,
		CategoryID: req.CategoryId,
		Category: models.Category{
			Id:   req.Category.Id,
			Name: req.Category.Name,
		},
	}

	savedProduct, err := h.service.SaveProduct(&product)

	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			"Faield to save the product %v",
			err.Error())
	}

	return &pb.Product{
		Id:         savedProduct.Id,
		Name:       savedProduct.Name,
		Quantity:   savedProduct.Quantity,
		Price:      savedProduct.Price,
		CategoryId: savedProduct.CategoryID,
		Category: &pb.Category{
			Id:   savedProduct.Category.Id,
			Name: savedProduct.Category.Name,
		},
	}, nil

}

func (h *ProductHandler) Update(c context.Context, req *pb.Product) (*pb.Product, error) {

	product := models.Product{
		Id:         req.Id,
		Name:       req.Name,
		Quantity:   req.Quantity,
		Price:      req.Price,
		CategoryID: req.CategoryId,
		Category: models.Category{
			Id:   req.Category.Id,
			Name: req.Category.Name,
		},
	}

	updatedProduct, err := h.service.UpdateProduct(product.Id, &product)

	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			"Faield to update the product %v",
			err.Error())
	}

	return &pb.Product{
		Id:         updatedProduct.Id,
		Name:       updatedProduct.Name,
		Quantity:   updatedProduct.Quantity,
		Price:      updatedProduct.Price,
		CategoryId: updatedProduct.CategoryID,
		Category: &pb.Category{
			Id:   updatedProduct.Category.Id,
			Name: updatedProduct.Category.Name,
		},
	}, nil

}

func (h *ProductHandler) Get(c context.Context, req *pb.ProductId) (*pb.Product, error) {
	product, err := h.service.GetProduct(req.Id)

	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			"Faield to get the product %v",
			err.Error())
	}

	return &pb.Product{
		Id:         product.Id,
		Name:       product.Name,
		Quantity:   product.Quantity,
		Price:      product.Price,
		CategoryId: product.CategoryID,
		Category: &pb.Category{
			Id:   product.Category.Id,
			Name: product.Category.Name,
		},
	}, nil

}

func (h *ProductHandler) List(req *pb.NoParams, stream pb.ProductService_ListServer) error {
	products, err := h.service.ListProducts()

	if err != nil {
		return status.Errorf(
			codes.Internal,
			"Faield to list the products %v",
			err.Error())
	}

	for _, product := range products {
		pbProduct := &pb.Product{
			Id:         product.Id,
			Name:       product.Name,
			Quantity:   product.Quantity,
			Price:      product.Price,
			CategoryId: product.CategoryID,
			Category: &pb.Category{
				Id:   product.Category.Id,
				Name: product.Category.Name,
			},
		}

		if err := stream.Send(pbProduct); err != nil {
			return status.Errorf(codes.Internal, "Failed to send product: %v", err.Error())
		}
	}

	return nil

}

func (h *ProductHandler) Delete(c context.Context, req *pb.ProductId) (*pb.DeleteResponse, error) {

	err := h.service.DeleteProduct(req.Id)

	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			"Faield to delete the product %v",
			err.Error())
	}

	return &pb.DeleteResponse{
		Message: "Successfully deleted the product",
	}, nil

}
