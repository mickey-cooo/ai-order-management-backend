package services

import "testing"

func TestComparePassword(t *testing.T) {
	h, err := HashPassword("s3cret")
	if err != nil {
		t.Fatal(err)
	}
	if !ComparePassword(h, "s3cret") {
		t.Fatal("correct password rejected")
	}
	if ComparePassword(h, "wrong") {
		t.Fatal("wrong password accepted")
	}
	if ComparePassword("garbage", "s3cret") {
		t.Fatal("malformed hash accepted")
	}
}

func TestHashAndVerifyPassword(t *testing.T) {
	h, err := HashPassword("s3cret")
	if err != nil {
		t.Fatal(err)
	}
	if ok, err := VerifyPassword("s3cret", h); err != nil || !ok {
		t.Fatalf("correct password rejected: ok=%v err=%v", ok, err)
	}
	if ok, _ := VerifyPassword("wrong", h); ok {
		t.Fatal("wrong password accepted")
	}
	h2, _ := HashPassword("s3cret")
	if h == h2 {
		t.Fatal("salt not random")
	}
	for _, bad := range []string{
		"garbage",
		"$argon2id$v=19$m=65536,t=0,p=4$c2FsdA$a2V5",
		"$argon2id$v=19$m=65536,t=1,p=0$c2FsdA$a2V5",
		"$argon2id$v=19$m=4294967295,t=1,p=4$c2FsdA$a2V5",
		"$argon2id$v=19$m=65536,t=4294967295,p=4$c2FsdA$a2V5",
	} {
		if _, err := VerifyPassword("x", bad); err != ErrInvalidHash {
			t.Fatalf("%q: want ErrInvalidHash, got %v", bad, err)
		}
	}
}

func TestGeneratePassword(t *testing.T) {
	for _, n := range []int{-1, 0, minGeneratedLen - 1} {
		if _, err := GeneratePassword(n); err != ErrPasswordShort {
			t.Fatalf("length %d: want ErrPasswordShort, got %v", n, err)
		}
	}
	p, err := GeneratePassword(minGeneratedLen)
	if err != nil || len(p) != minGeneratedLen {
		t.Fatalf("got %q err=%v", p, err)
	}
}
