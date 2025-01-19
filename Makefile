.ONESHELL:
.PHONY: all step1 step2 step3

tailwind-watch:
	npx tailwindcss -i ./static/css/input.css -o ./static/css/output.css --watch 
templ-watch:
	templ generate --watch 
air:
	air 
watch:
	${MAKE} -j3 tailwind-watch templ-watch air 

upclient-generate:
	openapi-generator-cli generate   -i openapi.json   -g go   -o ./datafetcher/upclient   --additional-properties packageName=upclient   --git-user-id esteanes   --git-repo-id expense-manager/datafetcher/upclient

build:
	templ generate && gofmt -s -w . && go mod tidy && go build -o expense-manager

docker:
	docker build -t expense-manager:latest -f Dockerfile .
