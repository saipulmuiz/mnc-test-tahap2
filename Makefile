generate-mocks:
	# repositories
	@mockgen -destination=./repositories/mocks/mock_user_repository.go -package=mocks github.com/saipulmuiz/mnc-test-tahap2/repositories UserRepo
	@mockgen -destination=./repositories/mocks/mock_topup_repository.go -package=mocks github.com/saipulmuiz/mnc-test-tahap2/repositories TopupRepo
	@mockgen -destination=./repositories/mocks/mock_payment_repository.go -package=mocks github.com/saipulmuiz/mnc-test-tahap2/repositories PaymentRepo
	@mockgen -destination=./repositories/mocks/mock_transfer_repository.go -package=mocks github.com/saipulmuiz/mnc-test-tahap2/repositories TransferRepo
	@mockgen -destination=./repositories/mocks/mock_transaction_repository.go -package=mocks github.com/saipulmuiz/mnc-test-tahap2/repositories TransactionRepo
	@mockgen -destination=./repositories/mocks/mock_tx_manager_repository.go -package=mocks github.com/saipulmuiz/mnc-test-tahap2/repositories TxManagerRepo
