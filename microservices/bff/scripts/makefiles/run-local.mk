-include $(ENV_DIR)/env.local.conf

LOCAL_ENV := \
	$(BASE_ENV) \

run-local: protoc
	env $(LOCAL_ENV) go run cmd/bff_gateway/main.go