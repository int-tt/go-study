package main

import (
	"bytes"
	"fmt"
	"reflect"
	"strconv"

	"github.com/gopl/ch07/eval"
)

type Loop struct {
	count int
}

func Display(name string, x interface{}) {
	fmt.Printf("Display %s (%T): \n", name, x)
	l := &Loop{
		count: 0,
	}
	l.display(name, reflect.ValueOf(x))
}

func (l *Loop) display(path string, v reflect.Value) {
	l.count++
	if l.count > 25 {
		return
	}

	switch v.Kind() {
	case reflect.Invalid:
		fmt.Printf("%s = invalid\n", path)
	case reflect.Slice, reflect.Array:
		for i := 0; i < v.Len(); i++ {
			l.display(fmt.Sprintf("%s[%d]", path, i), v.Index(i))
		}
	case reflect.Struct:
		for i := 0; i < v.NumField(); i++ {
			fieldPath := fmt.Sprintf("%s.%s", path, v.Type().Field(i).Name)
			l.display(fieldPath, v.Field(i))
		}
	case reflect.Map:
		for _, key := range v.MapKeys() {
			l.display(fmt.Sprintf("%s[%s]", path, formatAtom(key)), v.MapIndex(key))
		}
	case reflect.Ptr:
		if v.IsNil() {
			fmt.Printf("%s = nil\n", path)
		} else {
			l.display(fmt.Sprintf("(*%s)", path), v.Elem())
		}
	case reflect.Interface:
		if v.IsNil() {
			fmt.Printf("%s.type = nil\n", path)
		} else {
			fmt.Printf("%s.type = %s\n", path, v.Elem().Type())
			l.display(path+".value", v.Elem())
		}
	default:
		fmt.Printf("%s = %s\n", path, formatAtom(v))
	}
}

func formatAtom(v reflect.Value) string {
	switch v.Kind() {
	case reflect.Invalid:
		return "invalid"
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return strconv.FormatInt(v.Int(), 10)
	case reflect.Uint, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return strconv.FormatUint(v.Uint(), 10)
	case reflect.Bool:
		return strconv.FormatBool(v.Bool())
	case reflect.String:
		return strconv.Quote(v.String())
	case reflect.Chan, reflect.Func, reflect.Ptr, reflect.Slice, reflect.Map:
		return v.Type().String() + "0x" + strconv.FormatUint(uint64(v.Pointer()), 16)
	case reflect.Struct:
		var buf bytes.Buffer

		buf.WriteString(v.Type().String())
		buf.WriteRune('{')
		for i := 0; i < v.NumField(); i++ {
			buf.WriteString(fmt.Sprintf("%s: %s", v.Type().Field(i).Name, formatAtom(v.Field(i))))
			if i < (v.NumField() - 1) {
				buf.WriteString(", ")
			}
		}
		buf.WriteRune('}')
		return buf.String()
	case reflect.Array:
		var buf bytes.Buffer

		buf.WriteString(v.Type().String())
		buf.WriteRune('[')
		for i := 0; i < v.Len(); i++ {
			buf.WriteString(formatAtom(v.Index(i)))
			if i < (v.Len() - 1) {
				buf.WriteString(", ")
			}
		}
		buf.WriteRune(']')
		return buf.String()
	default:
		return v.Type().String() + " value"
	}
}

type Movie struct {
	Title, Subtitle string
	Year            int
	Color           bool
	Actor           map[string]string
	Oscars          []string
	Sequel          *string
}

func main() {
	e, _ := eval.Parse("sqrt(A / pi)")
	Display("e", e)

	strangelove := Movie{
		Title:    "Dr. Strangelove",
		Subtitle: "How I Learned to Stop Worring and Love the Bomb",
		Year:     1964,
		Color:    false,
		Actor: map[string]string{
			"Dr. Strangelove":            "Peter Sellers",
			"Grp, Capt. Lionel Mandrake": "Peter Sellers",
			"Gen. Buck Turgidson":        "George C. Scott",
			"Brig. Gen. Jack D. Ripper":  "Sterling hayden",
			`Maj. T.J. "King" Kong`:      "Slim Pickens",
		},
		Oscars: []string{
			"Best Actor (Nomin.)",
			"Best Adapted Screenplay (Nomin.)",
			"Best Director (Nomin.)",
			"Best Piture (Nomin.)",
		},
	}
	Display("strangelove", strangelove)

	name := [2]string{"hello", "world"}
	salaries := map[[2]string]int{
		name: 1701,
	}
	Display("salaries", salaries)
}
