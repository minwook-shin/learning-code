Python 3.7.0a4 (v3.7.0a4:07c9d85, Jan  9 2018, 07:07:02) [MSC v.1900 64 bit (AMD64)] on win32
Type "copyright", "credits" or "license()" for more information.
>>> a= 12
>>> b= 6
>>> bin(a)
'0b1100'
>>> bin(b)
'0b110'
>>> c = a & b
>>> bin(c)
'0b100'
>>> d = a | b
>>> bin(d)
'0b1110'
>>> e = a ^ b
>>> bin(e)
'0b1010'
>>> f = ~a
>>> bin(f)
'-0b1101'
>>> f2 = ~b
>>> bin(f2)
'-0b111'
>>> g = a<<2
>>> bin(g)
'0b110000'
>>> g2 = a>>2
>>> bin(g2)
'0b11'
>>> int(f)
-13
>>> 