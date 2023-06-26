int power(int base, int exponent) {
    int result = 1;
    int i = 0;
    while (i < exponent) {
        result = result * base;
        i = i + 1;
    }
    return result;
}

int main() {
    int result = power(2, 5);
    return result;
}
