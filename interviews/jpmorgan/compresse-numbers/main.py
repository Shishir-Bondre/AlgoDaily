def compressNumbers(nums):
    prev = nums[0]
    count = 1
    res = ""
    for i in range(1, len(nums)):
        # we do necessarily need to parse numbers to int, just treat them as string
        c = nums[i]
        if c == prev:
            count += 1
        else:
            res += str(count) + ' ' + prev + ' '
            prev = c
            count = 1
    res += str(count) + ' ' + prev
    return res


# input() reads a string with a line of input, stripping the ' ' (newline) at the end.
# This is all you need for most Code Jam problems.
# read input to know N
t = int(input())
# read N lines
for i in range(1, t + 1):
    ri = input()  # raw_input() for python2.7
    nums = ri.split()
    result = compressNumbers(nums)
    print(result)