ABIGEN := abigen
EHGEN := ehgen

API_BINDINGS := bindings
ZARBAN := $(API_BINDINGS)/zarban
CHAINLINK := $(API_BINDINGS)/chainlink
IERC20 := $(API_BINDINGS)/ierc20
UNISWAPV3 := $(API_BINDINGS)/uniswapv3

devtools:
	cd x && go build codegen/events_handler.go

abi-gen-ierc20:
	$(ABIGEN) --abi=$(IERC20)/ierc20.json --pkg=ierc20 --out=$(IERC20)/ierc20.go

ABI_GEN_TARGETS := zarjoin median osm jug vat spot gemjoin cdpmanager proxy proxy_registry deployment lendingpool lendingpool_address_provider lendingpool_address_provider_registry ui_pool_data_provider wallet_balance_provider dog ilkregistry clipper abacus flopper

$(ABI_GEN_TARGETS):
	$(ABIGEN) --abi=$(ZARBAN)/$@/$@.json --pkg=$@ --out=$(ZARBAN)/$@/$@.go

abi-gen-vow:
	$(ABIGEN) --abi=$(ZARBAN)/vow/vow.json --pkg=vow --out=$(ZARBAN)/vow/vow.go --alias Sin=TotalSin

CHAINLINK_TARGETS := access_controlled_offchain_aggregator

$(CHAINLINK_TARGETS):
	$(ABIGEN) --abi=$(CHAINLINK)/$@/$@.json --pkg=$@ --out=$(CHAINLINK)/$@/$@.go
	$(EHGEN) --abi=$(CHAINLINK)/$(@:eh_gen_%=%)/$(@:eh_gen_%=%).json --output-dir=$(CHAINLINK)/$(@:eh_gen_%=%) --contract=$$(echo $(@:eh_gen_%=%) | perl -pe "s/_(.)/\u\1/g; s/^(.)/\u\1/g")


UNISWAPV3_TARGETS := pool quoter

$(UNISWAPV3_TARGETS):
	$(ABIGEN) --abi=$(UNISWAPV3)/$@/$@.json --pkg=$@ --out=$(UNISWAPV3)/$@/$@.go

abi-gen: abi-gen-ierc20 $(ABI_GEN_TARGETS) $(CHAINLINK_TARGETS) abi-gen-vow $(UNISWAPV3_TARGETS)

code-gen: devtools abi-gen eh-gen

EH_GEN_TARGETS := median osm vat vow zarjoin jug spot cdpmanager proxy_registry proxy lendingpool gemjoin dog ilkregistry clipper flopper

$(EH_GEN_TARGETS:%=eh_gen_%):
	$(EHGEN) --abi=$(ZARBAN)/$(@:eh_gen_%=%)/$(@:eh_gen_%=%).json --output-dir=$(ZARBAN)/$(@:eh_gen_%=%) --contract=$$(echo $(@:eh_gen_%=%) | perl -pe "s/_(.)/\u\1/g; s/^(.)/\u\1/g")

EH_GEN_UNISWAPV3_TARGETS := pool

$(EH_GEN_UNISWAPV3_TARGETS:%=eh_gen_%):
	$(EHGEN) --abi=$(UNISWAPV3)/$(@:eh_gen_%=%)/$(@:eh_gen_%=%).json --output-dir=$(UNISWAPV3)/$(@:eh_gen_%=%) --contract=$$(echo $(@:eh_gen_%=%) | perl -pe "s/_(.)/\u\1/g; s/^(.)/\u\1/g")

eh-gen: $(EH_GEN_TARGETS:%=eh_gen_%) $(EH_GEN_UNISWAPV3_TARGETS:%=eh_gen_%)

clean:
	find $(API_BINDINGS) -type f -name '*.go' -exec rm {} +

.PHONY: abi_gen eh_gen clean
