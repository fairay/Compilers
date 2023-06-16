int main() {
    int matrix1[2][2];
    int matrix2[2][2];
    int result[2][2];
    int i, j, k;

    // Заполняем первую матрицу
    matrix1[0][0] = 1;
    matrix1[0][1] = 2;
    matrix1[1][0] = 3;
    matrix1[1][1] = 4;

    // Заполняем вторую матрицу
    matrix2[0][0] = 5;
    matrix2[0][1] = 6;
    matrix2[1][0] = 7;
    matrix2[1][1] = 8;

    // Умножаем матрицы
    i = 0;
    while (i < 2) {
        j = 0;
        while (j < 2) {
            result[i][j] = 0;
            k = 0;
            while (k < 2) {
                result[i][j] = result[i][j] + matrix1[i][k] * matrix2[k][j];
                k = k + 1;
            }
            j = j + 1;
        }
        i = i + 1;
    }

    // Возвращаем значение первого элемента результирующей матрицы
    return result[0][0];
}
