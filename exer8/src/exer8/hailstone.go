package exer8
// TODO: your Hailstone, HailstoneSequenceAppend, HailstoneSequenceAllocate functions

func Hailstone(n uint) uint{
	if (n % 2 == 0){
		return n / 2; 
	}else{
		return (3 * n) + 1
	}
}

// ! BenchmarkHailSeqAppend-6          656902              1797 ns/op
func HailstoneSequenceAppend(n uint) []uint{
	seq := make([]uint,0);
	for (n != 1){
		seq = append (seq, n); 
		n = Hailstone(n); 
	}
	seq = append(seq,1)
	return seq	
}


// ! BenchmarkHailSeqAllocate-6       1493504               803 ns/op
func HailstoneSequenceAllocate(n uint) []uint{
	var l uint = 0; 
	var n1 = n
	for (n1 != 1){
		l += 1
		n1 = Hailstone(n1)
	}

	seq := make ([]uint, l+1)

	var i = 0
	for (i < len(seq)){
		seq[i] = n
		n = Hailstone(n)
		i += 1
	}

	return seq
}



