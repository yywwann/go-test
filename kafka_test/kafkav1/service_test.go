package kafkav1

import "testing"

func TestService_Deal(t *testing.T) {
	srv := New("test-topic", "group1", 1, []string{"172.16.101.107:9092"})
	srv.ListenMQ()
	select {}
}
