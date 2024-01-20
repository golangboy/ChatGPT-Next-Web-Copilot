.PHONY: help
help:
	@echo "Please use \`make <target>\` where <target> is one of"
	@echo "  dev       to start development server"
	@echo "  get-copilot-token       to get Github Copilot Plugin Token"

.PHONY: dev
dev:
	@echo "Starting development server..."
	@go run main.go

.PHONY: get-copilot-token
# The script below will automatically install Github Copilot CLI and obtain the Github Copilot Plugin Token through authorization
get-copilot-token:
	@echo "Getting Github Copilot Plugin Token..."
	@bash ./shells/get_copilot_token.sh
