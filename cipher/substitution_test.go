package cipher

import (
	"testing"
)

func TestSubstitutionNew(t *testing.T) {
	c, err := NewSubstitution(
		[]byte("CRYPTOGRAPHY"),
	)

	if c == nil || err != nil {
		t.Fatalf(`could not instantiate: %q`, err.Error())
	}
}

func TestSubstitutionBasicCrypt(t *testing.T) {
	c, err := NewSubstitution(
		[]byte("CRYPTOGRAPHY"),
	)
	if err != nil {
		t.Fatalf(`unexpected: could not instantiate: %q`, err.Error())
	}

	testCipherCrypt(c, t,
		[]byte("THEQUICKBROWNFOXJUMPSOVERTHELAZYDOG"),
		[]byte("QATLSHYDRMJVIOJWBSFKNJUTMQATECZXPJG"),
		[]byte("THEQUICKBROWNFOXJUMPSOVERTHELAZYDOG"),
	)
}

func TestSubstitutionBasicCryptByte(t *testing.T) {
	c, err := NewSubstitution(
		[]byte("CRYPTOGRAPHY"),
	)
	if err != nil {
		t.Fatalf(`unexpected: could not instantiate: %q`, err.Error())
	}

	testCipherCryptByte(c, t, 'T', 'Q', 'T')
}

func TestSubstitutionNewWithPlainAlphabet(t *testing.T) {
	c, err := NewSubstitution(
		[]byte("CRYPTOGRAPHY"),
		SubstitutionWithPlainAlphabet([]byte("ABCD")),
	)

	if c == nil || err != nil {
		t.Fatalf(`could not instantiate: %q`, err.Error())
	}
}

func TestSubstitutionNewWithNonUniquePlainAlphabet(t *testing.T) {
	c, err := NewSubstitution(
		[]byte("CRYPTOGRAPHY"),
		SubstitutionWithPlainAlphabet([]byte("ABCDB")),
	)

	if c == nil || err != nil {
		t.Fatalf(`could not instantiate: %q`, err.Error())
	}
}

func TestSubstitutionNewWithCipherAlphabet(t *testing.T) {
	c, err := NewSubstitution(
		[]byte("CRYPTOGRAPHY"),
		SubstitutionWithCipherAlphabet([]byte("ZYXWVUTSRQPONMLKJIHGFEDCBA")),
	)

	if c == nil || err != nil {
		t.Fatalf(`could not instantiate: %q`, err.Error())
	}
}

func TestSubstitutionNewWithAlphabets(t *testing.T) {
	c, err := NewSubstitution(
		[]byte("CRYPTOGRAPHY"),
		SubstitutionWithPlainAlphabet([]byte("ABCD")),
		SubstitutionWithCipherAlphabet([]byte("WXYZ")),
	)

	if c == nil || err != nil {
		t.Fatalf(`could not instantiate: %q`, err.Error())
	}
}

func TestSubstitutionNewErrorPlainBigger(t *testing.T) {
	c, err := NewSubstitution(
		[]byte("CRYPTOGRAPHY"),
		SubstitutionWithPlainAlphabet([]byte("ABCD")),
		SubstitutionWithCipherAlphabet([]byte("WXY")),
	)

	if c != nil || err == nil {
		t.Fatalf("did not fail")
	}
}

func TestSubstitutionNewErrorCipherBigger(t *testing.T) {
	c, err := NewSubstitution(
		[]byte("CRYPTOGRAPHY"),
		SubstitutionWithPlainAlphabet([]byte("ABC")),
		SubstitutionWithCipherAlphabet([]byte("WXYZ")),
	)

	if c != nil || err == nil {
		t.Fatalf("did not fail")
	}
}

func TestSubstitutionNewErrorPlainDuplicate(t *testing.T) {
	c, err := NewSubstitution(
		[]byte("CRYPTOGRAPHY"),
		SubstitutionWithPlainAlphabet([]byte("ABBD")),
		SubstitutionWithCipherAlphabet([]byte("WXYZ")),
	)

	if c != nil || err == nil {
		t.Fatalf("did not fail")
	}
}

