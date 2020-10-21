function divisors(num)
   return [i for i = 2:num-1 if num%i==0] 
end 

function primes(num)
    return [i for i = 2:num if length(divisors(i)) == 0]
end

function pythagTriples(num)
    return [(a,b,c) for a = 1:num, b = 1:num, c=1:num if a^2+b^2==c^2 && b>a]
end 

#TODO: Refactor 

function join(separ, strings)
    ret = ""
    #If last item dont append the separator
    arr = map((x) -> strings[end]==x ? x : x*separ  , strings)
    for item in arr
        ret *= item
    end
    return ret
end


#=
# * Divisors Tests
println(divisors(30))
println(divisors(64))
println(divisors(127))

# * Primes Tests
println(primes(7))
println(primes(100))

# * PythagTriples Tests
println(pythagTriples(10))
println(pythagTriples(30))
 
# * Tests for join func 
println(join("+", ["a","b","c"]))
println(join("+", ["a"]))
println(join("+", []))

=# 

