package testing

import (
	"fmt"
	"testing"

	"github.com/Homebrew-Blockchain-Club/minichain/typeconv"
)

func TestTypeConvInt(t *testing.T) {
	x := 5
	y := typeconv.ToBytes(5)
	x = typeconv.FromBytes[int](y)
	println(x)
}
func TestTypeConvStruct(t *testing.T) {
	type tmp struct {
		A, B int
	}
	x := tmp{
		A: 1, B: 2,
	}
	y := typeconv.ToBytes(x)
	x = typeconv.FromBytes[tmp](y)
	fmt.Printf("%v", x)
}
func TestTypeConvStructPointer(t *testing.T) {
	type tmp struct {
		A *int
	}
	b := 5
	a := tmp{
		A: &b,
	}
	c := typeconv.ToBytes(a)
	a = typeconv.FromBytes[tmp](c)
	fmt.Printf("%v", *a.A)
}

func TestTypeConvSlice(t *testing.T) {
	b := typeconv.ToBytes("114514")
	a := typeconv.FromBytes[string](b)
	println(a)
}

func TestInterface(t *testing.T) {
	type y struct {
		A int
	}
	a := y{
		A: 1,
	}
	b := typeconv.ToBytes(a)
	a = typeconv.FromBytes[any](b).(y)
	println(a.A)
}

func TestIsomorphismStruct(t *testing.T) {
	type y struct {
		A int
	}
	b := typeconv.ToBytes(struct {
		A int
	}{
		A: 5,
	})
	a := typeconv.FromBytes[y](b)
	fmt.Println(a)
}
