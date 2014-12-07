/* Package bitmap provides a simple bitmap which allows individual bits to be
set and read from a slice of bytes.

The package was inspired by the book "Programming Pearls," by Jon Bentley. The
first example in the book is a clever sorting solution which uses a bitmap to
sort a file containing up 10 million numeric values in a single pass without
loading them all into memory. 

http://www.cs.bell-labs.com/cm/cs/pearls/cto.html
*/
package bitmap
