package services

import (
	"context"
	"io"
	"os"

	"github.com/rs/zerolog"
	"github.com/zarbanio/auction-keeper/store"
)

var Logger zerolog.Logger

type CustomWriter struct {
	ctx    context.Context
	writer io.Writer
	istore store.IStore
}

func NewCustomWriter(ctx context.Context, istore store.IStore, writer io.Writer) *CustomWriter {
	return &CustomWriter{
		ctx:    ctx,
		writer: writer,
		istore: istore,
	}
}

func (cw *CustomWriter) Write(p []byte) (n int, err error) {
	_, err = cw.istore.CreateSysLog(cw.ctx, p)
	if err != nil {
		return 0, err
	}

	return cw.writer.Write(p)
}

func InitSysLogger(ctx context.Context, s store.IStore) error {
	customWriter := NewCustomWriter(ctx, s, os.Stdout)

	Logger = zerolog.New(customWriter).With().Caller().Logger().Output(
		customWriter)
	return nil
}
