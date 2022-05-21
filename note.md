QUICK GO REFERENCE
==================
## Array
```go
var x [5] int
var x [5][5] int
var x = [5]int{1,2,3,4,5}
x := [5]int{1,2,3,4,5}
```

## Slice
```go
var s = []int{1,2,3}
s := make([]int,10) //creaate a slice with length of 10 (10 slots) and each value is 0
s := make([]int,5,10) // create a slice with length of 5 and capacity of 10
s2 := s[n:m+1] // create new slice from old slice. n = indes you want to start from, m = index to end
```

## Map
```go
m := make(map[string]int)
delete(m, "key")
val, ok := m["key"] //if key does not exist ok will be false
key, value := range m //to be used in a loop
```
## String
```go
s := "Fanan"
b := []byte(s) //convert s to byte or []rune(s)
s[i] // accesses the character at index i of s. It is returned as byte.
ch := rune(97)
n := int('a')

//import the strings library
strings.Split(s, character)
strings.ToLower(s)
```

## Data Structures
Go doesn't have a lot of in-built data structures so you have to improvise with a few
1. For Stacks, you can use slices
2. for Queues you can use channels
```go
queue := make(chan int, size)
queue <- 10 // put 10 in queue
first := <- queue // get first character in queue
// to avoid blocking calls you can use select
select {
    case first := <- this.queue:
        // do something with first
    default:
        // avoid blocking calls
        fmt.Println("no activity")
}
```

## Numbers
```
const MaxUint = ^uint(0)
const MinUint = 0

const MaxInt = int(^uint(0) >> 1)
const MinInt = -MaxInt - 1

var kthDigit int = int(math.Floor(float64(a[i]) / math.Pow(10, float64(k)))) % 10//get the kth digit from behind
var numOfDigits int = int(math.Floor(math.Log10(float64(max)))) + 1//how many digits does this number have
```

PROBLEM SOLVING PATTERNS
===========================
1. Multiple Pointers - most cases, array must be sorted
2. Frequency counter - use maps to avoid nested loops
3. Window Slider - Used for a subset of data that exhibits a continuous property

MATH TIPS
============
1. `(idx % N)` will wrap around an array preventing the need of an IF statement
2. `x + (y - (x % y))` - next highest multiple of x from y
2. `n * (n + 1) / 2` total number of possible substrings of a string. It also returns sum of numbers from 0 to n
2. `n * (n - 1)/2` gives you number of palindromes in a repeated string
2. `[a * (1 - r^n)]/1-r` gives sum of n elements in a geometric sequence
2. `[floor(Log10(x)) + 1]` gives number of digits a number has
3. `[floor(x / pow(10,k) ) % 10]` gives the `kth` digit of number `x`
3. For any number K, the sum of 2 values (A & B) is evenly divisible by K if the sum of the remainders of `A/K + B/K `is `K`.
That is if `(A%K) + (B%K) == K`
3. `P(A and B) = P(A) * P(B)`
3. If A and B are mutually exclusive (e.g., if one happens, the other one can’t), `P(A or B) = P(A) + P(B)`
3. In general), `P(A or B) = P(A) + P(B) - P(A and B)`
3. `P(B given A) = P(A and B) / P(A)`
5. a ^ b = b ^ a; a^0 = a; a^a = 0; a^b^a = b where '^' = XOR (Can be used to find a single unique value in a list)

PITFALLS
===========
1. Try not to confuse **SUB-SET**, **SUB-SEQUENCE**, **SUB-ARRAY** / **SUB-STRING**
- **SUB-SET** => Any group of characters picked can suffice
- **SUB-SEQUENCE** => characters must come after one another but not immediately (e.g. For Fanan - Fnn, Fan, Fn is a subsequence not nanF)
- **SUB-ARRAY/SUB-STRING** => any group of elements that follow each other sequentially (e.g. For Fanan - Fan, anam nan, Fana)
