

code-gen:
	abigen --abi=bindings/abacus/abacus.json --pkg=abacus --out=bindings/abacus/abacus.go
	abigen --abi=bindings/clip/Clipper.json --pkg=clipper --out=bindings/clip/Clipper.go

	./eh-gen --abi=bindings/clip/Clipper.json  --output-dir=bindings/clip --contract=Clipper

clean:
	rm bindings/abacus/abacus.go
	rm bindings/clip/Clipper.go