func TestAtbashBasicCrypt(t *testing.T) {
	c, err := NewAtbash()
	if err != nil {
		t.Fatalf(`unexpected: could not instantiate: %q`, err.Error())
	}

	testCipherCrypt(c, t,
		[]byte("THEQUICKBROWNFOXJUMPSOVERTHELAZYDOG"),
		[]byte("GSVJFRXPYILDMULCQFNKHLEVIGSVOZABWLT"),
		[]byte("THEQUICKBROWNFOXJUMPSOVERTHELAZYDOG"),
	)
}

func TestCaesarBasicCrypt(t *testing.T) {
	c, err := NewCaesar(7)
	if err != nil {
		t.Fatalf(`unexpected: could not instantiate: %q`, err.Error())
	}

	testCipherCrypt(c, t,
		[]byte("THEQUICKBROWNFOXJUMPSOVERTHELAZYDOG"),
		[]byte("AOLXBPJRIYVDUMVEQBTWZVCLYAOLSHGFKVN"),
		[]byte("THEQUICKBROWNFOXJUMPSOVERTHELAZYDOG"),
	)
}

func TestCaesarNewRotationGreaterThanAlphabet(t *testing.T) {
	c, err := NewCaesar(59)
	if err != nil {
		t.Fatalf(`unexpected: could not instantiate: %q`, err.Error())
	}

	testCipherCrypt(c, t,
		[]byte("THEQUICKBROWNFOXJUMPSOVERTHELAZYDOG"),
		[]byte("AOLXBPJRIYVDUMVEQBTWZVCLYAOLSHGFKVN"),
		[]byte("THEQUICKBROWNFOXJUMPSOVERTHELAZYDOG"),
	)
}

func TestCaesarNewRotationLessThanZero(t *testing.T) {
	c, err := NewCaesar(-45)
	if err != nil {
		t.Fatalf(`unexpected: could not instantiate: %q`, err.Error())
	}

	testCipherCrypt(c, t,
		[]byte("THEQUICKBROWNFOXJUMPSOVERTHELAZYDOG"),
		[]byte("AOLXBPJRIYVDUMVEQBTWZVCLYAOLSHGFKVN"),
		[]byte("THEQUICKBROWNFOXJUMPSOVERTHELAZYDOG"),
	)
}

func TestRot13BasicCrypt(t *testing.T) {
	c, err := NewRot13()
	if err != nil {
		t.Fatalf(`unexpected: could not instantiate: %q`, err.Error())
	}

	testCipherCrypt(c, t,
		[]byte("THEQUICKBROWNFOXJUMPSOVERTHELAZYDOG"),
		[]byte("GURDHVPXOEBJASBKWHZCFBIREGURYNMLQBT"),
		[]byte("THEQUICKBROWNFOXJUMPSOVERTHELAZYDOG"),
	)
}

func BenchmarkSubstitutionEncrypt(b *testing.B) {
	c, _ := NewSubstitution([]byte("CRYPTOGRAPHY"))

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = c.Encrypt([]byte("INCRYPTOGRAPHYASUBSTITUTIONCIPHERISAMETHODOFENCRYPTINGINWHICHUNITSOFPLAINTEXTAREREPLACEDWITHTHECIPHERTEXTINADEFINEDMANNERWITHTHEHELPOFAKEY"))
	}
}

func BenchmarkSubstitutionDecrypt(b *testing.B) {
	c, _ := NewSubstitution([]byte("CRYPTOGRAPHY"))

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = c.Decrypt([]byte("HIYMXKQJGMCKAXCNSRNQHQSQHJIYHKATMHNCFTQAJPJOTIYMXKQHIGHIVAHYASIHQNJOKECHIQTWQCMTMTKECYTPVHQAQATYHKATMQTWQHICPTOHITPFCIITMVHQAQATATEKJOCDTX"))
	}
}
