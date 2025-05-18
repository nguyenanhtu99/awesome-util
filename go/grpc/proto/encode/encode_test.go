package encode_test

import (
	"awesome-util/go/grpc/proto/encode"
	"testing"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
)

func TestEncodeSmall(t *testing.T) {
	encodeInt32 := encode.EncodeInt32{Num: 1}
	encodeFixed32 := encode.EncodeFixed32{Num: 1}
	encodeSint32 := encode.EncodeSint32{Num: 1}

	bytesInt32, err := proto.Marshal(&encodeInt32)
	if err != nil {
		t.Fatal(err)
	}

	bytesFixed32, err := proto.Marshal(&encodeFixed32)
	if err != nil {
		t.Fatal(err)
	}

	bytesSint32, err := proto.Marshal(&encodeSint32)
	if err != nil {
		t.Fatal(err)
	}

	t.Log("Encode int32: ", bytesInt32)
	t.Log("Encode fixed32: ", bytesFixed32)
	t.Log("Encode sint32: ", bytesSint32)
}

func TestEncodeBig(t *testing.T) {
	encodeInt32 := encode.EncodeInt32{Num: 1747557659}
	encodeFixed32 := encode.EncodeFixed32{Num: 1747557659}
	encodeSint32 := encode.EncodeSint32{Num: 1747557659}

	bytesInt32, err := proto.Marshal(&encodeInt32)
	if err != nil {
		t.Fatal(err)
	}

	bytesFixed32, err := proto.Marshal(&encodeFixed32)
	if err != nil {
		t.Fatal(err)
	}

	bytesSint32, err := proto.Marshal(&encodeSint32)
	if err != nil {
		t.Fatal(err)
	}

	t.Log("Encode int32: ", bytesInt32)
	t.Log("Encode fixed32: ", bytesFixed32)
	t.Log("Encode sint32: ", bytesSint32)
}

func TestNegative(t *testing.T) {
	encodeInt32 := encode.EncodeInt32{Num: -1}
	encodeSint32 := encode.EncodeSint32{Num: -1}

	bytesInt32, err := proto.Marshal(&encodeInt32)
	if err != nil {
		t.Fatal(err)
	}

	bytesSint32, err := proto.Marshal(&encodeSint32)
	if err != nil {
		t.Fatal(err)
	}

	t.Log("Encode int32: ", bytesInt32)
	t.Log("Encode sint32: ", bytesSint32)
}

func TestRepeated(t *testing.T) {
	encodeRepeated := encode.Repeated{Num: []int32{1, 2, 3}, Name: []string{"a", "b", "c"}}

	bytesRepeated, err := proto.Marshal(&encodeRepeated)
	if err != nil {
		t.Fatal(err)
	}

	t.Log("Encode repeated: ", bytesRepeated)
}

func TestMap(t *testing.T) {
	encodeMap := encode.Map{Map: map[int32]string{1: "a", 2: "b", 3: "c"}}

	bytesMap, err := proto.Marshal(&encodeMap)
	if err != nil {
		t.Fatal(err)
	}

	t.Log("Encode map: ", bytesMap)
}

func TestOneof(t *testing.T) {
	encodeOneof := encode.Oneof{Oneof: &encode.Oneof_Num{1}}

	bytesOneof, err := proto.Marshal(&encodeOneof)
	if err != nil {
		t.Fatal(err)
	}

	t.Log("Encode oneof: ", bytesOneof)
}

func TestReserve(t *testing.T) {
	encodeReserve := encode.Reserve{Num: 1, Name: "a"}

	bytesReserve, err := proto.Marshal(&encodeReserve)
	if err != nil {
		t.Fatal(err)
	}

	t.Log("Encode reserve: ", bytesReserve)
}

func TestAny(t *testing.T) {
	data, err := anypb.New(&encode.EncodeInt32{Num: 1})
	if err != nil {
		t.Fatal(err)
	}
	encodeAny := encode.Any{Any: data}

	bytesAny, err := proto.Marshal(&encodeAny)
	if err != nil {
		t.Fatal(err)
	}

	t.Log("Encode any: ", bytesAny)
}
