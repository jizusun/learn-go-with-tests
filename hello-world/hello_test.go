package helloworld

import (
	"testing"
	// . "github.com/smartystreets/goconvey/convey"
)

// func TestHello(t *testing.T) {
// 	Convey("1 should equal 1", t, func() {
// 		So(Hello("Chris"), ShouldEqual, "Hello, Chris")
// 	})
// }

func TestHello(t *testing.T) {
	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("say hello to people", func(t *testing.T) {
		got := Hello("Chris", "English")
		want := "Hello, Chris"
		assertCorrectMessage(t, got, want)
	})

	t.Run("empty string defaults to 'World'", func(t *testing.T) {
		got := Hello("", "English")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"
		assertCorrectMessage(t, got, want)
	})
}

func TestHelloWithTable(t *testing.T) {
	type args struct {
		name     string
		language string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"empty string", args{"", "English"}, "Hello, World"},
		{"non-empty string", args{"Jizu", "English"}, "Hello, Jizu"},

		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Hello(tt.args.name, tt.args.language); got != tt.want {
				t.Errorf("🛑 Hello() = %v, want %v", got, tt.want)
			}
		})
	}
}
