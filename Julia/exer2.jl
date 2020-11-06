using Dates 

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

function merge(lst1, lst2)
    arr = []
    #basic merge 
    while length(lst1) > 0 && length(lst2) >0
        if lst1[begin] < lst2[begin]
            append!(arr,lst1[begin])
            lst1 = lst1[2:end]
        else
            append!(arr, lst2[begin])
            lst2 = lst2[2:end]   
        end
    end  

    #Concat the rest of the elemts
    arr = [arr; lst1; lst2]

    return arr
end


function mergeSort(lst)
    if length(lst) <= 1
        return lst
    else 
        # Split the array 
        left = lst[1:length(lst)÷2]
        right = lst[(length(lst)÷2)+1:end]
        
        #Recursively sort the halves
        left = mergeSort(left)
        right = mergeSort(right)
        
        #Merge the 2 sorted halves
        merge(left,right)
    end 
end



function isFriday(date)
    return  (Dates.dayname(date) == "Friday") ? true : false
end


function isPrimeDay(date)
    day = Dates.day(date)
    return (length(divisors(day)) == 0) ? true : false 
end



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


# * Testing Merge
println("Testing Merge")
println(merge([3,4],[5,6]))
println(merge([1],[2]))
println(merge([1],[2,3]))
println(merge([1],[]))
println(merge([],[1]))
println(merge([1,2],[]))
println(merge([],[1,2]))

println("================")

# * Testing MergeSort 
println("Testing MergeSort")
println(mergeSort([8,3,2,6]))
println(mergeSort([8,2,6,2]))
println(mergeSort([8,2,6]))
println(mergeSort([2,1]))
println(mergeSort([1]))
println(mergeSort([]))

 
#* Test if day is Friday
println("Testing Friday\n")
println("Tuesday 2020,10,20   ", isFriday(DateTime(2020,10,20))) #Tuesday
println("Wednesday 2020,10,21 ", isFriday(DateTime(2020,10,21))) #Wednesday
println("Thursday 2020,10,22  ", isFriday(DateTime(2020,10,22))) #Thursday
println("Friday 2020,10,23    ", isFriday(DateTime(2020,10,23))) #Friday
println("\n====================\n")

#* Test if day is Prime 
println("Testing PrimeDay\n")
println("20 ", isPrimeDay(DateTime(2020,10,20))) #Tuesday
println("21 ", isPrimeDay(DateTime(2020,10,21))) #Wednesday
println("22 ", isPrimeDay(DateTime(2020,10,22))) #Thursday
println("23 ", isPrimeDay(DateTime(2020,10,23))) #Friday