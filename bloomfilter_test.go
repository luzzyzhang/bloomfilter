package bloomfilter

import (
	"testing"
)

func TestBloomFilter(t *testing.T) {
	testCases := []struct {
		desc  string
		input string
		want  bool
	}{
		{
			desc:  "Basic feature",
			input: "https://www.google.com",
			want:  true,
		},
		{
			desc:  "Basic feature",
			input: "https://www.baidu.com",
			want:  true,
		},
		{
			desc:  "Basic feature",
			input: "https://www.456.com",
			want:  false,
		},
		{
			desc:  "Basic feature",
			input: "a",
			want:  true,
		},
		{
			desc:  "Basic feature",
			input: "b",
			want:  true,
		},
		{
			desc:  "Basic feature",
			input: "123",
			want:  true,
		},
		{
			desc:  "Basic feature",
			input: "a+",
			want:  false,
		},
		{
			desc:  "Basic feature",
			input: "[a]",
			want:  false,
		},
		// {
		// 	desc: "Special input",
		// },
		// {
		// 	desc: "Edege cases",
		// },
	}

	bf := New(512)
	lst := []string{
		"https://www.baidu.com",
		"https://www.google.com",
		"https://www.123.com",
		"a", "b", "c", "123",
	}
	for _, c := range lst {
		bf.Add([]byte(c))
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got := bf.Contain([]byte(tC.input))
			if got != tC.want {
				t.Errorf("bloomfilter.query(%q) == %t, want %t", tC.input, got, tC.want)
			}
		})
	}
}
