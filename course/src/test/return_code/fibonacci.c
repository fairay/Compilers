int main() {
    int n = 10;

    if (n > 100) {
        return -1;
    }

    int fib[100];
    fib[0] = 0;
    fib[1] = 1;

    int i = 2;
    while (i <= n) {
        fib[i] = fib[i-1] + fib[i-2];
        i = i + 1;
    }

    return fib[n];
}
