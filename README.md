# golang_project

`curl -O https://storage.googleapis.com/golang/go1.6.linux-amd64.tar.gz`

`sha256sum go1.6.linux-amd64.tar.gz`
`tar xvf go1.6.linux-amd64.tar.gz`
`sudo chown -R root:root ./go`
`sudo mv go /usr/local`

export PATH=$PATH:/usr/local/go/bin
export GOPATH=$HOME/sai/go
export GOROOT=/usr/local/go
export GOBIN=$GOPATH/bin
export PATH=$PATH:$GOROOT/bin




/* importent ponints of run code

go get  --- this one is for to get dependencies

go install ---- this one is for install dependencies

go run main.go ---- this one is for to run to run file name



/* importent points in code

Numeric Conversions

The most common numeric conversions are Atoi (string to int) and Itoa (int to string).

i, err := strconv.Atoi("-42")
s := strconv.Itoa(-42)
These assume decimal and the Go int type.

ParseBool, ParseFloat, ParseInt, and ParseUint convert strings to values:

b, err := strconv.ParseBool("true")
f, err := strconv.ParseFloat("3.1415", 64)
i, err := strconv.ParseInt("-42", 10, 64)
u, err := strconv.ParseUint("42", 10, 64)
The parse functions return the widest type (float64, int64, and uint64), but if the size argument specifies a narrower width the result can be converted to that narrower type without data loss:

s := "2147483647" // biggest int32
i64, err := strconv.ParseInt(s, 10, 32)
...
i := int32(i64)
FormatBool, FormatFloat, FormatInt, and FormatUint convert values to strings:

s := strconv.FormatBool(true)
s := strconv.FormatFloat(3.1415, 'E', -1, 64)
s := strconv.FormatInt(-42, 16)
s := strconv.FormatUint(42, 16)
AppendBool, AppendFloat, AppendInt, and AppendUint are similar but append the formatted value to a destination slice.



String Conversions



Quote and QuoteToASCII convert strings to quoted Go string literals. The latter guarantees that the result is an ASCII string, by escaping any non-ASCII Unicode with \u:

q := Quote("Hello, 世界")
q := QuoteToASCII("Hello, 世界")
QuoteRune and QuoteRuneToASCII are similar but accept runes and return quoted Go rune literals.

Unquote and UnquoteChar unquote Go string and rune literals.
