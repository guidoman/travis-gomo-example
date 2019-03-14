package ciexample

import "testing"

func TestHello(t *testing.T) {
	const hello = "Hello, world"
	if s := Hello(); s != hello {
		t.Errorf("Hello() = %v, want %v", s, hello)
	}
}
