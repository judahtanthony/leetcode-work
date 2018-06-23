func reverse(x int) int {
    n := 1;
    if x == 0 {
        return x;
    }
    INT_MAX := (1 << 31) - 1
    if x <= 0 {
        n *= -1;
        x *= -1;
    }
    r := 0;
    for x > 0 {
        // Overflow check.
        if r > INT_MAX / 10 {
            return 0;
        }
        r *= 10;
        ones := x % 10;
        // Overflow check.
        if r > INT_MAX - ones {
            return 0;
        }
        r += ones;
        x /= 10;
    }
    return r * n;
}
