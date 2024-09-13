.ONESHELL:
.PHONY: all step1 step2 step3

tailwind-watch:
	npx tailwindcss -i ./static/css/input.css -o ./static/css/output.css --watch 

templ-watch:
	templ generate --watch 

air:
	air 

all:
	${MAKE} -j3 tailwind-watch templ-watch air 

