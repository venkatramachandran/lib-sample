## Compiling the code in a shared library ##
go build -o libsample.so -buildmode=c-shared .

## Compiling the code into an Android aar ##
gomobile bind -target=android api

## Compiling C code ##
gcc -Wall -o sample -L $(pwd)  blah.c -lsample

## Compiling and executing Java code ##
javac -cp ./jna-4.2.2.jar:. sample.java && java -cp ./jna-4.2.2.jar:. sample
