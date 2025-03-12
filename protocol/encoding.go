package transfer

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"time"
)

// writeFixedString schreibt eine Zeichenkette mit fester Länge
func writeFixedString(buf *bytes.Buffer, s string, length int) {
	b := make([]byte, length)
	copy(b, s) // Falls s kürzer ist, wird der Rest mit 0-Padding gefüllt
	buf.Write(b)
}

// EncodeMsg serialisiert die `Msg`-Struktur in ein Byte-Array
func EncodeMsg(msg MsgHeader) ([]byte, error) {
	buf := new(bytes.Buffer)

	// Header serialisieren
	buf.WriteByte(byte(msg.Version)) // [1]byte
	buf.WriteByte(byte(msg.MsgType)) // [1]byte
	buf.WriteByte(byte(msg.Method))  // [1]byte

	// Timestamp als 8 Byte speichern
	timestampBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(timestampBytes, uint64(msg.Timestamp))
	buf.Write(timestampBytes)

	// Timeout als 4 Byte speichern
	timeoutBytes := make([]byte, 4)
	binary.BigEndian.PutUint32(timeoutBytes, uint32(msg.Timeout.Seconds()))
	buf.Write(timeoutBytes)

	// Strings in feste Byte-Größen konvertieren
	writeFixedString(buf, msg.Domain, 32)
	writeFixedString(buf, msg.Endpoint, 32)
	writeFixedString(buf, msg.Auth, 16)
	buf.WriteByte(byte(msg.PayloadType)) // [1]byte

	// PayloadSize als 64 Byte speichern
	payloadSizeBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(payloadSizeBytes, uint64(msg.PayloadSize))
	buf.Write(payloadSizeBytes)

	return buf.Bytes(), nil
}

func readInt(r io.Reader, n int) (int, error) {
	buf := make([]byte, n)
	if _, err := io.ReadFull(r, buf); err != nil {
		return 0, fmt.Errorf("Fehler beim Lesen eines %d-Byte-Werts: %w", n, err)
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

// readTime liest einen Unix-Timestamp (8 Byte) und konvertiert ihn zu `time.Time`
func readTime(r io.Reader) (time.Time, error) {
	val, err := readInt(r, 8)
	if err != nil {
		return time.Time{}, err
	}
	return time.Unix(int64(val), 0), nil
}

// readString liest einen String fester Länge aus dem Stream
func readString(r io.Reader, length int) (string, error) {
	buf := make([]byte, length)
	if _, err := io.ReadFull(r, buf); err != nil {
		return "", fmt.Errorf("Fehler beim Lesen eines Strings (%d Bytes): %w", length, err)
	}
	return string(buf), nil
}

// readBytes liest eine Länge + Byte-Daten aus dem Stream
func readBytes(r io.Reader, size int) ([]byte, error) {
	buf := make([]byte, size)
	if _, err := io.ReadFull(r, buf); err != nil {
		return nil, fmt.Errorf("Fehler beim Lesen von %d Bytes: %w", size, err)
	}
	return buf, nil
}

// DecodeMsgStream liest eine `Msg`-Struktur direkt aus einem Stream (z. B. TCP)
func DecodeMsgStream(r io.Reader) (MsgHeader, error) {
	var msg MsgHeader
	var err error

	// Header lesen
	if msg.Version, err = readInt(r, 1); err != nil {
		return msg, err
	}
	if msg.MsgType, err = readInt(r, 1); err != nil {
		return msg, err
	}
	if msg.Method, err = readInt(r, 1); err != nil {
		return msg, err
	}

	// Timestamp lesen
	if msg.Timestamp, err = readInt(r, 8); err != nil {
		return msg, err
	}

	// Timeout als Sekunden lesen
	timeoutSec, err := readInt(r, 4)
	if err != nil {
		return msg, err
	}
	msg.Timeout = time.Duration(timeoutSec) * time.Second

	// Strings lesen
	if msg.Domain, err = readString(r, 32); err != nil {
		return msg, err
	}
	if msg.Endpoint, err = readString(r, 32); err != nil {
		return msg, err
	}
	if msg.Auth, err = readString(r, 16); err != nil {
		return msg, err
	}
	if msg.PayloadType, err = readInt(r, 1); err != nil {
		return msg, err
	}

	// PayloadSize lesen
	if msg.PayloadSize, err = readInt(r, 8); err != nil {
		return msg, err
	}

	return msg, nil
}
