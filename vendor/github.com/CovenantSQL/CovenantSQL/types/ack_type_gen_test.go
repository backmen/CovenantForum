package types

// Code generated by github.com/CovenantSQL/HashStablePack DO NOT EDIT.

import (
	"bytes"
	"crypto/rand"
	"encoding/binary"
	"testing"
)

func TestMarshalHashAck(t *testing.T) {
	v := Ack{}
	binary.Read(rand.Reader, binary.BigEndian, &v)
	bts1, err := v.MarshalHash()
	if err != nil {
		t.Fatal(err)
	}
	bts2, err := v.MarshalHash()
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(bts1, bts2) {
		t.Fatal("hash not stable")
	}
}

func BenchmarkMarshalHashAck(b *testing.B) {
	v := Ack{}
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		v.MarshalHash()
	}
}

func BenchmarkAppendMsgAck(b *testing.B) {
	v := Ack{}
	bts := make([]byte, 0, v.Msgsize())
	bts, _ = v.MarshalHash()
	b.SetBytes(int64(len(bts)))
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bts, _ = v.MarshalHash()
	}
}

func TestMarshalHashAckHeader(t *testing.T) {
	v := AckHeader{}
	binary.Read(rand.Reader, binary.BigEndian, &v)
	bts1, err := v.MarshalHash()
	if err != nil {
		t.Fatal(err)
	}
	bts2, err := v.MarshalHash()
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(bts1, bts2) {
		t.Fatal("hash not stable")
	}
}

func BenchmarkMarshalHashAckHeader(b *testing.B) {
	v := AckHeader{}
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		v.MarshalHash()
	}
}

func BenchmarkAppendMsgAckHeader(b *testing.B) {
	v := AckHeader{}
	bts := make([]byte, 0, v.Msgsize())
	bts, _ = v.MarshalHash()
	b.SetBytes(int64(len(bts)))
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bts, _ = v.MarshalHash()
	}
}

func TestMarshalHashAckResponse(t *testing.T) {
	v := AckResponse{}
	binary.Read(rand.Reader, binary.BigEndian, &v)
	bts1, err := v.MarshalHash()
	if err != nil {
		t.Fatal(err)
	}
	bts2, err := v.MarshalHash()
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(bts1, bts2) {
		t.Fatal("hash not stable")
	}
}

func BenchmarkMarshalHashAckResponse(b *testing.B) {
	v := AckResponse{}
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		v.MarshalHash()
	}
}

func BenchmarkAppendMsgAckResponse(b *testing.B) {
	v := AckResponse{}
	bts := make([]byte, 0, v.Msgsize())
	bts, _ = v.MarshalHash()
	b.SetBytes(int64(len(bts)))
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bts, _ = v.MarshalHash()
	}
}

func TestMarshalHashSignedAckHeader(t *testing.T) {
	v := SignedAckHeader{}
	binary.Read(rand.Reader, binary.BigEndian, &v)
	bts1, err := v.MarshalHash()
	if err != nil {
		t.Fatal(err)
	}
	bts2, err := v.MarshalHash()
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(bts1, bts2) {
		t.Fatal("hash not stable")
	}
}

func BenchmarkMarshalHashSignedAckHeader(b *testing.B) {
	v := SignedAckHeader{}
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		v.MarshalHash()
	}
}

func BenchmarkAppendMsgSignedAckHeader(b *testing.B) {
	v := SignedAckHeader{}
	bts := make([]byte, 0, v.Msgsize())
	bts, _ = v.MarshalHash()
	b.SetBytes(int64(len(bts)))
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bts, _ = v.MarshalHash()
	}
}
