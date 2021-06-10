package kafka

import (
	"strconv"
	"testing"
)

//func BenchmarkPubv1_Send(b *testing.B) {
//	pub := &Pubv1{
//		KafkaPub: NewPubv1([]string{"127.0.0.1:9092"}),
//	}
//	b.ResetTimer()
//	for i:=0; i < b.N; i++ {
//		pub.Send(strconv.Itoa(i), "pubv1", strconv.Itoa(i))
//	}
//}

func BenchmarkPubv2_1(b *testing.B) {
	pub := &Pubv2{
		KafkaPub: NewPubv2([]string{"127.0.0.1:9092"}),
	}
	b.ResetTimer()
	for i:=0; i < b.N; i++ {
		pub.Send(strconv.Itoa(i), "goim-zb_otc-topic", strconv.Itoa(i))
	}
}

func BenchmarkPubv2_test1(b *testing.B) {
	pub := &Pubv2{
		KafkaPub: NewPubv2([]string{"127.0.0.1:9092"}),
	}
	b.ResetTimer()
	for i:=0; i < b.N; i++ {
		pub.Send(strconv.Itoa(i), "test1", strconv.Itoa(i))
	}
}

func BenchmarkPubv2_test2(b *testing.B) {
	pub := &Pubv2{
		KafkaPub: NewPubv2([]string{"127.0.0.1:9092"}),
	}
	b.ResetTimer()
	for i:=0; i < b.N; i++ {
		pub.Send(strconv.Itoa(i), "test2", strconv.Itoa(i))
	}
}

func BenchmarkPubv2_test4(b *testing.B) {
	pub := &Pubv2{
		KafkaPub: NewPubv2([]string{"127.0.0.1:9092"}),
	}
	b.ResetTimer()
	for i:=0; i < b.N; i++ {
		pub.Send(strconv.Itoa(i), "test4", strconv.Itoa(i))
	}
}

func BenchmarkPubv2_test8(b *testing.B) {
	pub := &Pubv2{
		KafkaPub: NewPubv2([]string{"127.0.0.1:9092"}),
	}
	b.ResetTimer()
	for i:=0; i < b.N; i++ {
		pub.Send(strconv.Itoa(i), "test8", strconv.Itoa(i))
	}
}

func BenchmarkPubv2_test16(b *testing.B) {
	pub := &Pubv2{
		KafkaPub: NewPubv2([]string{"127.0.0.1:9092"}),
	}
	b.ResetTimer()
	for i:=0; i < b.N; i++ {
		pub.Send(strconv.Itoa(i), "test16", strconv.Itoa(i))
	}
}