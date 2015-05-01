API_SERVER      = api
API_TEST_DIR    = test
WEB_SERVER      = web
DOCS_DIR        = api/docs

.PHONY: api-server
api-server:
	cd $(API_SERVER); go run main.go

.PHONY: test-api
api-test:
	jasmine-node api/test --verbose

.PHONY: serve
serve:
	cd $(WEB_SERVER); python -m SimpleHTTPServer

.PHONY: docs
docs:
	aglio -i $(DOCS_DIR)/api.md -o $(DOCS_DIR)/index.html
	@echo '----------------------'
	@echo 'API Docs updated.'
	@echo 'Available at /api/docs'
	@echo '----------------------'

.PHONY: install-go-deps
install-go-deps:
	@echo '--------------------------'
	@echo 'Getting Go dependencies...'
	@echo '--------------------------'
	cd $(API_SERVER); go get all; cd ..
	@echo '-----------------'
	@echo 'Packages installed.'
	@echo '-----------------'

.PHONY: update-go-deps
update-go-deps:
	@echo '---------------------------'
	@echo 'Updating Go dependencies...'
	@echo '---------------------------'
	cd $(API_SERVER); go get -u all; cd ..
	@echo '-----------------'
	@echo 'Packages updated.'
	@echo '-----------------'

# http://stackoverflow.com/questions/8889035/how-to-document-a-makefile
.PHONY: help
help:
	@echo  'Makefile to assist with initializing Ricetta.'
	@echo  ''
	@echo  'Usage:'
	@echo  ''
	@echo  '       make [command]'
	@echo  ''
	@echo  'The commands are:'
	@echo  ''
	@echo  '  api-server         - Load typical API server session using local database'
	@echo  '  docs               - Build Aglio docs for the API'
	@echo  '  install-go-deps    - Install Go package dependencies'
	@echo  '  update-go-deps     - Update Go package dependencies'
	@echo  ''
