

# You are given a list of integers and a target number. Your task is to implement
# a function that finds a sublist of the given list that sums up to the target number. If no such sublist exists, the function should return an empty list.
#For example, if the given list is [1, 2, 3, 4, 5] and the target number
# is 9, the function should return [2, 3, 4].
#You are allowed to use dynamic arrays in your solution.

# Pseudocode
# create a tmp that holds all the possible numbers
# create a total that holds the total value of the tmp
# create a start (of the sublist) that starts at 0
# once value exceeds target and the start isn't at the end (for statement), sub sublist[start] and remove it
# add one to start and let for statemenet keep running
#if equals target, return sublist, else return []

def find_sublist(arr, target):
    sublist = []
    total = 0
    start = 0
    for end in range(len(arr)):
        total += arr[end]
        sublist.append(arr[end])
        while total > target and start < end:
            total -= sublist.pop(start)
            start += 1
        if total == target:
            return sublist
    return []

print(find_sublist([1, 2, 3, 4, 5], 9)) # [2, 3, 4]
print(find_sublist([1, 2, 3, 4, 5], 5)) # [1, 2]
print(find_sublist([1, 2, 3, 4, 5], 15)) # []
