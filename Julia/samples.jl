##! Julia Weakness ##

function ListComprehensionWeakness(num)
    #Iterators cannot be reused in list comps in 
    #Julia, this means more work on the CPU, and redundancy

    #! Julia NON working example 
    ## This will not compile because i cannot be reused in a since its not declared in scope of b
    # array = [(a,b,c) for a in 1:num, b = a:num, c = 1:num if a^2+b^2==c^2]
   
    #! Working example
    #Not good in terms of efficiency since it will loop throught all permutations 
    #of a,b,c 
    array = [(a,b,c) for a = 1:num, b = 1:num, c=1:num if a^2+b^2==c^2 && b>a]
    #println("List Comp Pythag $(array)\n")
   
    ##! Working Python Example using List ListComprehensions
    # array = [(a,b,c) for a in 1:num, b = a:num, c = 1:num if a^2+b^2==c^2]

    return array
end 

function ListCompForLoop(num)
    #Optimized example not using list Comps without extra iterations
    # Runs much faster 
    array = []
    for a in 1:num
        for b in a:num
            for c in b:num
                if a^2+b^2==c^2
                    push!(array, (a,b,c))
                end
            end
        end
    end

    return array

end



##! Common Julia Usages ##

using Random
function serialCalc(x, y)
    for i in 1:y
        x += rand((1:10))
    end
    return x
end 

using Base.Threads
##Multi threaded example, set number of threads to >1 to see results 
##On windows set the env variable to "set JULIA_NUM_THREADS=4" if on CMD 
##On Linux "export JULIA_NUM_THREADS=4"
##If the computer has at least 4 threads

function parallelCalc(x,y)
    #Do serial calc for small y
    if (y <= 2000000)
       return serialCalc(x,y)
    #Otherwise split the process
    else
        half = Threads.@spawn (parallelCalc(0, y ÷ 2))
        right = parallelCalc( 0, y ÷ 2 )
        return fetch(half) + right
    end
   
end


## To run this code, CSV and Plots must be installed for 
## Reading data
## import Pkg; Pkg.add("CSV"); 
using CSV
using Statistics
function linearRegression(filename)
    data = CSV.File(filename)

    # Place rows into lists
    x = [row.GDP for row in data]
    y = [row.LE for row in data]

    n = length(x)
    ## Means of the X and Y
    mean_x = mean(x)
    mean_y = mean(y)

    #Cross deviation
    xy = sum([x[i]*y[i] for i in 1:n]) - (n * mean_x * mean_y)
    xx = sum([x[i]^2 for i in 1:n]) - (n * mean_x^2)

    a = xy / xx
    b = mean_y - (a * mean_x)

    println("Regression Line with  Y-Int = $(b), Slope = $(a)")

end 

using Printf
## Julias Random function number distribution 
function randomDistribution(nums, maxNum)
    rands = Dict{Integer,Integer}()

    for i in 1:nums
        gen = rand((1:maxNum))
        push!(rands, gen => get(rands,gen,0) + 1)

    end

    percents = Dict{Integer,Float32}()
    for key in rands
        push!(percents, key[1]=>get(rands,key[1],0) / nums)
    end
    
    #Sort Dictionary
    percents = sort(collect(percents), by= x->x[1])
    println("Distribution of $(nums) random values in the range: 1:$(maxNum)
===================================================")
    for key in percents
        println("Value: $(key[1]) => $(key[2])%")
    end 

end


println(@time ListComprehensionWeakness(10))
println(@time ListCompForLoop(10))

println("Regression Line Calculation")
linearRegression("SampleData.csv")
println("\nComparing serial vs parallel Processing\n ")
println("SET THE NUMBER OF THREADS > 1 OTHERWISE WOULD BE NO DIFFERENCE\n ")
println(@time serialCalc(0, 100000000))
println(@time parallelCalc(0, 100000000))

randomDistribution(1000,10)
println()
randomDistribution(100000,10)
println()
randomDistribution(10000000,10)


