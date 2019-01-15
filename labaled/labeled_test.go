package labeled

import (
	"testing"
	"time"
)

func TestBool(t *testing.T) {
	type args struct {
		label string
		value bool
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "True",
			args: args{
				label: "True",
				value: true,
			},
			want: true,
		},
		{
			name: "False",
			args: args{
				label: "False",
				value: false,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Bool(tt.args.label, tt.args.value); got != tt.want {
				t.Errorf("Bool() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestString(t *testing.T) {
	type args struct {
		label string
		value string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "String",
			args: args{
				label: "Random String",
				value: "random string",
			},
			want: "random string",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := String(tt.args.label, tt.args.value); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt8(t *testing.T) {
	type args struct {
		label string
		value int8
	}
	tests := []struct {
		name string
		args args
		want int8
	}{
		{
			name: "Int8",
			args: args{
				label: "Random Int8",
				value: 8,
			},
			want: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Int8(tt.args.label, tt.args.value); got != tt.want {
				t.Errorf("Int8() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUInt8(t *testing.T) {
	type args struct {
		label string
		value uint8
	}
	tests := []struct {
		name string
		args args
		want uint8
	}{
		{
			name: "UInt8",
			args: args{
				label: "Random UInt8",
				value: 8,
			},
			want: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UInt8(tt.args.label, tt.args.value); got != tt.want {
				t.Errorf("UInt8() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestByte(t *testing.T) {
	type args struct {
		label string
		value byte
	}
	tests := []struct {
		name string
		args args
		want byte
	}{
		{
			name: "Byte",
			args: args{
				label: "s Byte",
				value: 's',
			},
			want: 's',
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Byte(tt.args.label, tt.args.value); got != tt.want {
				t.Errorf("Byte() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt16(t *testing.T) {
	type args struct {
		label string
		value int16
	}
	tests := []struct {
		name string
		args args
		want int16
	}{
		{
			name: "Int16",
			args: args{
				label: "Random Int16",
				value: 16,
			},
			want: 16,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Int16(tt.args.label, tt.args.value); got != tt.want {
				t.Errorf("Int16() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUInt16(t *testing.T) {
	type args struct {
		label string
		value uint16
	}
	tests := []struct {
		name string
		args args
		want uint16
	}{
		{
			name: "UInt16",
			args: args{
				label: "Random UInt16",
				value: 16,
			},
			want: 16,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UInt16(tt.args.label, tt.args.value); got != tt.want {
				t.Errorf("UInt16() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt32(t *testing.T) {
	type args struct {
		label string
		value int32
	}
	tests := []struct {
		name string
		args args
		want int32
	}{
		{
			name: "Int32",
			args: args{
				label: "Random Int32",
				value: 32,
			},
			want: 32,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Int32(tt.args.label, tt.args.value); got != tt.want {
				t.Errorf("Int32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRune(t *testing.T) {
	type args struct {
		label string
		value rune
	}
	tests := []struct {
		name string
		args args
		want rune
	}{
		{
			name: "Rune",
			args: args{
				label: "j Rune",
				value: 'j',
			},
			want: 'j',
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Rune(tt.args.label, tt.args.value); got != tt.want {
				t.Errorf("Rune() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUInt32(t *testing.T) {
	type args struct {
		label string
		value uint32
	}
	tests := []struct {
		name string
		args args
		want uint32
	}{
		{
			name: "UInt32",
			args: args{
				label: "Random UInt32",
				value: 32,
			},
			want: 32,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UInt32(tt.args.label, tt.args.value); got != tt.want {
				t.Errorf("UInt32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt64(t *testing.T) {
	type args struct {
		label string
		value int64
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "Int64",
			args: args{
				label: "Random Int64",
				value: 64,
			},
			want: 64,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Int64(tt.args.label, tt.args.value); got != tt.want {
				t.Errorf("Int64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUInt64(t *testing.T) {
	type args struct {
		label string
		value uint64
	}
	tests := []struct {
		name string
		args args
		want uint64
	}{
		{
			name: "UInt64",
			args: args{
				label: "Random UInt64",
				value: 64,
			},
			want: 64,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UInt64(tt.args.label, tt.args.value); got != tt.want {
				t.Errorf("UInt64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt(t *testing.T) {
	type args struct {
		label string
		value int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Int",
			args: args{
				label: "Random Int",
				value: 8,
			},
			want: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Int(tt.args.label, tt.args.value); got != tt.want {
				t.Errorf("Int() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUInt(t *testing.T) {
	type args struct {
		label string
		value uint
	}
	tests := []struct {
		name string
		args args
		want uint
	}{
		{
			name: "UInt",
			args: args{
				label: "Random UInt",
				value: 8,
			},
			want: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UInt(tt.args.label, tt.args.value); got != tt.want {
				t.Errorf("UInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUIntPtr(t *testing.T) {
	type args struct {
		label string
		value uintptr
	}
	tests := []struct {
		name string
		args args
		want uintptr
	}{
		{
			name: "UIntPtr",
			args: args{
				label: "Random UIntPtr",
				value: 8,
			},
			want: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UIntPtr(tt.args.label, tt.args.value); got != tt.want {
				t.Errorf("UIntPtr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFloat32(t *testing.T) {
	type args struct {
		label string
		value float32
	}
	tests := []struct {
		name string
		args args
		want float32
	}{
		{
			name: "Float32",
			args: args{
				label: "Random Float32",
				value: 32.0,
			},
			want: 32.0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Float32(tt.args.label, tt.args.value); got != tt.want {
				t.Errorf("Float32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFloat64(t *testing.T) {
	type args struct {
		label string
		value float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "Float64",
			args: args{
				label: "Random Float64",
				value: 64.0,
			},
			want: 64.0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Float64(tt.args.label, tt.args.value); got != tt.want {
				t.Errorf("Float64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestComplex64(t *testing.T) {
	type args struct {
		label string
		value complex64
	}
	tests := []struct {
		name string
		args args
		want complex64
	}{
		{
			name: "Complex64",
			args: args{
				label: "Random Complex64",
				value: 64 + 0i,
			},
			want: 64 + 0i,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Complex64(tt.args.label, tt.args.value); got != tt.want {
				t.Errorf("Complex64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestComplex128(t *testing.T) {
	type args struct {
		label string
		value complex128
	}
	tests := []struct {
		name string
		args args
		want complex128
	}{
		{
			name: "Complex128",
			args: args{
				label: "Random Complex128",
				value: 128 + 0i,
			},
			want: 128 + 0i,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Complex128(tt.args.label, tt.args.value); got != tt.want {
				t.Errorf("Complex128() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTime(t *testing.T) {
	now := time.Now()
	type args struct {
		label string
		value time.Time
	}
	tests := []struct {
		name string
		args args
		want time.Time
	}{
		{
			name: "Time",
			args: args{
				label: "Now",
				value: now,
			},
			want: now,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Time(tt.args.label, tt.args.value); got != tt.want {
				t.Errorf("Time() = %v, want %v", got, tt.want)
			}
		})
	}
}
