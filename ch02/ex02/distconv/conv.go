package distconv

func MToMi(meter Meter) Mile {
	return Mile(meter * 0.0006)
}

func MiToM(mile Mile) Meter {
	return Meter(mile * 1600)
}
