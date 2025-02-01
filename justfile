build-html:
	go run main.go

build-pdf:
	html2pdf index.html -o Nilesh_Kevlani.pdf

build-dir: build-html build-pdf
	mkdir build -p
	mv index.html build/
	mv Nilesh_Kevlani.pdf build/
	cp favicon.ico build/
	cp -r assets build/

