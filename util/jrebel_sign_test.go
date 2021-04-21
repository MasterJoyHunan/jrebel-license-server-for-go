package util

import "testing"

func TestJrebelSign(t *testing.T) {
	type args struct {
		clientRandomness string
		guid             string
		validFrom        string
		validUntil       string
		offline          string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"", args{
			clientRandomness: "1",
			guid:             "2",
			validFrom:        "",
			validUntil:       "",
			offline:          "false",
		}, "IOlHi1qQx4CmpMcxIv+VEjMu/F1kpLaRmO2ONQEEQOA/dxWIyNBxhqUWFzasqzsuKtBpqQP2qShf70hNEg7zuOXIchP5U6BnXrE3FFAeVEnmFWYFQHlzqToyp2qGiP1uMujl5p0sTd92jBXsdlhzVrgwdCRWDS/fYXCzlNGG51I="},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := JrebelSign(tt.args.clientRandomness, tt.args.guid, tt.args.validFrom, tt.args.validUntil, tt.args.offline); got != tt.want {
				t.Errorf("JrebelSign() = %v, want %v", got, tt.want)
			}
		})
	}
}




func TestJrebelRsaSign(t *testing.T) {
	type args struct {
		b []byte
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"", args{[]byte("<PingResponse><message></message><responseCode>OK</responseCode><salt>xxx</salt></PingResponse>")},
			"831a8f75189d1454442548320e56e9ce58588b99b6a4842fb5f7f9e0f90b37db65364ba0e656c532656b9c024280e548e6718e38216623854beeba3a492492ef"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := JrebelRsaSign(tt.args.b); got != tt.want {
				t.Errorf("JrebelRsaSign() = %v, want %v", got, tt.want)
			}
		})
	}
}
