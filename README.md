# go-paystack

This is a Go client library for Paystack, a payments platform that allows you to accept payments from customers in 154+ countries.

## Installation

```bash
go get https://github.com/rammyblog/go-paystack
```

## Usage

```go
import "https://github.com/rammyblog/go-paystack"

client := paystack.NewClient("your_secret_key")

// Create a new transaction
transaction, err := client.Transaction.Create(&paystack.TransactionRequest{
    Amount:   5000,
    Email:    "customer@example.com",
    Currency: "NGN",
})
```

## Testing

To run the tests, execute the following command:

```bash
go test ./...
```

## Contributing

Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License

[MIT](https://choosealicense.com/licenses/mit/)


## TODO

- [ ] Add more tests
- [ ] Add more examples
- [ ] Add more documentation
- [ ] Fix Logging
