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
		// {
		// 	desc: "Special input",
		// },
		// {
		// 	desc: "Edege cases",
		// },
	}
	bf := bloomfilter{}
	lst := []string{"https://www.baidu.com", "https://www.google.com", "https://www.123.com"}
	for _, c := range lst {
		bf.add(c)
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got := bf.query(tC.input)
			// fmt.Printf("input: %v, bitarray: %v\n", tC.input, bf.bitarray)
			if got != tC.want {
				t.Errorf("bloomfilter.query(%q) == %t, want %t", tC.input, got, tC.want)
			}
		})
	}
}
