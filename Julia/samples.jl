
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


println("Regression Line Calculation")
linearRegression("SampleData.csv")
println("\nComparing serial vs parallel Processing ")
println("SET THE NUMBER OF THREADS > 1")
println(@time serialCalc(0, 100000000))
println(@time parallelCalc(0, 100000000))

randomDistribution(1000,10)
println()
randomDistribution(100000,10)
println()
randomDistribution(10000000,10)


