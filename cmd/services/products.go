package services

// import (
// 	"context"
// 	"fmt"
// 	"log"
// 	productPb "testing/pb/product"

// 	"go.mongodb.org/mongo-driver/bson"
// 	"go.mongodb.org/mongo-driver/bson/primitive"
// 	"go.mongodb.org/mongo-driver/mongo"
// 	"go.mongodb.org/mongo-driver/mongo/options"
// )

// type ProductService struct {
// 	productPb.UnimplementedProductServiceServer
// 	collection *mongo.Collection
// }

// func NewProductService(client *mongo.Client) *ProductService {
// 	return &ProductService{collection: client.Database("testing-h8").Collection("products")}
// }

// func (s *ProductService) mustEmbedUnimplementedProductServiceServer() {}
// func (s *ProductService) GetAllProducts(ctx context.Context, req *productPb.GetAllProductsRequest) (*productPb.GetAllProductsResponse, error) {
// 	log.Println("Get All Products")
// 	// Set options
// 	findOptions := options.Find()

// 	// Find all documents
// 	cursor, err := s.collection.Find(ctx, bson.M{}, findOptions)

// 	if err != nil {
// 		return nil, err
// 	}

// 	defer cursor.Close(ctx)

// 	// Iterate over documents
// 	var products []*productPb.Product

// 	for cursor.Next(ctx) {
// 		var product productPb.Product
// 		if err = cursor.Decode(&product); err != nil {
// 			return nil, err
// 		}
// 		products = append(products, &product)
// 	}

// 	if err = cursor.Err(); err != nil {
// 		return nil, err
// 	}

// 	return &productPb.GetAllProductsResponse{Products: products}, nil
// }

// func (s *ProductService) GetProductByID(ctx context.Context, req *productPb.GetProductByIdRequest) (*productPb.GetProductByIdResponse, error) {

// 	// Mengambil produk berdasarkan ID
// 	var result bson.M
// 	id, _ := primitive.ObjectIDFromHex(req.Id)
// 	err := s.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&result)
// 	if err != nil {
// 		if err == mongo.ErrNoDocuments {
// 			return nil, fmt.Errorf("Produk dengan ID %s tidak ditemukan", req.Id)
// 		}
// 		log.Fatal(err)
// 	}

// 	// Membuat respons dari produk yang ditemukan
// 	res := &productPb.GetProductByIdResponse{
// 		Id:          result["_id"].(primitive.ObjectID).Hex(),
// 		Name:        result["name"].(string),
// 		Description: result["description"].(string),
// 		Price:       result["price"].(float32),
// 		Stock:       result["stock"].(int32),
// 	}

// 	return res, nil
// }

// func (s *ProductService) AddProduct(ctx context.Context, req *productPb.AddProductRequest) (*productPb.AddProductResponse, error) {
// 	fmt.Println("Tambah produk")
// 	product := bson.M{
// 		"name":        req.Name,
// 		"description": req.Description,
// 		"price":       req.Price,
// 		"stock":       req.Stock,
// 	}
// 	result, err := s.collection.InsertOne(ctx, product)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	// Membuat respons dari produk yang ditambahkan
// 	res := &productPb.AddProductResponse{
// 		Id: result.InsertedID.(primitive.ObjectID).Hex(),
// 	}

// 	return res, nil
// }

// func (s *ProductService) UpdateProduct(ctx context.Context, req *productPb.UpdateProductRequest) (*productPb.UpdateProductResponse, error) {
// 	// Mengubah produk berdasarkan ID
// 	id, _ := primitive.ObjectIDFromHex(req.Id)
// 	filter := bson.M{"_id": id}
// 	update := bson.M{
// 		"$set": bson.M{
// 			"name":        req.Name,
// 			"description": req.Description,
// 			"price":       req.Price,
// 			"stock":       req.Stock,
// 		},
// 	}
// 	_, err := s.collection.UpdateOne(ctx, filter, update)
// 	log.Fatal(err)

// 	// Membuat respons dari produk yang diubah
// 	res := &productPb.UpdateProductResponse{
// 		Message: fmt.Sprintf("Produk dengan ID %s berhasil diubah", req.Id),
// 	}

// 	return res, nil
// }

// func (s *ProductService) DeleteProduct(ctx context.Context, req *productPb.DeleteProductRequest) (*productPb.DeleteProductResponse, error) {
// 	// Menghapus produk berdasarkan ID
// 	id, _ := primitive.ObjectIDFromHex(req.Id)
// 	filter := bson.M{"_id": id}
// 	_, err := s.collection.DeleteOne(ctx, filter)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	// Membuat respons dari produk yang dihapus
// 	res := &productPb.DeleteProductResponse{
// 		Message: fmt.Sprintf("Produk dengan ID %s berhasil dihapus", req.Id),
// 	}

// 	return res, nil
// }
