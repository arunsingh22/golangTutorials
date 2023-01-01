
- **Character Set**
A character set or simply charset is a table of unique numbers assigned to different characters like letters, numbers and other symbols. There are many characters set like ASCII, UTF, ISO, and others. For example, the value of character A in the ASCII character set is 65 (decimal).

UTF stands for "UCS (Unicode) Transformation Format". The UTF-8 encoding can be used to represent any Unicode character. Depending on a Unicode character's numeric value, the corresponding UTF-8 character is a 1, 2, or 3 byte sequence.

Grapheme = a single unit of human writing systemm 
eg: d, 👌, (❁´◡`❁) etc.

- In unicode, a grapheme != one single codepoint != a byte.
    - Use unicode-unaware functions iff every grapheme is a byte: ASCII or UTF-8 with ASCII chars.
    - Use unicode-aware functions to properly handle codepoints.
    

**NOTE** 
1. Golang doesn't have char datatype 
2. In golang things are either string or rune(int32 aka 4 bytes), therefore i can say that
   in golang the char len is 4bytes by default, unlike other languages where the char len is either 1 byte or 2 bytes
3. To deal with individual rune(aka char) we use **unicode pkg** and to deal wi
   strings we need **strings pkg** 




