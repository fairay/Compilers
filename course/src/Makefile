all: output.exe

buildDir = build

$(buildDir):
	mkdir $(buildDir)

output.obj:
	clang -c -o $(buildDir)/output.obj $(buildDir)/output.ll

mainCRTStartup.obj: $(buildDir)
	clang -c -o $(buildDir)/mainCRTStartup.obj c/mainCRTStartup.c

output.exe: output.obj mainCRTStartup.obj
	lld-link /subsystem:console /out:$(buildDir)/output.exe $(buildDir)/output.obj $(buildDir)/mainCRTStartup.obj

generate:
	go generate ./...

clear:
	rm -rf $(buildDir)
