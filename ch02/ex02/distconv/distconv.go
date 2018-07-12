package distconv

import "fmt"

type Meter float64
type Mile float64

func (meter Meter) String() string {
	return fmt.Sprintf("%gm", meter)
}

func (mile Mile) String() string {
	return fmt.Sprintf("%gmi", mile)
}
