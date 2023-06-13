int main() {
    int arr[5];
    arr[0] = 10;
    arr[1] = arr[0] + 5;
    arr[2] = arr[1] * 2;
    arr[3] = arr[2] - arr[1];
    arr[4] = arr[3] / 3;

    arr[0] = arr[1] + arr[2];
    arr[1] = arr[3] - arr[4];
    arr[2] = arr[4] * arr[0];
    arr[3] = arr[0] / arr[4];
    arr[4] = arr[2] % arr[1];

    return arr[4];
}
