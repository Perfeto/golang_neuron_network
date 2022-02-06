package machine_learn

import "testing"

func Test_sigmoid(t *testing.T) {
	type args struct {
		x float32
	}
	tests := []struct {
		name string
		args args
		want float32
	}{
		{
			name: "want 0",
			args: args{
				x: 0,
			},
			want: 0.5,
		},
		{
			name: "want 0.55",
			args: args{
				x: 0.22,
			},
			want: 0.5547792351072148,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sigmoid(tt.args.x); got != tt.want {
				t.Errorf("sigmoid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_calculateDelta(t *testing.T) {
	type args struct {
		errValue float32
	}
	tests := []struct {
		name string
		args args
		want float32
	}{
		{
			name: "want 0.24",
			args: args{
				errValue: 0.60,
			},
			want: 0.22878424045665732,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sigmoidDerivative(tt.args.errValue); got != tt.want {
				t.Errorf("sigmoidDerivative() = %v, want %v", got, tt.want)
			}
		})
	}
}
