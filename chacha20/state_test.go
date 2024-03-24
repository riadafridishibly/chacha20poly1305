package chacha20

import "testing"

func Test_quarterRound(t *testing.T) {
	type args struct {
		a uint32
		b uint32
		c uint32
		d uint32
	}
	tests := []struct {
		name  string
		args  args
		wantA uint32
		wantB uint32
		wantC uint32
		wantD uint32
	}{
		{
			name: "Test Vector 2.1.1 (rfc)",
			args: args{
				a: 0x11111111,
				b: 0x01020304,
				c: 0x9b8d6f43,
				d: 0x01234567,
			},
			wantA: 0xea2a92f4,
			wantB: 0xcb1cf8ce,
			wantC: 0x4581472e,
			wantD: 0x5881c4bb,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2, got3 := quarterRound(tt.args.a, tt.args.b, tt.args.c, tt.args.d)
			if got != tt.wantA {
				t.Errorf("quarterRound() got = %v, want %v", got, tt.wantA)
			}
			if got1 != tt.wantB {
				t.Errorf("quarterRound() got1 = %v, want %v", got1, tt.wantB)
			}
			if got2 != tt.wantC {
				t.Errorf("quarterRound() got2 = %v, want %v", got2, tt.wantC)
			}
			if got3 != tt.wantD {
				t.Errorf("quarterRound() got3 = %v, want %v", got3, tt.wantD)
			}
		})
	}
}
