clean:
	rm -rf build

build:
	go run main.go
	mkdir build/pdf -p
	mv index.html build/
	cp -r static/* build/
	html2pdf --background --paper A4 build/index.html -o Nilesh_Kevlani.pdf
	mv Nilesh_Kevlani.pdf build/pdf/

server: build
	python -m http.server --directory build

check-typo:
	typos

check-lint-go:
	golangci-lint run
	stylelint --configBasedir=/usr/lib/node_modules "**/*.css"

check-lint-css:
	stylelint --configBasedir=/usr/lib/node_modules "**/*.css"

check-lint: check-lint-go check-lint-css

check: check-typo check-lint
