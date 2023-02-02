-include $(ENV_DIR)/env.local.conf

LOCAL_ENV := \
	$(BASE_ENV) \

run-local: protoc run-mysql-local
	env $(LOCAL_ENV) go run cmd/plants_gateway/main.go