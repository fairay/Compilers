int main() {
    int a = 30;
    int b = 20;
    int result;

    if (a > b) {
        int temp = a * 2;
        result = temp - 5;
        return result;
    } else if (a < b) {
        int temp = b / 2;
        result = temp + 10;
        return result;
    } else {
        int temp = a + b;
        result = temp * 3;
        return result;
    }
    // TODO: no return here
    return result;
}
