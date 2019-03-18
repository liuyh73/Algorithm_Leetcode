const INT_MAX = int(^uint(0) >> 1)
const INT_MIN = ^INT_MAX
func max(nums ...int) (maxx int) {
    maxx = INT_MIN
    for _, num := range(nums) {
        if maxx < num {
            maxx = num
        }
    }
    return
}

func min(nums ...int) (minn int) {
    minn = INT_MAX
    for _, num := range(nums) {
        if minn > num {
            minn = num
        }
    }
    return 
}

func maxProduct(nums []int) int {
    lenNums := len(nums)
    localMin := nums[0]
    localMax := nums[0]
    res := nums[0]
    for i:=1; i<lenNums; i++ {
        tempMax := localMax
        localMax = max(localMax * nums[i], localMin * nums[i], nums[i])
        localMin = min(tempMax * nums[i], localMin * nums[i], nums[i])
        res = max(localMax, res)
    }
    return res
}