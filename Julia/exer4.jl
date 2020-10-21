using Dates


function divisors(num)
    return [i for i = 2:num-1 if num%i==0] 
end 
 

function isFriday(date)
    return  (Dates.dayname(date) == "Friday") ? true : false
end




function isPrimeDay(date)
    day = Dates.day(date)
    return (length(divisors(day)) == 0) ? true : false 
end



 
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