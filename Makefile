

code-gen:
	abigen --abi=bindings/abacus/abacus.json --pkg=abacus --out=bindings/abacus/abacus.go
	abigen --abi=bindings/clip/clipper.json --pkg=clipper --out=bindings/clip/clipper.go
	abigen --abi=bindings/vat/vat.json --pkg=vat --out=bindings/vat/vat.go
	abigen --abi=bindings/vow/vow.json --pkg=vow --out=bindings/vow/vow.go --alias Sin=TotalSin
	abigen --abi=bindings/dog/dog.json --pkg=dog --out=bindings/dog/dog.go
	abigen --abi=bindings/flopper/flopper.json --pkg=flopper --out=bindings/flopper/flopper.go
	abigen --abi=bindings/uniswap_v3_quoter/uniswap_v3_quoter.json --pkg=uniswap_v3_quoter --out=bindings/uniswap_v3_quoter/uniswap_v3_quoter.go

	./eh-gen --abi=bindings/clip/clipper.json  --output-dir=bindings/clip --contract=Clipper
	./eh-gen --abi=bindings/vat/vat.json  --output-dir=bindings/vat --contract=Vat
	./eh-gen --abi=bindings/vow/vow.json  --output-dir=bindings/vow --contract=Vow
	./eh-gen --abi=bindings/uniswap_v3_quoter/uniswap_v3_quoter.json  --output-dir=bindings/uniswap_v3_quoter --contract=UniswapV3Quoter

clean:
	rm bindings/abacus/abacus.go
	rm bindings/clip/clipper.go
	rm bindings/vat/vat.go
	rm bindings/vow/vow.go
	rm bindings/dog/dog.go
	rm bindings/flopper/flopper.go
