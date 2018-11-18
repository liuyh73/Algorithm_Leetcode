func max(a, b int) int {
    if a < b {
        return b
    }
    return a
}

func maxProfit(prices []int) int {
    days := len(prices)
    if days == 0 {
        return 0
    }
    f := make([]int, days)
    minn := prices[0]
    maxx := prices[days-1]
    sum := 0
    for i:=1; i<days; i++ {
        if prices[i] > prices[i-1] {
            f[i] = max(f[i-1], prices[i] - minn)
        } else {
            f[i] = f[i-1]
        }

        if prices[i] < minn {
            minn = prices[i]
        }
    }

    for i:=days-1; i>0; i-- {
        maxx = max(maxx, prices[i])
        if maxx - prices[i] > 0 {
            sum = max(sum, maxx - prices[i] + f[i-1])
        }
    }

    if maxx - prices[0] > 0 {
        sum = max(sum, maxx - prices[0])
    }

    return sum 
}