func searchRange(nums []int, target int) []int {
    leftOccurrence := -1
    rightOccurrence := -1

    // Find first occurrence
    start, end := 0, len(nums)-1

    for start <= end {
        mid := start + (end-start)/2

        if nums[mid] == target {
            leftOccurrence = mid
            end = mid - 1
        } else if nums[mid] < target {
            start = mid + 1
        } else {
            end = mid - 1
        }
    }

    // Target not found
    if leftOccurrence == -1 {
        return []int{-1, -1}
    }

    // Find last occurrence
    start, end = 0, len(nums)-1

    for start <= end {
        mid := start + (end-start)/2

        if nums[mid] == target {
            rightOccurrence = mid
            start = mid + 1
        } else if nums[mid] < target {
            start = mid + 1
        } else {
            end = mid - 1
        }
    }

    return []int{leftOccurrence, rightOccurrence}
}