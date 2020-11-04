using Base.Threads
function serialCalc(x, y)
    for i in 1:y
        x += 1
    end
    return x
end 

##Multi threaded example, set number of threads to >1 to see results with "julia --threads 4"
function parallelCalc(x,y)
    @threads for i in 1:y
        atomic_add!(x, UInt32(1))
    end
    return x[]
end


## To run this code, CSV and Plots must be installed for 
## Reading data
## import Pkg; Pkg.add("CSV")
using CSV
function linearRegression(filename)
    data = CSV.File(filename)
    x = 0;y = 0; x_sq = 0; xy = 0;
    n = 0
    for row in data
        y += row.LE; x += row.GDP
        xy += row.LE + row.GDP; x_sq += row.GDP ^ 2
        n +=1 
    end

    a = ((y * x_sq) - (x * xy)) / ((n*x_sq) - (x^2)) 
    b = ((n*xy) - (x * y)) / ((n * x_sq) - x^2)

    print("Regression Line with  Y-Int = $(a), Slope = $(b)")
end 


#linearRegression("SampleData.csv")
println(@time serialCalc(0, 100000000000))
println(@time parallelCalc(Atomic{UInt32}(0), 100000000000))