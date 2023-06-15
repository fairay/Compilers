int main() {
    int number = 100;
    int count = 0;
    int shouldContinue = 1;

    while (shouldContinue) {
        if (number % 2 == 0) {
            number = number / 2;
        } else {
            number = number - 3;
            count = count + 1;
        }

        if (count >= 5) {
            shouldContinue = 0;
        }
    }

    return number;
}