define i32 @main() {
main:
	%0 = alloca i32
	store i32 10, i32* %0
	%1 = load i32, i32* %0
	%2 = icmp sgt i32 %1, 100
	%3 = alloca [100 x i32]
	%4 = alloca i32
	br i1 %2, label %if-1, label %else-1

if-1:
	%5 = sub i32 0, 1
	ret i32 %5

else-1:
	br label %main-1

main-1:
	%6 = getelementptr [100 x i32], [100 x i32]* %3, i32 0, i32 0
	store i32 0, i32* %6
	%7 = getelementptr [100 x i32], [100 x i32]* %3, i32 0, i32 1
	store i32 1, i32* %7
	store i32 2, i32* %4
	br label %while.cond-4

while.cond-4:
	%8 = load i32, i32* %4
	%9 = load i32, i32* %0
	%10 = icmp sle i32 %8, %9
	br i1 %10, label %while.body-4, label %main-4

while.body-4:
	%11 = load i32, i32* %4
	%12 = getelementptr [100 x i32], [100 x i32]* %3, i32 0, i32 %11
	%13 = load i32, i32* %4
	%14 = sub i32 %13, 1
	%15 = getelementptr [100 x i32], [100 x i32]* %3, i32 0, i32 %14
	%16 = load i32, i32* %4
	%17 = sub i32 %16, 2
	%18 = getelementptr [100 x i32], [100 x i32]* %3, i32 0, i32 %17
	%19 = load i32, i32* %15
	%20 = load i32, i32* %18
	%21 = add i32 %19, %20
	store i32 %21, i32* %12
	%22 = load i32, i32* %4
	%23 = add i32 %22, 1
	store i32 %23, i32* %4
	br label %while.cond-4

main-4:
	%24 = load i32, i32* %0
	%25 = getelementptr [100 x i32], [100 x i32]* %3, i32 0, i32 %24
	%26 = load i32, i32* %25
	ret i32 %26
}

define i32 @mainCRTStartup() {
0:
	%1 = call i32 @main()
	ret i32 %1
}
