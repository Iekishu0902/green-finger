MSV_CODE := bff
MK_DIR := $(dir $(lastword $(MAKEFILE_LIST)))/scripts/makefiles
ENV_DIR := $(dir $(lastword $(MAKEFILE_LIST)))/env

include $(MK_DIR)/protoc.mk
include $(ENV_DIR)/env.local.conf
BASE_ENV := \
	LOGLEVEL=$(LOGLEVEL) \
	SERVER_PORT=$(SERVER_PORT)

include $(MK_DIR)/run-local.mk
