func getConcatenation(nums []int) []int {
n:=len(nums)
ans:=make([]int,2*n)

for index,num:=range nums {
    ans[index]=num
    ans[index+n]=num
}
return ans
}
