package kafkav1

import "testing"

func TestService_Deal(t *testing.T) {
	srv := New("test1", "goim-zb_otc-group", 16, []string{"127.0.0.1:9092"})
	srv.ListenMQ()
	select {

	}
}
