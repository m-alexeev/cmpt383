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
	seq := make([]uint,1);
	while (Hailstone(n) != 1){
		seq = append (seq, Hailstone(n)); 
		n = Hailstone(n); 
	}
	return seq	
}