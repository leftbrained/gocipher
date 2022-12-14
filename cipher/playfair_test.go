package cipher

import (
	"testing"
)

func TestPlayfairNew(t *testing.T) {
	c, err := NewPlayfair(
		[]byte("CRYPTOGRAPHY"),
	)

	if c == nil || err != nil {
		t.Fatalf(`could not instantiate: %q`, err.Error())
	}
}

func TestPlayfairNewErrorNotSquare(t *testing.T) {
	c, err := NewPlayfair(
		[]byte("CRYPTOGRAPHY"),
		PlayfairWithAlphabet([]byte("ABCDEFGHIKLMNOPQRSTUVWXY"), map[byte]byte{}),
	)

	if c != nil || err == nil {
		t.Fatalf("did not fail")
	}
}

func TestPlayfairBasicCrypt(t *testing.T) {
	c, err := NewPlayfair(
		[]byte("CRYPTOGRAPHY"),
	)
	if err != nil {
		t.Fatalf(`unexpected: could not instantiate: %q`, err.Error())
	}

	testCipherCrypt(c, t,
		[]byte("THEQUICKBROWNFOXJUMPSOVERTHELAZYDOGS"),
		[]byte("PBIMXDTDGTAUWNHUDXQRLBRMYCGINOWTLDBM"),
		[]byte("THEQUICKBROWNFOXIUMPSOVERTHELAZYDOGS"),
	)
}

func TestPlayfairKeyWithJCrypt(t *testing.T) {
	c, err := NewPlayfair(
		[]byte("CRYPTOJRAPHY"),
	)
	if err != nil {
		t.Fatalf(`unexpected: could not instantiate: %q`, err.Error())
	}

	testCipherCrypt(c, t,
		[]byte("THEQUICKBROWNFOXJUMPSOVERTHELAZYDOGS"),
		[]byte("PBGMVOTDITAUWNHUOVQRLBRMYCIGNOWTLDKQ"),
		[]byte("THEQUICKBROWNFOXIUMPSOVERTHELAZYDOGS"),
	)
}

func TestPlayfairDoubleLetterCrypt(t *testing.T) {
	c, err := NewPlayfair(
		[]byte("CRYPTOGRAPHY"),
	)
	if err != nil {
		t.Fatalf("unexpected: could not instantiate")
	}

	testCipherCrypt(c, t,
		[]byte("OOF"),
		[]byte("HUAD"),
		[]byte("OXOF"),
	)
}

func TestPlayfairDoubleXCrypt(t *testing.T) {
	c, err := NewPlayfair(
		[]byte("CRYPTOGRAPHY"),
	)
	if err != nil {
		t.Fatalf("unexpected: could not instantiate")
	}

	testCipherCrypt(c, t,
		[]byte("XXF"),
		[]byte("PXWI"),
		[]byte("XQXF"),
	)
}

func TestPlayfairExtraLetterCrypt(t *testing.T) {
	c, err := NewPlayfair(
		[]byte("CRYPTOGRAPHY"),
	)
	if err != nil {
		t.Fatalf("unexpected: could not instantiate")
	}

	testCipherCrypt(c, t,
		[]byte("D"),
		[]byte("IU"),
		[]byte("DX"),
	)
}

func TestPlayfairTrailingSkippedLettersCrypt(t *testing.T) {
	c, err := NewPlayfair(
		[]byte("CRYPTOGRAPHY"),
	)
	if err != nil {
		t.Fatalf("unexpected: could not instantiate")
	}

	testCipherCrypt(c, t,
		[]byte("THEQUICKBROWNFOXJUMPSOVERTHELAZYDOGS?$@."),
		[]byte("PBIMXDTDGTAUWNHUDXQRLBRMYCGINOWTLDBM"),
		[]byte("THEQUICKBROWNFOXIUMPSOVERTHELAZYDOGS"),
	)
}

func BenchmarkPlayfairEncrypt(b *testing.B) {
	c, _ := NewPlayfair(
		[]byte("CRYPTOGRAPHY"),
	)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = c.Encrypt([]byte("INCRYPTOGRAPHYASUBSTITUTIONCIPHERISAMETHODOFENCRYPTINGINWHICHUNITSOFPLAINTEXTAREREPLACEDWITHTHECIPHERTEXTINADEFINEDMANNERWITHTHEHELPOFAKEY"))
	}
}

func BenchmarkPlayfairDecrypt(b *testing.B) {
	c, _ := NewPlayfair(
		[]byte("CRYPTOGRAPHY"),
	)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = c.Decrypt([]byte("FQRYPTCBEGHYAPBNZOZBKPZCDHLYQHGIPENBVMPBDLADFMRYPTPKMAFQXADPOXQFBZADCQHFSYIVYBGMGMCQOYFEXFPBPBDRQHGIYCIVPKWFEFIKMFELFWMFYVKPBPGIGIQCADBFFR"))
	}
}
