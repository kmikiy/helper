package ptr

import (
	"testing"
)

func TestBool(t *testing.T) {
	type args struct {
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
				value: true,
			},
			want: true,
		},
		{
			name: "False",
			args: args{
				value: false,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Bool(tt.args.value); *got != tt.want {
				t.Errorf("Bool() = %v, want %v", *got, tt.want)
			}
		})
	}
}

func TestString(t *testing.T) {
	type args struct {
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
				value: "random string",
			},
			want: "random string",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := String(tt.args.value); *got != tt.want {
				t.Errorf("String() = %v, want %v", *got, tt.want)
			}
		})
	}
}

func TestInt8(t *testing.T) {
	type args struct {
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
				value: 8,
			},
			want: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Int8(tt.args.value); *got != tt.want {
				t.Errorf("Int8() = %v, want %v", *got, tt.want)
			}
		})
	}
}

func TestUInt8(t *testing.T) {
	type args struct {
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
				value: 8,
			},
			want: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UInt8(tt.args.value); *got != tt.want {
				t.Errorf("UInt8() = %v, want %v", *got, tt.want)
			}
		})
	}
}

func TestByte(t *testing.T) {
	type args struct {
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
				value: 's',
			},
			want: 's',
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Byte(tt.args.value); *got != tt.want {
				t.Errorf("Byte() = %v, want %v", *got, tt.want)
			}
		})
	}
}

func TestInt16(t *testing.T) {
	type args struct {
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
				value: 16,
			},
			want: 16,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Int16(tt.args.value); *got != tt.want {
				t.Errorf("Int16() = %v, want %v", *got, tt.want)
			}
		})
	}
}

func TestUInt16(t *testing.T) {
	type args struct {
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
				value: 16,
			},
			want: 16,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UInt16(tt.args.value); *got != tt.want {
				t.Errorf("UInt16() = %v, want %v", *got, tt.want)
			}
		})
	}
}

func TestInt32(t *testing.T) {
	type args struct {
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
				value: 32,
			},
			want: 32,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Int32(tt.args.value); *got != tt.want {
				t.Errorf("Int32() = %v, want %v", *got, tt.want)
			}
		})
	}
}

func TestRune(t *testing.T) {
	type args struct {
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
				value: 'j',
			},
			want: 'j',
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Rune(tt.args.value); *got != tt.want {
				t.Errorf("Rune() = %v, want %v", *got, tt.want)
			}
		})
	}
}

func TestUInt32(t *testing.T) {
	type args struct {
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
				value: 32,
			},
			want: 32,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UInt32(tt.args.value); *got != tt.want {
				t.Errorf("UInt32() = %v, want %v", *got, tt.want)
			}
		})
	}
}

func TestInt64(t *testing.T) {
	type args struct {
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
				value: 64,
			},
			want: 64,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Int64(tt.args.value); *got != tt.want {
				t.Errorf("Int64() = %v, want %v", *got, tt.want)
			}
		})
	}
}

func TestUInt64(t *testing.T) {
	type args struct {
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
				value: 64,
			},
			want: 64,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UInt64(tt.args.value); *got != tt.want {
				t.Errorf("UInt64() = %v, want %v", *got, tt.want)
			}
		})
	}
}

func TestInt(t *testing.T) {
	type args struct {
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
				value: 8,
			},
			want: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Int(tt.args.value); *got != tt.want {
				t.Errorf("Int() = %v, want %v", *got, tt.want)
			}
		})
	}
}

func TestUInt(t *testing.T) {
	type args struct {
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
				value: 8,
			},
			want: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UInt(tt.args.value); *got != tt.want {
				t.Errorf("UInt() = %v, want %v", *got, tt.want)
			}
		})
	}
}

func TestUIntPtr(t *testing.T) {
	type args struct {
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
				value: 8,
			},
			want: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UIntPtr(tt.args.value); *got != tt.want {
				t.Errorf("UIntPtr() = %v, want %v", *got, tt.want)
			}
		})
	}
}

func TestFloat32(t *testing.T) {
	type args struct {
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
				value: 32.0,
			},
			want: 32.0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Float32(tt.args.value); *got != tt.want {
				t.Errorf("Float32() = %v, want %v", *got, tt.want)
			}
		})
	}
}

func TestFloat64(t *testing.T) {
	type args struct {
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
				value: 64.0,
			},
			want: 64.0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Float64(tt.args.value); *got != tt.want {
				t.Errorf("Float64() = %v, want %v", *got, tt.want)
			}
		})
	}
}

func TestComplex64(t *testing.T) {
	type args struct {
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
				value: 64 + 0i,
			},
			want: 64 + 0i,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Complex64(tt.args.value); *got != tt.want {
				t.Errorf("Complex64() = %v, want %v", *got, tt.want)
			}
		})
	}
}

func TestComplex128(t *testing.T) {
	type args struct {
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
				value: 128 + 0i,
			},
			want: 128 + 0i,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Complex128(tt.args.value); *got != tt.want {
				t.Errorf("Complex128() = %v, want %v", *got, tt.want)
			}
		})
	}
}
