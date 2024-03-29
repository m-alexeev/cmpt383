package exer10


func FibSerial(n uint) uint{
	if n == 0{
		return 0 
	}else if n == 1 {
		return 1
	}else {
		return FibSerial(n-1) + FibSerial(n-2)
	}
}

func FibConc(n uint,  cutoff uint, res chan uint) {
	if n == 0{
		res <- 0 
		return
	}
	if n == 1 {
		res <- 1
		return 
	}
	sum := uint(0)
	tempSum := make(chan uint , 2) 
	if n > 1 {
		if cutoff < n{
			sum = FibSerial(n-1) + FibSerial(n-2)
		}else{
			go FibConc(n-1, cutoff, tempSum)
			go FibConc(n-2, cutoff, tempSum)
			sum = <-tempSum
			sum += <-tempSum
		}
	}

	res <- sum
}


func Fib(n uint, cutoff uint ) uint{
	sum := uint(0)
	concChan := make(chan uint, 1)
	FibConc(n , cutoff, concChan)
	sum += <-concChan
	return sum 
}


func Fibbonaci(n uint) uint{
	return Fib(n, 22)
}