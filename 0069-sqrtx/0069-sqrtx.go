
func mySqrt(x int) int {
    if x == 0 {
        return 0
    }

    left, right := 1, x
    res := 0

    for left <= right {
        mid := (left + right)/2

        if mid <= x/mid {
            res = mid
            left = mid + 1
        } else {
            right = mid - 1
        }
    }

    return res
}
