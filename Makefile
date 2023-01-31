

code-gen:
	abigen --abi=bindings/abacus/abacus.json --pkg=abacus --out=bindings/abacus/abacus.go
	abigen --abi=bindings/clip/Clipper.json --pkg=clipper --out=bindings/clip/Clipper.go
	abigen --abi=bindings/uniswap_v3_quoter/uniswap_v3_quoter.json --pkg=uniswap_v3_quoter --out=bindings/uniswap_v3_quoter/uniswap_v3_quoter.go
	abigen --abi=bindings/vat/vat.json --pkg=vat --out=bindings/vat/vat.go

	./eh-gen --abi=bindings/clip/Clipper.json  --output-dir=bindings/clip --contract=Clipper
	./eh-gen --abi=bindings/uniswap_v3_quoter/uniswap_v3_quoter.json  --output-dir=bindings/uniswap_v3_quoter --contract=UniswapV3Quoter
	./eh-gen --abi=bindings/vat/vat.json  --output-dir=bindings/vat --contract=vat

clean:
	rm bindings/abacus/abacus.go
	rm bindings/clip/Clipper.go
