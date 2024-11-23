package services

import (
	"errors"
	"net/http"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"

	"github.com/saipulmuiz/mnc-test-tahap2/models"
	"github.com/saipulmuiz/mnc-test-tahap2/params"
	"github.com/saipulmuiz/mnc-test-tahap2/repositories/mocks"
)

func Test_ProductService_GetProducts(t *testing.T) {
	type testCase struct {
		name             string
		wantError        bool
		page, size       int
		onGetProducts    func(mock *mocks.MockProductRepo)
		expectedResponse *params.Response
	}

	var testTable []testCase
	testTable = append(testTable, testCase{
		name:      "success",
		wantError: false,
		page:      1,
		size:      10,
		onGetProducts: func(mock *mocks.MockProductRepo) {
			mock.EXPECT().GetProducts(1, 10).Return(&[]models.Product{
				{ProductID: 1, ProductName: "Product 1", ProductPrice: 100, Stock: 50, MinStock: 10},
			}, int64(1), nil)
		},
		expectedResponse: &params.Response{
			Status: http.StatusOK,
			Payload: params.ResponseWithPagination{
				Message: "Success to get product",
				Data: []params.GetProductRes{
					{
						ProductID:    1,
						ProductName:  "Product 1",
						Description:  "",
						ProductPrice: 100,
						Stock:        50,
						MinStock:     10,
						CreatedAt:    "0001-01-01 00:00:00",
						UpdatedAt:    "0001-01-01 00:00:00",
					},
				},
				Pagination: params.PaginationResponse{
					CurrentPage:  1,
					PageSize:     10,
					TotalCount:   1,
					TotalPages:   1,
					FirstPage:    1,
					NextPage:     1,
					LastPage:     1,
					CurrentCount: 1,
				},
			},
		},
	})

	for _, tc := range testTable {
		t.Run(tc.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			productRepo := mocks.NewMockProductRepo(mockCtrl)

			if tc.onGetProducts != nil {
				tc.onGetProducts(productRepo)
			}

			service := &ProductService{productRepo: productRepo}

			resp := service.GetProducts(tc.page, tc.size)

			if tc.wantError {
				assert.NotNil(t, resp)
				assert.Equal(t, http.StatusBadRequest, resp.Status)
			} else {
				assert.Equal(t, tc.expectedResponse, resp)
			}
		})
	}
}

func Test_ProductService_GetProductById(t *testing.T) {
	type testCase struct {
		name              string
		wantError         bool
		productId         int
		onFindProductById func(mock *mocks.MockProductRepo)
		expectedResponse  *params.Response
	}

	var testTable []testCase
	testTable = append(testTable, testCase{
		name:      "success",
		wantError: false,
		productId: 1,
		onFindProductById: func(mock *mocks.MockProductRepo) {
			mock.EXPECT().FindById(1).Return(&models.Product{
				ProductID: 1, ProductName: "Product 1", ProductPrice: 100, Stock: 50, MinStock: 10,
			}, nil)
		},
		expectedResponse: &params.Response{
			Status: http.StatusOK,
			Payload: params.ResponseSuccess{
				Message: "Success Get Product By Id",
				Data: params.GetProductRes{
					ProductID:    1,
					ProductName:  "Product 1",
					Description:  "",
					ProductPrice: 100,
					Stock:        50,
					MinStock:     10,
					CreatedAt:    "0001-01-01 00:00:00",
					UpdatedAt:    "0001-01-01 00:00:00",
				},
			},
		},
	})
	testTable = append(testTable, testCase{
		name:      "product not found",
		wantError: true,
		productId: 1,
		expectedResponse: &params.Response{
			Status: http.StatusNotFound,
			Payload: params.ResponseErrorMessage{
				Error: "Product not found",
			},
		},
		onFindProductById: func(mock *mocks.MockProductRepo) {
			mock.EXPECT().FindById(1).Return(nil, gorm.ErrRecordNotFound)
		},
	})

	for _, tc := range testTable {
		t.Run(tc.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			productRepo := mocks.NewMockProductRepo(mockCtrl)

			if tc.onFindProductById != nil {
				tc.onFindProductById(productRepo)
			}

			service := &ProductService{productRepo: productRepo}

			resp := service.GetProductById(tc.productId)

			if tc.wantError {
				assert.Equal(t, http.StatusNotFound, resp.Status)
				assert.Equal(t, tc.expectedResponse, resp)
			} else {
				assert.Equal(t, tc.expectedResponse, resp)
			}
		})
	}
}

