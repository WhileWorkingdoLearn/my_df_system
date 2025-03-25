package msg

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"strconv"
	"strings"
)

/*
Encoding Package for decoding and encoding protocol for inter node communication of fileserver nodes.
*/
const (
	none int = iota
	initialized
	done
)

// writeFixedString writes a fixed-length string to a buffer.
// If the input is shorter than the specified length, the remaining space
// is filled with null bytes (zero-padding).
// Parameters:
//   - buf: A pointer to a bytes.Buffer where the fixed-length string will be written.
//   - s: The input string to be written into the buffer.
//   - length: The fixed length of the string in bytes.
func writeFixedString(buf *bytes.Buffer, s string, length int) {
	b := make([]byte, length)
	copy(b, s) // Falls s kürzer ist, wird der Rest mit 0-Padding gefüllt
	buf.Write(b)
}

func writeboolToByte(b bool) byte {
	if b {
		return 1
	}
	return 0
}

func writeChecksum(buffer *bytes.Buffer) {
	checksum := computeCRC32Checksum(buffer.Bytes())
	checksumbytes := make([]byte, 4)
	binary.BigEndian.PutUint32(checksumbytes, checksum)
	buffer.Write(checksumbytes)
}

// EncodeMsg serializes a MsgHeader structure into a byte slice.
// The function converts each field of the MsgHeader into its binary representation
// using big-endian format and writes it into a byte buffer.
// Parameters:
//   - msg: The MsgHeader structure to be encoded.
//
// Returns:
//   - A byte slice containing the serialized data.
//   - An error if encoding fails.

func EncodeMsgHeader(msg MsgHeader) ([]byte, error) {
	buffer := new(bytes.Buffer)

	buffer.WriteByte(byte(msg.Version)) // [1]byte
	buffer.WriteByte(byte(':'))

	buffer.WriteByte(byte(msg.MsgType)) // [1]byte
	buffer.WriteByte(byte(':'))

	errorBytes := make([]byte, 4)
	binary.BigEndian.PutUint16(errorBytes, uint16(msg.Error))
	buffer.Write(errorBytes)
	buffer.WriteByte(byte(':'))

	buffer.WriteByte(byte(msg.Method)) // [1]byte
	buffer.WriteByte(byte(':'))

	timestampBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(timestampBytes, uint64(msg.Timestamp.Unix()))
	buffer.Write(timestampBytes)
	buffer.WriteByte(byte(':'))

	timeoutBytes := make([]byte, 4)
	binary.BigEndian.PutUint32(timeoutBytes, uint32(msg.Timeout.Seconds()))
	buffer.Write(timeoutBytes)
	buffer.WriteByte(byte(':'))

	writeFixedString(buffer, msg.Domain, 32)
	buffer.WriteByte(byte(':'))
	writeFixedString(buffer, msg.Endpoint, 32)
	buffer.WriteByte(byte(':'))

	buffer.WriteByte(writeboolToByte(msg.HasAuth))
	buffer.WriteByte(byte(':'))

	if msg.HasAuth {
		writeFixedString(buffer, msg.Auth, 16)
		buffer.WriteByte(byte(':'))
	}
	buffer.WriteByte(writeboolToByte(msg.HasPayload))
	buffer.WriteByte(byte(':'))

	if msg.HasPayload {

		buffer.WriteByte(byte(msg.PayloadType))
		buffer.WriteByte(byte(':'))

		payloadSizeBytes := make([]byte, 8)
		binary.BigEndian.PutUint64(payloadSizeBytes, uint64(msg.PayloadSize))
		buffer.Write(payloadSizeBytes)
		buffer.WriteByte(byte(':'))
	}

	writeChecksum(buffer)
	buffer.Write([]byte(":-:"))

	return buffer.Bytes(), nil
}

type parser struct {
	step   int
	parsed int
	buffer *bytes.Buffer
	done   bool
}

func (p *parser) parse(data string) (int, error) {
	idx := strings.Index(data, ":")
	if idx == 0 {
		p.done = true
		return 0, nil
	}

	switch p.step {
	case 0:
		versionAsString := strings.TrimSpace(data[:idx])
		fmt.Println(versionAsString)
		if version, err := strconv.ParseInt(versionAsString, 10, 64); err != nil {
			return 0, fmt.Errorf("error with converting version")
		} else {
			if p.buffer != nil {
				p.buffer.WriteByte(byte(version))
			}
		}
		p.step++
	case 1:
		p.done = true
	default:
		p.done = true
		return 0, fmt.Errorf("parsing not possible")
	}

	return idx, nil
}

func EncodeMsgHeaderFromString(data string) ([]byte, error) {
	end := strings.Index(data, ":-:")
	if end == -1 {
		return nil, fmt.Errorf("no data end was provided")
	}

	dataToParse := data[:end]

	//buffer := new(bytes.Buffer)
	p := parser{buffer: bytes.NewBuffer(make([]byte, 0))}
	parsedFromString := 0
	for !p.done {
		parsed, err := p.parse(dataToParse[parsedFromString:])
		if err != nil {
			return nil, err
		}
		parsedFromString += parsed
		fmt.Println(parsed)
	}
	fmt.Println(p.buffer.Bytes())

	return p.buffer.Bytes(), nil
}
