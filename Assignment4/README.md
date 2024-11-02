# www-phone
A Phone Book application that works with HTTP

Run without compiling:
```
go run handler.go www-phone.go
```

Compile and run:
```
go build handler.go www-phone.go
./handlers
```

Running Curl commands from CLI:

List Command:
```
curl localhost:1234/list
```

Output:
```
Jane Doe 0800123456
Michalis Tsoukalos 2101112223
Mike Tsoukalos 2109416471
```
