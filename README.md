# go-paystack

This is a Go client library for Paystack, a payments platform that allows you to accept payments from customers in 154+ countries.

## Installation

```bash
go get https://github.com/rammyblog/go-paystack
```

## Usage

```go
import (
        "https://github.com/rammyblog/go-paystack"
        "context"
        "time"

        )


ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
client := paystack.NewClient("your_secret_key")

// Create a new transaction
	resp, err := c.Transaction.Initialize(ctx, &paystack.TransactionRequest{
		Amount:      100000,
		Email:       "Onas@gmail.com",
		Currency:    "NGN",
		Reference:   "yinmu",
		CallbackURL: "https://ngrok.com/rammybloh",
	})
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Printf("\n Initialize transaction \n-%+v\n", resp.AuthorizationURL)
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
