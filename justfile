clean:
	rm -rf public

build:
	hugo --minify
	html2pdf --background --paper A4 public/index.html -o public/pdf/Nilesh_Kevlani.pdf

server:
	hugo server --minify

lint:
	typos
	yamllint .
	stylelint "**/*.css"

lint-ci:
	stylelint "**/*.css"
