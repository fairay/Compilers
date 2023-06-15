int main() {
    int matrix[3][3];

    {
        int i = 0;
        while (i < 9) {
            matrix[i/3][i%3] = i;
            i = i + 1;
        }
    }

    int count = 0;
    { 
        int i = 0;
        while (i < 9) {
            count = count + matrix[i/3][i%3];
            i = i + 1;
        }
    }
    return count;
}
