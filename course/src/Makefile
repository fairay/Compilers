all: output.exe

buildDir = build

$(buildDir):
	mkdir $(buildDir)

output.ll:
	go run .

output.obj:
	clang -c -o $(buildDir)/output.obj $(buildDir)/output.ll

output.exe: output.obj
	lld-link /subsystem:console /out:$(buildDir)/output.exe $(buildDir)/output.obj

generate:
	go generate ./...

run: output.exe
	-$(buildDir)/output.exe

full: output.ll run

clear:
	rm -rf $(buildDir)
