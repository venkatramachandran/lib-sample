# Go Library Sample #
=================

The requirements that I started with was to write a library in Go that can:

* Be compiled into a shared object library to be used in a Java webapp running on Tomcat
* Be compiled into an AAR to be used in an Android app

This is the code I ended up with. The code is organized into:

1. An `api` package that has the majority of the code.
2. A `main` package that exports all methods from `api`.
3. A COMPILE file to indicate how to compile the code for Android (using `gomobile`) and for Java
4. A bunch of sample code to invoke the .so (from C, Java (using JNI) and python)


This borrows heavily from http://blog.ralch.com/tutorial/golang-sharing-libraries/
