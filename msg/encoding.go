package msg

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"hash/crc32"
	"io"
	"time"
)

/*
Encoding Package for decoding and encoding protocol for inter node communication of fileserver nodes.

*/

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

func WriteboolToByte(b bool) byte {
	if b {
		return 1
	}
	return 0
}

func computeCRC32Checksum(data []byte) uint32 {
	return crc32.ChecksumIEEE(data)
}

func WriteChecksum(buffer *bytes.Buffer) {
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
	buffer.WriteByte(byte(msg.MsgType)) // [1]byte

	errorBytes := make([]byte, 4)
	binary.BigEndian.PutUint16(errorBytes, uint16(msg.Error))
	buffer.Write(errorBytes)

	buffer.WriteByte(byte(msg.Method)) // [1]byte

	timestampBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(timestampBytes, uint64(msg.Timestamp))
	buffer.Write(timestampBytes)

	timeoutBytes := make([]byte, 4)
	binary.BigEndian.PutUint32(timeoutBytes, uint32(msg.Timeout.Seconds()))
	buffer.Write(timeoutBytes)

	writeFixedString(buffer, msg.Domain, 32)
	writeFixedString(buffer, msg.Endpoint, 32)

	buffer.WriteByte(WriteboolToByte(msg.HasAuth))

	if msg.HasAuth {
		writeFixedString(buffer, msg.Auth, 16)
	}
	buffer.WriteByte(WriteboolToByte(msg.HasPayload))

	if msg.HasPayload {

		buffer.WriteByte(byte(msg.PayloadType)) 

		payloadSizeBytes := make([]byte, 8)
		binary.BigEndian.PutUint64(payloadSizeBytes, uint64(msg.PayloadSize))
		buffer.Write(payloadSizeBytes)
	}

	WriteChecksum(buffer)

	return buffer.Bytes(), nil
}

// readInt reads an integer of the specified byte size from a stream.
// The function supports reading 1-byte, 4-byte, and 8-byte integers in big-endian format.
func readInt(r io.Reader, n int) (int, error) {
	buf := make([]byte, n)
	if _, err := io.ReadFull(r, buf); err != nil {
		return 0, fmt.Errorf("fehler beim lesen eines %d-Byte-Werts: %w", n, err)
	}
	switch n {
	case 1:
		return int(buf[0]), nil
	case 4:
		return int(binary.BigEndian.Uint32(buf)), nil
	case 8:
		return int(binary.BigEndian.Uint64(buf)), nil
	default:
		return 0, fmt.Errorf("ungültige Byte-Größe: %d", n)
	}
}

func readString(r io.Reader, length int) (string, error) {
	buf := make([]byte, length)
	if _, err := io.ReadFull(r, buf); err != nil {
		return "", fmt.Errorf("fehler beim lesen eines strings (%d Bytes): %w", length, err)
	}
	return string(buf), nil
}

func readBool(r io.Reader) (bool, error) {
	var buffer [1]byte
	_, err := r.Read(buffer[:])
	if err != nil {
		return false, fmt.Errorf("error reading byte: %w", err)
	}

	switch buffer[0] {
	case 0:
		return false, nil
	case 1:
		return true, nil
	default:
		return false, fmt.Errorf("invalid boolean byte: %d", buffer[0])
	}
}

func ValidateCRC32Checksum(data []byte, expectedChecksum uint32) bool {
	return computeCRC32Checksum(data) == expectedChecksum
}

// DecodeMsgStream reads a MsgHeader structure from a stream.
// This function reads a binary stream and decodes it into a MsgHeader structure,
// following the expected field sizes and format.
func DecodeMsgHeader(reader io.Reader) (MsgHeader, error) {
	headerBuffer := bytes.NewBuffer([]byte{})
	buffreader := io.TeeReader(reader, headerBuffer)
	var msg MsgHeader
	var err error

	if msg.Version, err = readInt(buffreader, 1); err != nil {
		return msg, err
	}

	if msg.MsgType, err = readInt(buffreader, 1); err != nil {
		return msg, err
	}

	if msg.Error, err = readInt(buffreader, 4); err != nil {
		return msg, err
	}

	if msg.Method, err = readInt(buffreader, 1); err != nil {
		return msg, err
	}

	if msg.Timestamp, err = readInt(buffreader, 8); err != nil {
		return msg, err
	}

	timeoutSec, err := readInt(buffreader, 4)
	if err != nil {
		return msg, err
	}
	msg.Timeout = time.Duration(timeoutSec) * time.Second

	if msg.Domain, err = readString(buffreader, 32); err != nil {
		return msg, fmt.Errorf("error with domain: %w", err)
	}
	if msg.Endpoint, err = readString(buffreader, 32); err != nil {
		return msg, fmt.Errorf("error with endpoint: %w", err)
	}

	if msg.HasAuth, err = readBool(buffreader); err != nil {
		return msg, fmt.Errorf("error with hasAuth: %w", err)
	} else {
		if msg.HasAuth {
			if msg.Auth, err = readString(buffreader, 16); err != nil {
				return msg, fmt.Errorf("error with auth: %w", err)
			}
		}

	}

	if msg.HasPayload, err = readBool(buffreader); err != nil {
		return msg, fmt.Errorf("error with haspayload: %w", err)
	} else {
		if msg.HasPayload {
			if msg.PayloadType, err = readInt(buffreader, 1); err != nil {
				return msg, fmt.Errorf("error with payloadtype: %w", err)
			}

			if msg.PayloadSize, err = readInt(buffreader, 8); err != nil {
				return msg, fmt.Errorf("error with payloadsize: %w", err)
			}
		}
	}

	if msg.Checksum, err = readInt(reader, 4); err != nil {
		return msg, err
	}

	isValid := ValidateCRC32Checksum(headerBuffer.Bytes(), uint32(msg.Checksum))
	if !isValid {
		return msg, fmt.Errorf("invalid msg header")
	}

	headerBuffer.Reset()

	return msg, nil
}
