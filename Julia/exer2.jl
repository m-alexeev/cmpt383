function divisors(num)
   return [i for i = 2:num-1 if num%i==0] 
end 

function primes(num)
    return [i for i = 2:num if length(divisors(i)) == 0]
end

function pythagTriples(num)
    return [(a,b,c) for a = 1:num, b = 1:num, c=1:num if a^2+b^2==c^2 && b>a]
end 


function join(separ, strings)
    ret = ""
    for i in 1:length(strings)
        if i != length(strings)
            ret *= strings[i] * separ
        else
            ret *= strings[i]
        end
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