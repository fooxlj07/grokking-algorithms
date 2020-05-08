def binary_search(nums: list, num: int):
    min, max = 0, len(nums)-1
    if nums[min] == num:
        return min+1
    if nums[max] == num:
        return max+1

    while min < max:
        mid = (min + max) // 2
        if nums[mid] == num:
            return mid+1
        if nums[mid] < num:
            min = mid
        if nums[mid] > num:
            max = mid

    return 0


if __name__ == "__main__":
    print(binary_search([i for i in range(1, 100)], 50))
    print(binary_search([i for i in range(1, 100)], 98))
