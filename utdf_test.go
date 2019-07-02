package utdfgo

import (
	"bytes"
	"crypto/rand"
	"encoding/binary"
	"io/ioutil"
	"testing"
)

type UDTFPacket struct {
	b1      byte
	b2      byte
	b3      byte
	router1 uint8
	router2 uint8
	YY      uint8
	sic     uint16
	vid     uint16
	scY     uint32
	mcS     uint32
}

var packet UDTFPacket

const pl = 75

// build fake utdf packet for testing
func TestBuildUTDF(t *testing.T) {
	packet = UDTFPacket{
		b1:      1,
		b2:      2,
		b3:      3,
		router1: 'D',
		router2: 'D',
		YY:      19,
		sic:     3675,
		vid:     1,
		scY:     3772500,
		mcS:     0,
	}

	buf := new(bytes.Buffer)

	err := binary.Write(buf, binary.BigEndian, packet)
	if err != nil {
		t.Fatal("Cannot package utdf packet")
	}

	data := make([]byte, 57)
	rand.Read(data)

	err = binary.Write(buf, binary.BigEndian, data)
	if err != nil {
		t.Fatal("Cannot package utdf packet")
	}

	err = ioutil.WriteFile("utdf", buf.Bytes(), 0644)
	if err != nil {
		t.Fatal("Cannot write packet to file")
	}

}

// ensure the first 13 bytes are correct
func TestUTDFHeader(t *testing.T) {
	p := Run("utdf")
	op := p[0]

	// make sure utdf packet is 75 bytes
	if len(op) != pl {
		t.Error("Packet in file not equal to", pl)
	}

	// packet year
	if op.GetYear() != 2000+int(+packet.YY) {
		t.Error("Unexpected Year:", op.GetYear())
	}

	// packet router
	if op[3] != packet.router1 && op[4] != packet.router2 {
		t.Error("Unexpecting router:", op[3])
	}

	// packet SIC
	if op.GetSIC() != uint64(packet.sic) {
		t.Error("Unexpected SIC:", op.GetSIC())
	}

	// packet VID
	if op.GetVID() != uint64(packet.vid) {
		t.Error("Unexpected VID:", op.GetVID())
	}

	// packet seconds
	if op.GetSeconds() != int(packet.scY) {
		t.Error("Unexpected seconds of year:", op.GetSeconds())
	}

	// packet microseconds
	if op.GetMicroseconds() != int(packet.mcS) {
		t.Error("Unexpected microseconds:", op.GetMicroseconds())
	}
}
