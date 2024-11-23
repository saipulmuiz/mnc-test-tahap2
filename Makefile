generate-mocks:
	# mockery
	@mockery --dir ./repositories --case underscore --output ./repositories/mocks --name UserRepo
	@mockery --dir ./repositories --case underscore --output ./repositories/mocks --name ProductRepo
	@mockery --dir ./repositories --case underscore --output ./repositories/mocks --name CartRepo
	@mockery --dir ./repositories --case underscore --output ./repositories/mocks --name OrderRepo

	# repositories
	@mockgen -destination=./repositories/mocks/mock_user_repository.go -package=mocks github.com/saipulmuiz/mnc-test-tahap2/repositories UserRepo
	@mockgen -destination=./repositories/mocks/mock_product_repository.go -package=mocks github.com/saipulmuiz/mnc-test-tahap2/repositories ProductRepo
	@mockgen -destination=./repositories/mocks/mock_cart_repository.go -package=mocks github.com/saipulmuiz/mnc-test-tahap2/repositories CartRepo
	@mockgen -destination=./repositories/mocks/mock_order_repository.go -package=mocks github.com/saipulmuiz/mnc-test-tahap2/repositories OrderRepo
