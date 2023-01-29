
# Write a function largest_average(arr) that takes in a list of integers 
# as input, and returns the largest average of any contiguous sublist in the
# input list. If the input list is empty, the function should return None. 
# Note that the input list may have negative numbers.

# Pseudocode:
# create an empty array, fill with values
# create function which checks if list is empty, return None
# initialize temp and count in function
# add each number from the array to temp, then increment count by one 
# if there is no number numbers to add from array to temp, return temp[0]/count

def largestavg(arr):
    if len(arr) < 1:
        return None
    tmp = 0
    count = 0 
    for i in range(len(arr)):
        tmp += arr[i]
        count += 1
    return tmp/count
        

myarray = [None] *5


myarray[0],myarray[1],myarray[2],myarray[3],myarray[4] = 1, 2, 3, 4, 5

print(largestavg(myarray))
