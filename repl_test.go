package main

import (
	"fmt"
	"testing"
	"time"

	"github.com/felixsolom/pokedexcli/internal/pokecache"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "   hello, world  ",
			expected: []string{"hello,", "world"},
		},
		{
			input:    "   Dream a little Dream !",
			expected: []string{"dream", "a", "little", "dream", "!"},
		},
		{
			input:    " He was more   Than surprised    ",
			expected: []string{"he", "was", "more", "than", "surprised"},
		},
	}
	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("slice length: %v is not equal to expeted length %v",
				len(actual), len(c.expected))
		}
		for i := range actual {
			word := actual[i]
			expetedWord := c.expected[i]
			if word != expetedWord {
				t.Errorf("Word: %v doesn't match expected word %v", word, expetedWord)
			}
		}
	}
}

func TestAddGet(t *testing.T) {
	cache := pokecache.NewCache(50 * time.Millisecond)
	cases := []struct {
		key string
		val []byte
	}{
		{
			key: "https://example.com",
			val: []byte("testdata"),
		},
		{
			key: "https://example.com/path",
			val: []byte("moretestdata"),
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("Test case %v", i), func(t *testing.T) {
			cache.Add(c.key, c.val)
			val, ok := cache.Get(c.key)
			if !ok {
				t.Errorf("expected to find key")
				return
			}
			if string(val) != string(c.val) {
				t.Errorf("expected to find value")
				return
			}
		})
	}
}

func TestReapLoop(t *testing.T) {
	cache := pokecache.NewCache(50 * time.Millisecond)
	cache.Add("https://example.com", []byte("testdata"))

	_, ok := cache.Get("https://example.com")
	if !ok {
		t.Errorf("expected to find key")
		return
	}

	time.Sleep(100 * time.Millisecond)

	_, ok = cache.Get("https://example.com")
	if ok {
		t.Errorf("expected to not find key")
		return
	}
}
