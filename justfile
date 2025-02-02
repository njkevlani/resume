clean:
	rm -rf build

build-html:
	go run main.go

build-pdf:
	html2pdf --background --paper A4 build/index.html -o Nilesh_Kevlani.pdf
	mv Nilesh_Kevlani.pdf build/pdf/

build-dir: build-html && build-pdf
	mkdir build/pdf -p
	mv index.html build/
	cp -r static/* build/

server: build-dir
	python -m http.server --directory build
