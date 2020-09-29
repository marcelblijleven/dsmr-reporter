package dsmr

import (
	"fmt"
	"github.com/howeyc/crc16"
)

var (
	polynomial uint16 = 0xA001
	crc16Table        = crc16.MakeTableNoXOR(polynomial)
)

// GenerateCRC16Checksum returns a checksum as 4 hexadecimal characters of the provided slice of bytes
// using a table with the IBM (0xA001) polynomial and no XOR in or XOR out
func GenerateCRC16Checksum(data []byte) string {
	checksum := crc16.Checksum(data, crc16Table)
	return fmt.Sprintf("%04X", checksum)
}
