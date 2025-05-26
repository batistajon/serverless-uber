# Env vars for uber_receipt_ready lambda

UBER_RECEIPT_READY_PATH=./cmd/lambdas/uber_receipt_ready
UBER_RECEIPT_READY_BINARY_NAME=${UBER_RECEIPT_READY_PATH}/bin/bootstrap
UBER_RECEIPT_READY_ZIP_NAME=${UBER_RECEIPT_READY_PATH}/uber_receipt_ready.zip
UBER_RECEIPT_READY_ENTRYPOINT=${UBER_RECEIPT_READY_PATH}/main.go

# Build the Go binary for Linux (e.g. AWS Lambda)
build_uber_receipt_ready:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -tags lambda.norpc -o $(UBER_RECEIPT_READY_BINARY_NAME) $(UBER_RECEIPT_READY_ENTRYPOINT)

zip_uber_receipt_ready: build_uber_receipt_ready
	zip -j $(UBER_RECEIPT_READY_ZIP_NAME) $(UBER_RECEIPT_READY_BINARY_NAME)

clean_uber_receipt_ready:
	rm -f $(UBER_RECEIPT_READY_BINARY_NAME) $(UBER_RECEIPT_READY_ZIP_NAME)
