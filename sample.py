#!/usr/bin/python
import ctypes
lib = ctypes.CDLL("./libsample.so")
lib.Hello("World")
lib.Hello2.restype = ctypes.c_char_p
print lib.Hello2("New World")