func Test_ProductService_CreateProduct(t *testing.T) {
	type testCase struct {
		name             string
		wantError        bool
		expectedResponse *params.Response
		request          params.CreateProductReq
		onCreateProduct  func(mock *mocks.MockProductRepo)
	}

	var testTable []testCase
	testTable = append(testTable, testCase{
		name:      "success",
		wantError: false,
		request: params.CreateProductReq{
			ProductName: "Product 1", Description: "Description", ProductPrice: 100, Stock: 50, MinStock: 10,
		},
		onCreateProduct: func(mock *mocks.MockProductRepo) {
			mock.EXPECT().CreateProduct(gomock.Any()).Return(&models.Product{
				ProductID: 1, ProductName: "Product 1", ProductPrice: 100, Stock: 50, MinStock: 10,
			}, nil)
		},
		expectedResponse: &params.Response{
			Status: http.StatusCreated,
			Payload: params.ResponseSuccess{
				Message: "Create product successfully",
				Data: &models.Product{
					ProductID:    1,
					ProductName:  "Product 1",
					Description:  "",
					ProductPrice: 100,
					Stock:        50,
					MinStock:     10,
				},
			},
		},
	})

	testTable = append(testTable, testCase{
		name:      "error creating product",
		wantError: true,
		expectedResponse: &params.Response{
			Status: http.StatusBadRequest,
			Payload: params.ResponseErrorMessage{
				Error: "some error",
			},
		},
		request: params.CreateProductReq{
			ProductName: "Product 1", Description: "Description", ProductPrice: 100, Stock: 50, MinStock: 10,
		},
		onCreateProduct: func(mock *mocks.MockProductRepo) {
			mock.EXPECT().CreateProduct(gomock.Any()).Return(nil, errors.New("some error"))
		},
	})

	for _, tc := range testTable {
		t.Run(tc.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			productRepo := mocks.NewMockProductRepo(mockCtrl)

			if tc.onCreateProduct != nil {
				tc.onCreateProduct(productRepo)
			}

			service := &ProductService{productRepo: productRepo}

			resp := service.CreateProduct(tc.request)

			if tc.wantError {
				assert.NotNil(t, resp)
				assert.Equal(t, tc.expectedResponse.Status, resp.Status)
				assert.Equal(t, tc.expectedResponse.Payload, resp.Payload)
			} else {
				assert.Equal(t, http.StatusCreated, resp.Status)
				assert.Equal(t, tc.expectedResponse, resp)
			}
		})
	}
}

