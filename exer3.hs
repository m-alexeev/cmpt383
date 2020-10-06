import Data.Time.Calendar
import Data.Time.Calendar.OrdinalDate

merge :: Ord a => [a] -> [a] -> [a]
merge [] [] = []
merge [] (y:ys) = y : merge [] ys 
merge (x:xs) [] = x : merge xs [] 
merge (x:xs) (y:ys) 
    | x <= y = x : merge xs (y:ys) 
    | x > y = y : merge (x:xs) ys

mergeSort :: Ord a => [a] -> [a]
mergeSort [] = []
mergeSort [x] = [x]
mergeSort x = merge first last
    where   half = length x `div` 2
            first = mergeSort (take half x)
            last  = mergeSort (drop half x)


daysInYear :: Integer -> [Day]
daysInYear y =[jan1 .. dec3]
    where   jan1 = fromGregorian y 0 0 
            dec3 = fromGregorian y 12 30

isFriday :: Day -> Bool
isFriday x 
    | snd (mondayStartWeek x) == 5 = True
    | otherwise = False


divisors :: Int->[Int]
divisors n = [i | i <- [2..(n `div` 2)], n `mod` i == 0] -- range (2 .. n/2), keeps all values n % i == 0

getDay :: (Integer, Int, Int) -> Int
getDay (y,m,d) = d

isPrimeDay :: Day -> Bool
isPrimeDay x
    | divisors (getDay (toGregorian x)) == [] = True
    | otherwise = False


primeFridays :: Integer -> [Day]
primeFridays x = filter isPrimeDay(fridays)
    where   fridays =  filter isFriday(daysInYear x)
