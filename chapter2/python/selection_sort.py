def selection_sort(arr: []) -> []:
    sorted_arr = []
    l = len(arr)
    for _ in range(l):
        m = max(arr)
        sorted_arr.append(m)
        arr.remove(m)
    return sorted_arr


if __name__ == "__main__":
    print(selection_sort([4,1,3,2]))