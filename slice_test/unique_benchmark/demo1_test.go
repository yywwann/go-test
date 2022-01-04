package main

import (
	"math/rand"
	"reflect"
	"testing"
	"time"
)

func TestUniqueIntsWithMap(t *testing.T) {
	type args struct {
		slice []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "1",
			args: args{
				slice: []int{1, 1, 1},
			},
			want: []int{1},
		},
		{
			name: "2",
			args: args{
				slice: []int{1, 2, 2, 3, 3, 4},
			},
			want: []int{1, 2, 3, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UniqueIntsWithMap(tt.args.slice); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UniqueIntsWithMap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUniqueIntsWithSort(t *testing.T) {
	type args struct {
		slice []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "1",
			args: args{
				slice: []int{1, 1, 1},
			},
			want: []int{1},
		},
		{
			name: "2",
			args: args{
				slice: []int{1, 2, 2, 3, 3, 4},
			},
			want: []int{1, 2, 3, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UniqueIntsWithSort(tt.args.slice); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UniqueIntsWithSort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func generateWithCap(n int) []int {
	rand.Seed(time.Now().UnixNano())
	nums := make([]int, 0, n)
	for i := 0; i < n; i++ {
		nums = append(nums, rand.Int()%100)
	}
	return nums
}

var tmpSlice = generateWithCap(100000)

func BenchmarkUniqueIntsWithMap(b *testing.B) {
	for n := 0; n < b.N; n++ {
		UniqueIntsWithMap(tmpSlice)
	}
}

func BenchmarkUniqueIntsWithSort(b *testing.B) {
	for n := 0; n < b.N; n++ {
		UniqueIntsWithSort(tmpSlice)
	}
}

func BenchmarkUniqueIntsWithSort2(b *testing.B) {
	for n := 0; n < b.N; n++ {
		UniqueIntsWithSort2(tmpSlice)
	}
}
