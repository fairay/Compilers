int main() {
    int array[9];
    int value = 1;
    int i = 0;

    while (i < 9) {
        array[i] = value;
        value = value + 1;
        i = i + 1;
    }

    int sum = 0;
    int j = 0;

    while (j < 9) {
        int k = 0;
        while (k < array[j]) {
            sum = sum + array[j];
            k = k + 1;
        }
        j = j + 1;
    }

    return sum;
}
