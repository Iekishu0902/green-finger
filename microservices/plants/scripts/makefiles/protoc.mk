PB_PATH := infrastructure/rpc
BUF_PATH := tools/protogen
PROTO_PATH := api
ROOT_DIR := $(abspath .)

protoc: 
	rm -rf $(PB_PATH)/*.pb*.go
	mkdir -p $(BUF_PATH)/buf
	cp $(BUF_PATH)/buf.gen.yaml $(BUF_PATH)/buf.yaml $(BUF_PATH)/buf.lock $(BUF_PATH)/buf
	cp $(PROTO_PATH)/*.proto $(BUF_PATH)/buf
	cd ./$(BUF_PATH)/buf && \
	buf generate -v --debug --timeout 10m && \
	mv ./*.pb*.go $(ROOT_DIR)/$(PB_PATH)
	rm -rf $(BUF_PATH)/buf
