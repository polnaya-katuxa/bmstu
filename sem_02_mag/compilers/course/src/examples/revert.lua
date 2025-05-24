function revert_array(arr)
    l = 0
    for i, v in arr do
        l = l + 1
    end

    res = {}
    cur = 0
    for i = l-1, -1, -1 do
        res[cur] = arr[i]
        cur = cur + 1
    end

    return res
end

inp = {0,1,2,3,4}
print(revert_array(inp))