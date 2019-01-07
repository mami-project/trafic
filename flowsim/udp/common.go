package udp

import (
	"bytes"
	"encoding/binary"
	common "github.com/mami-project/trafic/flowsim/common"
	"time"
)

type myStruct struct {
	pktId   int64
	total   int64
	tStamp  int64
	padding [64*1024 - 24]byte
}

func DecodePacket(pkt []byte) myStruct {
	var result myStruct
	var vuelta int64

	// fmt.Printf("Received packet with %d bytes\n -> % x\n",len(pkt),pkt)
	binbuf := bytes.NewReader(pkt)

	err := binary.Read(binbuf, binary.BigEndian, &vuelta)
	common.FatalError(err)
	result.pktId = vuelta

	err = binary.Read(binbuf, binary.BigEndian, &vuelta)
	common.FatalError(err)
	result.total = vuelta

	err = binary.Read(binbuf, binary.BigEndian, &vuelta)
	common.FatalError(err)
	result.tStamp = vuelta
	// result.padding = pkt[16:]
	return result
}

func EncodePacket(input myStruct, plen int) []byte {
	var binbuf bytes.Buffer
	binary.Write(&binbuf, binary.BigEndian, input)

	return binbuf.Bytes()[:plen]
}

func MakeTimestamp() int64 {
	return time.Now().UnixNano() / (int64(time.Microsecond) / int64(time.Nanosecond))
}

func toTimestamp(t time.Time) int64 {
	// fmt.Println("toTimestamp (",t,")");
	return t.UnixNano() / (int64(time.Microsecond) / int64(time.Nanosecond))
}
