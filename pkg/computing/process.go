package computing

import (
	"io"
	"log/slog"
	"os"
	"os/signal"
)

func RunProcessWaitSignal(fn func() (io.Closer, error), sig ...os.Signal) {
	closer, err := fn()
	if err != nil {
		slog.Error("start process wait failed", slog.Any("error", err))
		os.Exit(1)
	}
	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, sig...)

	receivedSig := <-signalCh
	slog.Info("received signal, closing process", slog.String("signal", receivedSig.String()))

	if closer != nil {
		if err = closer.Close(); err != nil {
			slog.Error("close process resources failed", slog.Any("error", err))
			os.Exit(1)
		}
	}

	slog.Info("closed process resources", slog.String("signal", receivedSig.String()))
	os.Exit(0)
}
