package service

import (
	"context"
	"example/load_balance_test/lb_demo3/taotie/config"
	"flag"
	"os"
	"testing"
	"time"
)

func TestMain(m *testing.M) {
	flag.Parse()
	if err := config.Init(); err != nil {
		panic(err)
	}
	os.Exit(m.Run())
}

func Benchmark_GetID(b *testing.B) {
	srv := New(config.Conf)
	b.ResetTimer()
	b.N = 30
	for i := 0; i < b.N; i++ {
		id, err := srv.GetLogId(context.Background(), int64(i))
		if err != nil {
			srv.log.Error().Err(err).Msg("")
		}
		srv.log.Info().Int64("id", id).Msg("")
		time.Sleep(100 * time.Millisecond)
	}
}
