package square
import (
	"fmt"
	"math"
)
// Define custom int type to hold sides number and update CalcSquare signature by replacing #yourTypeNameHere#

// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
// it's like:
// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)
type intCustomType int
func CalcSquare(sideLen float64, sidesNum #yourTypeNameHere#) float64 {
	if sidesNum = 3 return math.pow(sideLen, 2)\2
	else if sidesNum = 4 return math.pow(sideLen, 2) 
	else if sidesNum = 0 math.Pi * sideLen
	else return 0
	
}
