divisors :: Int -> [Int]
divisors n = [i | i <- [2..(n `div` 2)], n `mod` i == 0] -- range (2 .. n/2), keeps all values n % i == 0
primes :: Int -> [Int]
primes n = [i | i <- [1..n], divisors i == []]

pythagorean :: Int -> [(Int, Int, Int)]
pythagorean n = [(a,b,c) | c <- [1..n], a<-[1..n],b<-[a..n] , a*a + b*b == c*c]

join x [] = []
join x [a] = a
join x (a : rest) = a++x ++ join x rest

fac' :: Int -> Int
fac' n = foldl (*) 1 [1..n]


hailstone x 
    |mod x 2 == 0 = div x 2
    |mod x 2 >  0 = 3 * x + 1


hailLen n = hailTail 0 n
  where
    hailTail a 1 = a
    hailTail a n = hailTail (a + 1) (hailstone n)


