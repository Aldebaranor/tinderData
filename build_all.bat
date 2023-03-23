set current_dir=%~dp0
set CGO_CFLAGS=-I%current_dir%cgo\include
set CGO_LDFLAGS=-L%current_dir%cgo\lib -l:libzmq.lib
set CGO_ENABLED=1


go build -o ./manifest/bin/tinder-data.exe ./main.go

pause