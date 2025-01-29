func calculate(a, b int) int {
    if a == 0 {
        return b
    }
    return calculate(b%a, a)
}