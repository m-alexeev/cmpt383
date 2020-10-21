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


println(mergeSort([8,3,2,6]))






# *  Integer div symbol: ÷  