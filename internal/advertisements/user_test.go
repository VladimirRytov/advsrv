package advertisements

import (
	"slices"
	"testing"
)

func TestNewUser(t *testing.T) {
	_, err := NewUser("Петя", "sdasdqweqweqw", false, 1)
	if err != nil {
		t.Fatal(err)
	}
}

func TestSetUserName(t *testing.T) {
	want := "Вася"
	user, err := NewUser("Петя", "sdasdqweqweqw", false, 1)
	if err != nil {
		t.Fatal(err)
	}
	user.SetName(want)
	if user.name != want {
		t.Fatalf("want %s,got %s", want, user.name)
	}
}

func TestSetUserPassword(t *testing.T) {
	want := "1234"
	user, err := NewUser("Петя", "sdasdqweqweqw", false, 1)
	if err != nil {
		t.Fatal(err)
	}
	user.SetPassword(want, true)
	if slices.Equal([]byte(want), user.password) {
		t.Fatalf("want %s,got %s", want, user.name)
	}
}

func TestSetSaltedPassword(t *testing.T) {
	rawPassword := "admin"
	u, err := NewUser("admin", rawPassword, true, int32(1|2))
	salted := u.saltPassword(rawPassword)
	u.SetPassword(string(salted), false)
	if err != nil || !slices.Equal(u.password, salted) {
		t.Fatalf("err %v, passwordEqual %v\n salted: %v\n setted: %v", err, slices.Equal(u.password, []byte(rawPassword)), salted, u.password)
	}
}

func TestSaltPassword(t *testing.T) {
	rawPassword := "sdasdqweqweqw"
	u, err := NewUser("Петя", "sdasdqweqweqw", true, int32(1|2))
	u.SetPassword(rawPassword, true)
	if err != nil || slices.Equal(u.password, []byte(rawPassword)) {
		t.Fatalf("err %v, passwordEqual %v\n salted: %v\n raw: %v", err, slices.Equal(u.password, []byte(rawPassword)), u.password, []byte(rawPassword))
	}
}

func BenchmarkSaltPassword(b *testing.B) {
	b.StopTimer()
	user, err := NewUser("Петя", "sdasdqweqweqw", false, int32(1|2))
	if err != nil {
		b.Fatal(err)
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		user.saltPassword("sdasdqweqweqw")
	}
}
