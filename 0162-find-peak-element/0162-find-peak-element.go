func findPeakElement(nums []int) int {
    n := len(nums)

    // Edge case: only one element
    if n == 1 {
        return 0
    }

    // Edge case: first element is a peak
    if nums[0] > nums[1] {
        return 0
    }

    // Edge case: last element is a peak
    if nums[n-1] > nums[n-2] {
        return n - 1
    }

    left, right := 1, n-2

    for left <= right {
        mid := left + (right-left)/2

        // Check if mid is a peak
        if nums[mid] > nums[mid-1] && nums[mid] > nums[mid+1] {
            return mid
        } else if nums[mid] < nums[mid+1] {
            // Increasing slope: search right
            left = mid + 1
        } else {
            // Decreasing slope: search left
            right = mid - 1
        }
    }

    return -1
}