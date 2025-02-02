clean:
	rm -rf build

build-html:
	go run main.go

build-pdf:
	html2pdf --background --paper A4 index.html -o Nilesh_Kevlani.pdf

build-dir: build-html build-pdf
	mkdir build/pdf -p
	mv index.html build/
	mv Nilesh_Kevlani.pdf build/pdf/
	cp favicon.ico build/
	cp -r assets build/
	cp -r static/* build/

server: build-dir
	python -m http.server --directory build
