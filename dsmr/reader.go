package dsmr

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

func Read(br *bufio.Reader, reporter Reporter, tz string) {
	for {
		// Peek the first byte to check if it's the start of the message
		if b, err := br.Peek(1); err == nil {
			if string(b) != "/" {
				// Consume the byte
				_, err := br.ReadByte()

				if err != nil {
					if err == io.EOF {
						reporter.Log("EOF reached")
						break
					}

					reporter.Error(err)
					continue
				}
				continue
			}
		} else {
			if err == io.EOF {
				reporter.Log("EOF reached")
				break
			}

			reporter.Error(err)
			continue
		}

		// From this point, the next byte will be "/" which is the start of the telegram
		// Read until '!' to get the entire telegram
		data, err := br.ReadBytes('!')

		if err != nil {
			reporter.Error(err)
			continue
		}

		// Read until newline to get the crc data
		crcData, err := br.ReadBytes('\n')

		if err != nil {
			reporter.Error(err)
			continue
		}

		checksum := GenerateCRC16Checksum(data)
		crc := strings.TrimSpace(string(crcData))

		if checksum != crc {
			reporter.Error(fmt.Errorf("crc does not match, expected %q got %q", crc, checksum))
			continue
		}

		telegram, err := ParseTelegram(data, tz)

		if err != nil {
			reporter.Error(err)
			continue
		}

		reporter.Update(telegram)
	}
}
