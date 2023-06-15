int main() {
    int a = 10;
    int b = 20;
    int c = 30;
    int result;

    if (a > b) {
        if (a > c) {
            result = a;
        } else {
            result = c;
        }
    } else {
        if (!(b < c)) {
            result = b;
        } else {
            result = c;
        }
    }

    return result;
}
