det a b c = b^2 - 4*a*c
quadsol1 a b c = (-b - sqrt (det a b c))/2*a
quadsol2 a b c = (-b + sqrt (det a b c))/2*a
third_a ls = ls !! 2
third_b (_:_:c:arr) = c

fact 0 = 1
fact x = x * fact (x-1)

hailstone x 
    |mod x 2 == 0 = div x 2
    |mod x 2 >= 0 = 3 * x + 1

hailLen 1 = 0
hailLen x = 1 + hailLen (hailstone x) 