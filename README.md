# itermania
proof of concept of Go 1.23 range over func

# example

```go
prime := Bind(Inc(2), func(n int) Gen[int] {
	return Where(Const(n), All(Not(Eq(Mod(Const(n), Range(2, n, 1)), Const(0)))))
})

for i := range Head(prime, 10)() {
	fmt.Println(i)
}
```

```
2
3
5
7
11
13
17
19
23
29
```

# test

```bash
GOEXPERIMENT=aliastypeparams GODEBUG=gotypesalias=1 go test
```
