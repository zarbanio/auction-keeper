package logger

import (
	"context"
	"io"
	"os"

	"github.com/rs/zerolog"
	"github.com/zarbanio/auction-keeper/store"
)

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
	_, err = cw.istore.CreateLogs(cw.ctx, p)
	if err != nil {
		return 0, err
	}

	return cw.writer.Write(p)
}

type Logger struct {
	ConsoleLogger zerolog.Logger
	Logger        zerolog.Logger
}

func NewLogger(ctx context.Context, s store.IStore) *Logger {
	customWriter := NewCustomWriter(ctx, s, os.Stdout)
	// logger writes logs to the database using CustomWriter
	logger := zerolog.New(customWriter).With().Timestamp().Logger().Output(customWriter)
	// consoleLogger writes logs only to the console, not to the database
	consoleLogger := zerolog.New(os.Stdout).With().Timestamp().Logger()

	return &Logger{
		ConsoleLogger: consoleLogger,
		Logger:        logger,
	}
}
