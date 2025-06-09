package observability

import (
	"context"
	"errors"
	"io"
	"log/slog"
	"net/http"
	"os"
	"runtime/pprof"
	"time"
)

const (
	PprofLabelComponent string = "component"
)

func RunPprof(ctx context.Context, addr string) io.Closer {
	s := &http.Server{
		Addr:              addr,
		Handler:           http.DefaultServeMux,
		ReadHeaderTimeout: time.Second,
	}
	slog.Info("starting pprof http server.", slog.String("addr", addr))

	go pprof.Do(ctx, pprof.Labels(PprofLabelComponent, "pprof-server"), func(ctx context.Context) {
		if err := s.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			slog.Error("start pprof http server failed.", slog.Any("error", err))
			os.Exit(1)
		}
	})
	return s
}
