// tempconv パッケージは摂氏(Celsius)と華氏(Fahrenheit)の温度変換を行います
package tempconv

import "fmt"

type Celsius float64
type Fahrenheit float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

func (c Celsius) String() string    { return fmt.Sprintf("%g °C", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%g °F", f) }
