package tempconv

// CToF は摂氏の温度を華氏へ変換します
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

// FtoC は華氏の温度を摂氏へ変換します
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

// CToK は摂氏の温度を絶対温度へ変換します
func CToK(c Celsius) Kelvin { return Kelvin(c + 273.15) }

// KToC は絶対温度を摂氏へ変換します
func KToC(c Kelvin) Celsius { return Celsius(c - 273.15) }

// FToK は華氏の温度を絶対温度へ変換します
func FToK(f Fahrenheit) Kelvin { return Kelvin(f) }

// CTok は絶対温度を華氏へ変換します
func KToF(k Kelvin) Fahrenheit { return Fahrenheit(k) }
