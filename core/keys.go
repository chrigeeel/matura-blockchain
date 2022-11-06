package core

import (
	"crypto"
	"crypto/ed25519"
	"crypto/rand"
	"encoding/json"
	"fmt"

	"github.com/mr-tron/base58"
)

const (
	PublicKeyLength = 32
	SignatureLength = 64
)

type PublicKey [PublicKeyLength]byte

var EmptyPublicKey = PublicKey{}

func (p PublicKey) String() string {
	return base58.Encode(p[:])
}

func (p PublicKey) MarshalJSON() ([]byte, error) {
	return json.Marshal(base58.Encode(p[:]))
}

func (p *PublicKey) UnmarshalJSON(data []byte) (err error) {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}

	*p, err = PublicKeyFromBase58(s)
	if err != nil {
		return fmt.Errorf("invalid public key %q: %w", s, err)
	}
	return
}

func (p PublicKey) Equals(pb PublicKey) bool {
	return p == pb
}

func (p PublicKey) ToPointer() *PublicKey {
	return &p
}

func MustPublicKeyFromBase58(in string) PublicKey {
	out, err := PublicKeyFromBase58(in)
	if err != nil {
		panic(err)
	}
	return out
}

func PublicKeyFromBase58(in string) (out PublicKey, err error) {
	val, err := base58.Decode(in)
	if err != nil {
		return out, fmt.Errorf("decode: %w", err)
	}

	if len(val) != PublicKeyLength {
		return out, fmt.Errorf("invalid length, expected %v, got %d", PublicKeyLength, len(val))
	}

	copy(out[:], val)
	return
}

type PrivateKey []byte

func (k PrivateKey) String() string {
	return base58.Encode(k)
}

func (k PrivateKey) PublicKey() PublicKey {
	p := ed25519.PrivateKey(k)
	pub := p.Public().(ed25519.PublicKey)

	var publicKey PublicKey
	copy(publicKey[:], pub)

	return publicKey
}

func MustPrivateKeyFromBase58(in string) PrivateKey {
	out, err := PrivateKeyFromBase58(in)
	if err != nil {
		panic(err)
	}
	return out
}

func PrivateKeyFromBase58(privkey string) (PrivateKey, error) {
	res, err := base58.Decode(privkey)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func NewRandomPrivateKey() (PrivateKey, error) {
	pub, priv, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		return nil, err
	}
	var publicKey PublicKey
	copy(publicKey[:], pub)
	return PrivateKey(priv), nil
}

func (k PrivateKey) Sign(payload []byte) (Signature, error) {
	p := ed25519.PrivateKey(k)
	signData, err := p.Sign(rand.Reader, payload, crypto.Hash(0))
	if err != nil {
		return Signature{}, err
	}

	var signature Signature
	copy(signature[:], signData)

	return signature, err
}

type Signature [SignatureLength]byte

var EmptySignature = Signature{}

func (p Signature) String() string {
	return base58.Encode(p[:])
}

func (sig Signature) IsZero() bool {
	return sig == EmptySignature
}
func (sig Signature) Equals(pb Signature) bool {
	return sig == pb
}

func (p Signature) MarshalJSON() ([]byte, error) {
	return json.Marshal(base58.Encode(p[:]))
}

func (p *Signature) UnmarshalJSON(data []byte) (err error) {
	var s string
	err = json.Unmarshal(data, &s)
	if err != nil {
		return
	}

	dat, err := base58.Decode(s)
	if err != nil {
		return err
	}

	if len(dat) != SignatureLength {
		return fmt.Errorf("invalid length for Signature, expected 64, got %d", len(dat))
	}

	target := Signature{}
	copy(target[:], dat)
	*p = target
	return
}

func MustSignatureFromBase58(in string) Signature {
	out, err := SignatureFromBase58(in)
	if err != nil {
		panic(err)
	}
	return out
}

func SignatureFromBase58(in string) (out Signature, err error) {
	val, err := base58.Decode(in)
	if err != nil {
		return
	}

	if len(val) != SignatureLength {
		err = fmt.Errorf("invalid length, expected 64, got %d", len(val))
		return
	}
	copy(out[:], val)
	return
}

func (s Signature) Verify(pubkey PublicKey, payload []byte) bool {
	return ed25519.Verify(pubkey[:], payload, s[:])
}
