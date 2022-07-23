# Go unit testing
This is an example for doing unit testing in go
## How to run
```sh
# Run all test cases
go test ./tests -v

# Run all test cases under TestInt
go test ./tests -v -run TestInt 

# Run sub test cases under TestInt using regex
go test ./tests -v run TestInt/one
```