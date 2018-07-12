package tempconv

// 摂氏の温度を華氏へ変換します
func CToF(c Celsius) Fahrenherit {
	return Fahrenherit(c*9/5 + 32)
}

// 華氏の温度を摂氏へ変換します
func FToC(f Fahrenherit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

func CToK(c Celsius) Kelvin {
	return Kelvin(c + 273.15)
}

func KToC(k Kelvin) Celsius {
	return Celsius(k - 273.15)
}
