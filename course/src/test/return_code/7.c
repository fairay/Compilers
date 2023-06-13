int main() {
    int arr[6];
    int a = 10, b = 20, c = 30, d = 40, e = 50;
    arr[0] = a + b;
    arr[1] = arr[0] * (c - d);
    arr[2] = arr[1] / (a * b);
    arr[3] = arr[2] - e;
    arr[4] = arr[0] + arr[1] + arr[2] + arr[3];

    int index = (arr[4] % 5 + 5) % 5;
    arr[5] = arr[index];
    
    return arr[5];
}
