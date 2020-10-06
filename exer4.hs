
pascal :: Integer -> [Integer]
pascal 0 = [1]
pascal 1 = [1,1]
pascal x = [1] ++ [ uncurry (+) i | i <- tuples] ++[1]
    where   p = pascal (x - 1)
            tuples = zip p (tail p)
        

addPair :: Num a => (a,a) -> a
addPair = uncurry (+) 

withoutZeros :: (Eq a,Num a) => [a] -> [a]
withoutZeros = filter (/=0)


findElt :: Eq a => a -> [a] -> Maybe Int
findElt item arr =  getIndex 0
    where 
        getIndex :: Int -> Maybe Int
        getIndex index  
            | item == arr !!index   = Just (index)
            | index+1 == length arr = Nothing
            | otherwise             = getIndex (index+1)

