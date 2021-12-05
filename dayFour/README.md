### Day 4

Giant Squid

Super happy for today i've decided to use Go at the start of this project!

For the Bingo problem having the ability to create multi dimensional arrays is a must!

In these puzzles I've chosen to use Tensors (3 Dimensional Arrays) which are fairly hard to visualize but are very powerful.

Here is an example of a 2 Dimensional Array more commonly known as a Matrix.

```
[0,0,0,0]
[0,5,0,0]
```
In Go we can access values in the Matrix like this.

```go
fmt.Println(matrix[1][1])

// Output: 5
```

Now we can visualize a 2 Dimensional array lets try a tensor.

```go
// Matrix 1 (0)
[0,0,0,0]
[0,0,0,0]
[0,0,0,0]

// Matrix 2 (1)
[0,0,0,0]
[0,5,0,0]
[0,0,0,0]
```

Now for a tensor we can do this.

```go
fmt.Println(tensor[1][1][1])

// Output: 5
```

Thats the best way I can try and explain this code