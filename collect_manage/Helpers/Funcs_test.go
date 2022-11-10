package Helpers

import "testing"

func init() {

}

func TestUrlDecode(t *testing.T) {
	var f Funcs
	f = Funcs{}

	s := f.UrlDecode("https%3A%2F%2Fgithub.com%2FCompVis%2Fstable-diffusion")

	t.Logf("done: %s", s)
}
