package rgb

import (
	"math"
)
func GenerateRGB(i int) (int,int,int){
	f:=0.1
	return int(math.Sin(f*float64(i)+0)*127 + 128),
        int(math.Sin(f*float64(i)+2*math.Pi/3)*127 + 128),
        int(math.Sin(f*float64(i)+4*math.Pi/3)*127 + 128)
}
