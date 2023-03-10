MSV_CODE := plants
MK_DIR := $(dir $(lastword $(MAKEFILE_LIST)))/scripts/makefiles
ENV_DIR := $(dir $(lastword $(MAKEFILE_LIST)))/env

include $(MK_DIR)/protoc.mk
include $(MK_DIR)/mysql.mk
include $(ENV_DIR)/env.local.conf
BASE_ENV := \
	LOGLEVEL=$(LOGLEVEL) \
	SERVER_PORT=$(SERVER_PORT) \
	MYSQL_HOST=$(MYSQL_HOST) \
	MYSQL_USENAME=$(MYSQL_USENAME) \
	MYSQL_PASSWORD=$(MYSQL_PASSWORD) \
	MYSQL_DATABASE=$(MYSQL_DATABASE) \
	MYSQL_PORT=$(MYSQL_PORT)

include $(MK_DIR)/run-local.mk
