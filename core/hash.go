package core

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"

	"github.com/mr-tron/base58"
)

const (
	HashLength = 32
)

type Hash [HashLength]byte

func (h Hash) String() string {
	return base58.Encode(h[:])
}

func (h Hash) MarshalJSON() ([]byte, error) {
	return json.Marshal(base58.Encode(h[:]))
}

func (h *Hash) UnmarshalJSON(data []byte) (err error) {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}

	*h, err = HashFromBase58(s)
	if err != nil {
		return fmt.Errorf("invalid hash %q: %w", s, err)
	}
	return
}

func (h Hash) MarshalText() ([]byte, error) {
	return []byte(base58.Encode(h[:])), nil
}

func (h *Hash) UnmarshalText(data []byte) error {
	return h.Set(string(data))
}

func (h *Hash) Set(s string) (err error) {
	*h, err = HashFromBase58(s)
	if err != nil {
		return fmt.Errorf("invalid hash %s: %w", s, err)
	}
	return
}

func MustHashFromBase58(in string) Hash {
	out, err := HashFromBase58(in)
	if err != nil {
		panic(err)
	}
	return out
}

func HashFromBase58(in string) (out Hash, err error) {
	val, err := base58.Decode(in)
	if err != nil {
		return out, fmt.Errorf("decode: %w", err)
	}

	if len(val) != HashLength {
		return out, fmt.Errorf("invalid length, expected %v, got %d", HashLength, len(val))
	}

	copy(out[:], val)
	return
}

func HashBytes(data []byte) Hash {
	return sha256.Sum256(data)
}

func HashString(data string) Hash {
	return sha256.Sum256([]byte(data))
}
