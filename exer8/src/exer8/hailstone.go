package exer8

// TODO: your Hailstone, HailstoneSequenceAppend, HailstoneSequenceAllocate functions

func Hailstone(n uint) uint{
	if (n % 2 == 0){
		return n / 2; 
	}else{
		return (3 * n) + 1
	}
}



func HailstoneSequenceAppend(n uint) []uint{
	seq := make([]uint,0);
	seq = append(seq, n)
	for (Hailstone(n) != 1){
		seq = append (seq, Hailstone(n)); 
		n = Hailstone(n); 
	}
	seq = append(seq, 1)
	return seq	
}

func HailstoneSequenceAllocate(n uint) []uint{
	seq := make([]uint, 0)
	return seq
}



