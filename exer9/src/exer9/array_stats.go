package exer9

import(
	"math/rand"
	"time"	
	"math"
)

func RandomArray(length int, maxInt int) []int {
	// TODO: create a new random generator with a decent seed; create an array with length values from 0 to values-1.
	rand.Seed(time.Now().UnixNano() + int64(rand.Intn(100)))

	a := make([]int, length)
	for i := 0; i < length; i++{
		a[i] = rand.Intn(maxInt)
	}
	return a
}


type PartSum struct{
	x int
	x_sq int
}


func Sums(arr []int,sums chan PartSum){
	x := 0
	x_sq := 0
	for i:=0; i < len(arr); i++{
		x += arr[i]
		x_sq += arr[i]*arr[i]
	}
	sums <- PartSum{x,x_sq}
}


func MeanStddev(arr []int, chunks int) (mean, stddev float64) {
	if len(arr)%chunks != 0 {
		panic("You promised that chunks would divide slice size")
	}
	numItems := len(arr) / chunks
	// TODO: calculate the mean and population standard deviation of the array, breaking the array into chunks segments
	// and calculating on them in parallel.
	partSum := PartSum{0,0}

	ch := make(chan PartSum)
	for i:=0 ; i < chunks; i++{
		go Sums(arr[i * numItems: (i+1) * numItems],ch)
	}

	totalSum := PartSum{0,0}
	for j :=0; j < chunks; j++{
		partSum = <-ch
		totalSum.x += partSum.x
		totalSum.x_sq += partSum.x_sq
	}

	var x float64 = float64(totalSum.x)
	var x_sq float64 = float64(totalSum.x_sq)
	var n float64 = float64(len(arr))

	mean = x / n
	stddev = math.Sqrt((x_sq/ n) - math.Pow(mean,2))	


	return mean, stddev

}
