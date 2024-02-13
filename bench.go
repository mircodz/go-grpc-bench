package main

import (
	"fmt"
	"os"
	"reflect"
	"testing"
	"text/tabwriter"

	"google.golang.org/protobuf/proto"

	"github.com/CrowdStrike/csproto"

	"github.com/mircodezorzi/go-proto-bench/pb"
)

type Message interface {
	MarshalVT() ([]byte, error)
	MarshalToVT(dAtA []byte) (int, error)
	UnmarshalVT(dAtA []byte) error

	csproto.Marshaler
	csproto.MarshalerTo
	csproto.Unmarshaler
	proto.Message
}

func main() {
	w := tabwriter.NewWriter(os.Stdout, 1, 1, 5, ' ', 0)

	for _, bench := range []struct {
		name    string
		message func() Message
		t       reflect.Type
	}{
		{
			name: "timeseries - small",
			message: func() Message {
				m := &pb.TimeSeries{}
				m.Labels = randLabels(100, 10)
				m.Samples = randSamples(100)
				return m
			},
			t: reflect.TypeOf(pb.TimeSeries{}),
		},
		{
			name: "timeseries - medium",
			message: func() Message {
				m := &pb.TimeSeries{}
				m.Labels = randLabels(10000, 10)
				m.Samples = randSamples(10000)
				return m
			},
			t: reflect.TypeOf(pb.TimeSeries{}),
		},
		{
			name: "timeseries - large",
			message: func() Message {
				m := &pb.TimeSeries{}
				m.Labels = randLabels(1000000, 10)
				m.Samples = randSamples(1000000)
				return m
			},
			t: reflect.TypeOf(pb.TimeSeries{}),
		},
		{
			name: "address book - small",
			message: func() Message {
				m := &pb.AddressBook{}
				m.Folders = randomFolders(10, 10, 2)
				return m
			},
			t: reflect.TypeOf(pb.AddressBook{}),
		},
		{
			name: "address book - medium",
			message: func() Message {
				m := &pb.AddressBook{}
				m.Folders = randomFolders(100, 100, 10)
				return m
			},
			t: reflect.TypeOf(pb.AddressBook{}),
		},
		{
			name: "address book - large",
			message: func() Message {
				m := &pb.AddressBook{}
				m.Folders = randomFolders(1000, 1000, 10)
				return m
			},
			t: reflect.TypeOf(pb.AddressBook{}),
		},
	} {
		message := bench.message()

		marshalVt := testing.Benchmark(func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_, _ = message.MarshalVT()
			}
		})

		marshalCs := testing.Benchmark(func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_, _ = message.Marshal()
			}
		})

		marshalProto := testing.Benchmark(func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_, _ = proto.Marshal(message)
			}
		})

		bs, _ := proto.Marshal(message)

		unmarshalVt := testing.Benchmark(func(b *testing.B) {
			m := reflect.New(bench.t).Interface().(Message)
			for i := 0; i < b.N; i++ {
				_ = m.UnmarshalVT(bs)
			}
		})

		unmarshalCs := testing.Benchmark(func(b *testing.B) {
			m := reflect.New(bench.t).Interface().(Message)
			for i := 0; i < b.N; i++ {
				_ = m.Unmarshal(bs)
			}
		})

		unmarshalProto := testing.Benchmark(func(b *testing.B) {
			m := reflect.New(bench.t).Interface().(Message)
			for i := 0; i < b.N; i++ {
				_ = proto.Unmarshal(bs, m)
			}
		})

		fmt.Fprintf(w, "%s\t%s\t%d ns/op\t%d ns/op\t%s/op\n", bench.name, "vtproto", marshalVt.NsPerOp(), unmarshalVt.NsPerOp(), Bytes(unmarshalVt.MemBytes/uint64(unmarshalVt.N)))
		fmt.Fprintf(w, "\t%s\t%d ns/op\t%d ns/op\t%s/op\n", "csproto", marshalCs.NsPerOp(), unmarshalCs.NsPerOp(), Bytes(unmarshalCs.MemBytes/uint64(unmarshalCs.N)))
		fmt.Fprintf(w, "\t%s\t%d ns/op\t%d ns/op\t%s/op\n", "proto", marshalProto.NsPerOp(), unmarshalProto.NsPerOp(), Bytes(unmarshalProto.MemBytes/uint64(unmarshalProto.N)))
	}

	w.Flush()
}
