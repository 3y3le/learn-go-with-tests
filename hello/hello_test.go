package main

import "testing"

// t 是你在测试框架中的 hook（钩子）
func TestHello(t *testing.T) {
	t.Run("say Hello to people", func(t *testing.T) {
		got := Hello("Chris", "")
		expect := "Hello,Chris"

		assertCorrectMessage(t, got, expect)
	})

	t.Run("say hello world when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		expect := "Hello,World!"

		assertCorrectMessage(t, got, expect)
	})
}

func assertCorrectMessage(t *testing.T, got string, expect string) {
	if got != expect {
		t.Errorf("got '%q' expect '%q'", got, expect)
	}
}
