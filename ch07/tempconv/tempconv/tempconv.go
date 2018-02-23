package tempconv

import (
	"flag"
	"fmt"
)

type Celsius float64
type Fahrenheit float64
type Kelvin float64

// CToF は摂氏の温度を華氏へ変換します
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

// FtoC は華氏の温度を摂氏へ変換します
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

func CToK(c Celsius) Kelvin { return Kelvin(c + 273.15) }

func KToC(k Kelvin) Celsius { return Celsius(k - 273.15) }

func KToF(k Kelvin) Fahrenheit { return CToF(KToC(k)) }

func FToK(f Fahrenheit) Kelvin { return CToK(FToC(f)) }

func (c Celsius) String() string    { return fmt.Sprintf("%g°C", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%g°F", f) }
func (k Kelvin) String() string     { return fmt.Sprintf("%g°K", k) }

type celsiusFlag struct{ Celsius }

func (f *celsiusFlag) Set(s string) error {
	var uint string
	var value float64
	fmt.Sscanf(s, "%f%s", &value, &uint)
	switch uint {
	case "C", "°C":
		f.Celsius = Celsius(value)
		return nil
	case "F", "°F":
		f.Celsius = FToC(Fahrenheit(value))
		return nil
	}
	return fmt.Errorf("invalid tempeature %q", s)
}

// CelsiusFlangは指定された名前、デフォルト値,使い方を持つCelsiusフラグ
// を定義しており、そのフラグ変数のアドレスを返します。
// フラグ引数は度数と単位です。例えば"100C"です。
func CelsiusFlag(name string, value Celsius, usage string) *Celsius {
	f := celsiusFlag{value}
	flag.CommandLine.Var(&f, name, usage)
	return &f.Celsius
}
