package events

import (
	"html/template"
	"os"
)

type GenData struct {
	Package               string
	BindingPath           string
	BindingEventName      string
	BindingEventSignature string
	BindingContract       string
}

func CodeGen(d GenData, output string) error {
	t := template.Must(template.New("event_handler").Parse(eventHandlerTemplate))
	f, err := os.Create(output)
	if err != nil {
		return err
	}
	return t.Execute(f, d)
}

var eventHandlerTemplate = `// Code generated - DO NOT EDIT.
// This file is a generated event handler and any manual changes will be lost.

package {{.Package}}

import (
	"errors"

	"github.com/zarbanio/auction-keeper/x/events"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/zarbanio/auction-keeper/x/eth"
)

type {{.BindingEventName}}Handler struct {
	addr string
	binding  *{{.BindingContract}}
	callback events.CallbackFn[{{.BindingContract}}{{.BindingEventName}}]
}

func (h *{{.BindingEventName}}Handler) ID() string {
	return h.addr + ":" + "{{.BindingEventSignature}}"
}

func (h *{{.BindingEventName}}Handler) DecodeLog(log eth.Log) (interface{}, error) {
	return h.binding.Parse{{.BindingEventName}}(log.Log)
}

func (h *{{.BindingEventName}}Handler) HandleEvent(log eth.Log, event interface{}) error {
	e, ok := event.({{.BindingContract}}{{.BindingEventName}})
	if !ok {
		return errors.New("event type is not {{.BindingContract}}{{.BindingEventName}}")
	}
	return h.callback(log, e)
}

func (h *{{.BindingEventName}}Handler) DecodeAndHandle(log eth.Log) error {
	e, err := h.binding.Parse{{.BindingEventName}}(log.Log)
	if err != nil {
		return err
	}
	return h.callback(log, *e)
}

func New{{.BindingEventName}}Handler(addr common.Address, eth *ethclient.Client, callback events.CallbackFn[{{.BindingContract}}{{.BindingEventName}}]) events.Handler {
	b, err := New{{.BindingContract}}(addr, eth)
	if err != nil {
		panic(err)
	}
	return &{{.BindingEventName}}Handler{
		addr: addr.String(),
		binding:  b,
		callback: callback,
	}
}
`
