function fibonacci(n)
    a = 0
    b = 1

    for i = 0, n, 1 do
        a, b = b, a + b
    end

    return a
end

print(fibonacci(7))