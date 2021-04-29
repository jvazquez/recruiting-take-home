DOCKER = $(shell which docker)
BUILD_ARG = $(if $(filter  $(NOCACHE), 1),--no-cache)
BASE_PATH?="$(shell pwd)"
ENV_FILE?=.env
TEST_CONTAINER_NAME?=strings-test

.PHONY: all test_function_image run_function_tests

test_function_image:
	$(DOCKER) build $(BUILD_ARG) \
	-f build/package/test/Dockerfile \
	-t $(TEST_CONTAINER_NAME) .
run_function_tests:
	$(DOCKER) run \
    	--rm \
    	--env-file .env.test \
    	$(TEST_CONTAINER_NAME)
