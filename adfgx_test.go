package gocipher

import (
	"fmt"
	"testing"
)

func TestAdfgxNew(t *testing.T) {
	c, err := NewAdfgx(
		[]byte("ABYZ"),
	)

	if c == nil || err != nil {
		t.Fatalf(`could not instantiate`)
	}
}

func TestAdfgxNewErrorAlphabetSize(t *testing.T) {
	c, err := NewAdfgx(
		[]byte("ABYZ"),
		AdfgxWithNewPolybius(func(key []byte, opts ...PolybiusOption) (Cipher, error) {
			opts = append(opts, PolybiusWithAlphabet([]byte("ABCDEFGHIKLMNOPQRSTUVWXY")))
			return NewPolybius(key, opts...)
		}),
	)

	if c != nil || err == nil {
		t.Fatalf("did not fail")
	}
}

func TestAdfgxNewErrorInvalidPolybiusCipher(t *testing.T) {
	c, err := NewAdfgx(
		[]byte("ABYZ"),
		AdfgxWithNewPolybius(func(key []byte, opts ...PolybiusOption) (Cipher, error) {
			return nil, fmt.Errorf("random failure")
		}),
	)

	if c != nil || err == nil {
		t.Fatalf("did not fail")
	}
}

func TestAdfgxNewErrorInvalidTranspositionCipher(t *testing.T) {
	c, err := NewAdfgx(
		[]byte("ABYZ"),
		AdfgxWithNewTransposition(func(key []byte, opts ...TranspositionOption) (Cipher, error) {
			return nil, fmt.Errorf("random failure")
		}),
	)

	if c != nil || err == nil {
		t.Fatalf("did not fail")
	}
}

func TestAdfgxBasicCrypt(t *testing.T) {
	c, err := NewAdfgx(
		[]byte("CRYPTOGRAPHY"),
	)
	if err != nil {
		t.Fatalf("unexpected: could not instantiate")
	}

	testCipherCrypt(c, t,
		[]byte("THEQUICKBROWNFOXJUMPSOVERTHELAZYDOG"),
		[]byte("GFGAAGAFFGXGGXXFDDXFGXXDGGXGFXAFFGXGXXAGFFXGGADFAADAAFFAFDDDGDAGDDDX"),
		[]byte("THEQUICKBROWNFOXUMPSOVERTHELAZYDOG"),
	)
}

func BenchmarkAdfgxEncrypt(b *testing.B) {
	c, _ := NewAdfgx(
		[]byte("CRYPTOGRAPHY"),
	)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = c.Encrypt([]byte("INCRYPTOGRAPHYASUBSTITUTIONCIPHERISAMETHODOFENCRYPTINGINWHICHUNITSOFPLAINTEXTAREREPLACEDWITHTHECIPHERTEXTINADEFINEDMANNERWITHTHEHELPOFAKEY"))
	}
}

func BenchmarkAdfgxDecrypt(b *testing.B) {
	c, _ := NewAdfgx(
		[]byte("CRYPTOGRAPHY"),
	)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = c.Decrypt([]byte("XAGGDGFAFDGAGFXAGFFFDFADGDGDDFFXDDFFGAGDAAAGDFGGGGADDFDDDFXAADADDFGADFFAGFAAGDAGDAFDAGAAAGFXFDAGFDGXGDFXXDXGFGAAGFAFGGGGXXAXFXAGXFFXFXDDXAGAXXGFGFFGGAGXDXDFFFFAXGGFFGFDGGGFGFDAGGXGGDFGDDFGFGFFGFGAFXGFXGGFGXXADADFGGAGXFFAGAGDGDADDAFFXGFAADFFGDGAADFXAFXADXXDGXAGDDFFGAAGFGAXXGXG"))
	}
}
