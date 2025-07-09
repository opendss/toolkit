package observabilities

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

type PprofOptions struct {
	BindAddress string `json:"bindAddress,omitempty" yaml:"bind-address,omitempty"`
}

func RunPprof(ctx context.Context, option PprofOptions) io.Closer {
	s := &http.Server{
		Addr:              option.BindAddress,
		Handler:           http.DefaultServeMux,
		ReadHeaderTimeout: time.Second,
	}
	slog.Info("starting pprof http server.", slog.String("addr", option.BindAddress))

	go pprof.Do(ctx, pprof.Labels(PprofLabelComponent, "pprof-server"), func(ctx context.Context) {
		if err := s.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			slog.Error("start pprof http server failed.", slog.Any("error", err))
			os.Exit(1)
		}
	})
	return s
}
