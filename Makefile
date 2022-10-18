

code-gen:
	abigen --abi=abis/ierc20/ierc20.json --pkg=IERC20 --out=abis/ierc20/ierc20.go
	abigen --abi=abis/abacus/abacus.json --pkg=abacus --out=abis/abacus/abacus.go

clean:
	rm abis/ierc20/ierc20.go
