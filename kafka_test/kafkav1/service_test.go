package kafkav1

import "testing"

func TestService_Deal(t *testing.T) {
	srv := New("test-topic", "group1", 1, []string{"10.89.3.6:9092"})
	srv.ListenMQ()
	select {}
}
