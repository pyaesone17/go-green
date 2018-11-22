all: app

app:
	cd gogreensvc; make; cd ..

dist:
	cd gogreensvc; make dist; cd ..

clean:
	cd gogreensvc; make clean

test:
	ginkgo -v -r .

.PHONY: app clean test dist
