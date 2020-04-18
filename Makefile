BUILD_DIR = build

TARGET = 1.1.0-oas3
SPEC_URL = https://api.swaggerhub.com/apis/OneLogin-Auth/onelogin-api/$(TARGET)
SPEC_FILE = build/$(TARGET).json
SERVERS_JSON = src/servers.json

TEMPLATE = src/templates
GENERATOR = src/go-client.groovy
GENERATE_TMP_DIR = tmp
FILE_LIST = file_list

PACKAGE_NAME = github.com/yacchi/go-onelogin-oas
PACKAGE_VERSION = 0.1.3

.PHONY: all
all: $(SPEC_FILE) generate

$(BUILD_DIR):
	mkdir -p $(BUILD_DIR)

.PHONY: spec
spec: $(SPEC_FILE)

$(SPEC_FILE): $(BUILD_DIR) $(SERVERS_JSON)
	curl $(SPEC_URL) | jq -s add - src/servers.json > $(SPEC_FILE).tmp
	mv $(SPEC_FILE).tmp $(SPEC_FILE)

.PHONY: generate
generate: $(BUILD_DIR) $(SPEC_FILE)
	GO_PACKAGE_VERSION=$(PACKAGE_VERSION) groovy $(GENERATOR) generate \
	-o $(GENERATE_TMP_DIR) \
	--input-spec $(SPEC_FILE) \
	--generator-name MyGoClientCodegen \
	--template-dir $(TEMPLATE) \
	--package-name $(PACKAGE_NAME) \
	--model-package $(PACKAGE_NAME)/models \
	--api-package "onelogin $(PACKAGE_NAME)"

	rm $(GENERATE_TMP_DIR)/{.gitignore,.travis.yml,.openapi-generator-ignore,git_push.sh,go.mod,go.sum}
	rm -r $(GENERATE_TMP_DIR)/api

	find $(GENERATE_TMP_DIR) -name '*.go' | xargs goimports -w

	cat $(FILE_LIST) | xargs rm -v ; true
	find $(GENERATE_TMP_DIR) -type f | sed -e "s#^$(GENERATE_TMP_DIR)/*##" | sort > $(FILE_LIST)

	mv $(GENERATE_TMP_DIR)/* $(GENERATE_TMP_DIR)/.[!.]* .
	rmdir $(GENERATE_TMP_DIR)