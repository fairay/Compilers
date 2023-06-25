define i32 @main() {
main:
	%0 = alloca i32
	store i32 42, i32* %0
	br label %cond
cond:
	%1 = load i32, i32* %0
	%2 = icmp sge i32 %1, 7
	br i1 %2, label %body, label %next
body:
	%3 = load i32, i32* %0
	%4 = sub i32 %3, 7
	store i32 %4, i32* %0
	br label %cond
next:
	%5 = load i32, i32* %0
	ret i32 %5
}