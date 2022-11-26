package gocipher

import (
	"testing"
)

func TestAdfgxNew(t *testing.T) {
	c, err := NewAdfgx([]byte("ABCDEFGHIKLMNOPQRSTUVWXYZ"), []byte("ABYZ"))

	if c == nil || err != nil {
		t.Fatalf(`could not instantiate`)
	}
}

func TestAdfgxNewErrorAlphabetSize(t *testing.T) {
	c, err := NewAdfgx([]byte("ABCDEFGHIKLMNOPQRSTUVWXY"), []byte("ABYZ"))

	if c != nil || err == nil {
		t.Fatalf("did not fail")
	}
}

func TestAdfgxBasicCrypt(t *testing.T) {
	c, err := NewAdfgx([]byte("ABCDEFGHIKLMNOPQRSTUVWXYZ"), []byte("CRYPTOGRAPHY"))
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
	c, _ := NewAdfgx([]byte("ABCDEFGHIKLMNOPQRSTUVWXYZ"), []byte("CRYPTOGRAPHY"))

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = c.Encrypt([]byte("INCRYPTOGRAPHYASUBSTITUTIONCIPHERISAMETHODOFENCRYPTINGINWHICHUNITSOFPLAINTEXTAREREPLACEDWITHTHECIPHERTEXTINADEFINEDMANNERWITHTHEHELPOFAKEY"))
	}
}

func BenchmarkAdfgxDecrypt(b *testing.B) {
	c, _ := NewAdfgx([]byte("ABCDEFGHIKLMNOPQRSTUVWXYZ"), []byte("CRYPTOGRAPHY"))

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = c.Decrypt([]byte("XAGGDGFAFDGAGFXAGFFFDFADGDGDDFFXDDFFGAGDAAAGDFGGGGADDFDDDFXAADADDFGADFFAGFAAGDAGDAFDAGAAAGFXFDAGFDGXGDFXXDXGFGAAGFAFGGGGXXAXFXAGXFFXFXDDXAGAXXGFGFFGGAGXDXDFFFFAXGGFFGFDGGGFGFDAGGXGGDFGDDFGFGFFGFGAFXGFXGGFGXXADADFGGAGXFFAGAGDGDADDAFFXGFAADFFGDGAADFXAFXADXXDGXAGDDFFGAAGFGAXXGXG"))
	}
}