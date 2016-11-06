Package bitmap provides a simple bitmap which allows individual bits to be set and read within a slice of bytes.

The package was inspired by the book "Programming Pearls," by Jon Bentley. The
first example in the book is a clever sorting solution which uses a bitmap to
sort a file containing up 10 million numeric values in a single pass without
loading them all into memory. 

http://www.cs.bell-labs.com/cm/cs/pearls/cto.html

```
func ExampleBitmap () {
    b := bitmap.New(10)
    b.Set(2)
    fmt.Printf("2 in bitmap: %v. 7 in bitmap: %v.\n", b.IsSet(2), b.IsSet(7))
    // Output: 2 in bitmap: true. 7 in bitmap: false.
}

func ExampleValues () {
    b := bitmap.New(10)
    b.Set(2)
    b.Set(7)
    fmt.Println(b.Values())
    // Output: [2 7]
}
```
