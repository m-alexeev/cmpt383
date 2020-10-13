import Data.Ratio

myIterate :: (a -> a)-> a -> [a]
myIterate func a = a : (myIterate func(func a)) 


mySplitAt :: Int -> [a] -> ([a],[a])
mySplitAt x list = tup x [] list 
  where 
    tup :: Int -> [a] -> [a] -> ([a], [a])
    tup 0 ls1 ls2 = (ls1, ls2)
    tup _ ls1 []  = (ls1, [])
    tup x ls1 ls2 = tup (x - 1) ( ls1 ++ head(ls2) : []) (tail(ls2))



rationalSum :: Integral a => a -> [Ratio a]
rationalSum x = [ i % j | i <- [1..x],j <- [x-1,x-2..1], i+j ==x ]


rationalSumLowest :: Integral a => a -> [Ratio a]
rationalSumLowest x =  [ i % j | i <- [1..x-1] , let j = x - i,
                                 gcd x i == 1 && gcd x j == 1, i + j == x ]

-- Recursive version 
-- rationalSumLowest :: Integral a => a -> [Ratio a]
-- rationalSumLowest x = sumLowest 1 (x-1) x
--   where 
--     sumLowest _ _ 0 = []
--     sumLowest _ 0 _ = []
--     sumLowest y x z
--       | y > 0 && x > 0 && min  = y % x : sumLowest (y+1) (x-1) z
--       | otherwise = sumLowest (y+1) (x-1) z
--       where
--       min = gcd z x == 1 && gcd z y == 1


rationals :: [Ratio Integer]
rationals = concat(map rationalSumLowest [1..])


-- Helper functions 
splitAtSeparator :: Eq a => a -> [a] -> [[a]]
splitAtSeparator sep [] = []
splitAtSeparator sep content = first : splitAtSeparator sep rest
  where 
    first = takeWhile(/= sep) content
    firstlen = length first
    rest = drop (firstlen+1) content 


readInt :: String -> Int
readInt = read


fileName = "input.txt"

sumFile :: IO ()
sumFile = do 
    text <- readFile fileName
    let split = splitAtSeparator '\n' text
    let res = sum (map readInt split)
    print res