func Test_ProductService_UpdateProduct(t *testing.T) {
	type testCase struct {
		name               string
		wantError          bool
		expectedResponse   *params.Response
		productId          int
		request            params.UpdatedProductReq
		onCheckProductByID func(mock *mocks.MockProductRepo)
		onUpdateProduct    func(mock *mocks.MockProductRepo)
	}

	var testTable []testCase
	testTable = append(testTable, testCase{
		name:      "success",
		wantError: false,
		productId: 1,
		request: params.UpdatedProductReq{
			ProductName: "Updated Product", Description: "Updated Description", ProductPrice: 200, Stock: 60, MinStock: 20,
		},
		onCheckProductByID: func(mock *mocks.MockProductRepo) {
			mock.EXPECT().CheckProductByID(1, gomock.Any()).Return(&models.Product{ProductID: 1}, nil)
		},
		onUpdateProduct: func(mock *mocks.MockProductRepo) {
			mock.EXPECT().UpdateProduct(gomock.Any(), 1, gomock.Any()).Return(&models.Product{
				ProductID: 1, ProductName: "Updated Product", Description: "Updated Description", ProductPrice: 200, Stock: 60, MinStock: 20,
			}, nil)
		},
		expectedResponse: &params.Response{
			Status: http.StatusOK,
			Payload: params.ResponseSuccess{
				Message: "Product successfully updated",
				Data: &models.Product{
					ProductID:    1,
					ProductName:  "Updated Product",
					Description:  "Updated Description",
					ProductPrice: 200,
					Stock:        60,
					MinStock:     20,
				},
			},
		},
	})

	testTable = append(testTable, testCase{
		name:      "product not found",
		wantError: true,
		productId: 1,
		expectedResponse: &params.Response{
			Status: http.StatusNotFound,
			Payload: params.ResponseErrorMessage{
				Error: "Product not found",
			},
		},
		onCheckProductByID: func(mock *mocks.MockProductRepo) {
			mock.EXPECT().CheckProductByID(1, gomock.Any()).Return(nil, nil)
		},
	})

	for _, tc := range testTable {
		t.Run(tc.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			productRepo := mocks.NewMockProductRepo(mockCtrl)

			if tc.onCheckProductByID != nil {
				tc.onCheckProductByID(productRepo)
			}
			if tc.onUpdateProduct != nil {
				tc.onUpdateProduct(productRepo)
			}

			service := &ProductService{productRepo: productRepo}

			resp := service.UpdateProduct(tc.productId, tc.request)

			if tc.wantError {
				assert.NotNil(t, resp)
				assert.Equal(t, tc.expectedResponse.Status, resp.Status)
				assert.Equal(t, tc.expectedResponse, resp)
			} else {
				assert.NotNil(t, resp)
				assert.Equal(t, http.StatusOK, resp.Status)
				assert.Equal(t, tc.expectedResponse, resp)
			}
		})
	}
}

func Test_ProductService_DeleteProduct(t *testing.T) {
	type testCase struct {
		name               string
		wantError          bool
		expectedResponse   *params.Response
		productId          int
		onCheckProductByID func(mock *mocks.MockProductRepo)
		onDeleteProduct    func(mock *mocks.MockProductRepo)
	}

	var testTable []testCase
	testTable = append(testTable, testCase{
		name:      "success",
		wantError: false,
		productId: 1,
		onCheckProductByID: func(mock *mocks.MockProductRepo) {
			mock.EXPECT().CheckProductByID(1, gomock.Any()).Return(&models.Product{ProductID: 1}, nil)
		},
		onDeleteProduct: func(mock *mocks.MockProductRepo) {
			mock.EXPECT().DeleteProduct(1).Return(nil)
		},
		expectedResponse: &params.Response{
			Status: http.StatusOK,
			Payload: params.ResponseSuccessMessage{
				Message: "Success delete product",
			},
		},
	})

	testTable = append(testTable, testCase{
		name:      "product not found",
		wantError: true,
		productId: 1,
		expectedResponse: &params.Response{
			Status: http.StatusNotFound,
			Payload: params.ResponseErrorMessage{
				Error: "Product not found",
			},
		},
		onCheckProductByID: func(mock *mocks.MockProductRepo) {
			mock.EXPECT().CheckProductByID(1, gomock.Any()).Return(models.Product{}, nil)
		},
	})

	for _, tc := range testTable {
		t.Run(tc.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			productRepo := mocks.NewMockProductRepo(mockCtrl)

			if tc.onCheckProductByID != nil {
				tc.onCheckProductByID(productRepo)
			}
			if tc.onDeleteProduct != nil {
				tc.onDeleteProduct(productRepo)
			}

			service := &ProductService{productRepo: productRepo}

			resp := service.DeleteProduct(tc.productId)

			if tc.wantError {
				assert.NotNil(t, resp)
				assert.Equal(t, tc.expectedResponse.Status, resp.Status)
				assert.Equal(t, tc.expectedResponse, resp)
			} else {
				assert.NotNil(t, resp)
				assert.Equal(t, http.StatusOK, resp.Status)
				assert.Equal(t, tc.expectedResponse, resp)
			}
		})
	}
}
