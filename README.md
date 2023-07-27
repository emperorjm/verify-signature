# Sign Arbitrary Messages

This application allows you to verify a signed message from any Cosmos SDK based application.

## Requirements

- go >= 1.20 (https://go.dev/dl)

## Configuration

Make a copy of the **.env.example** file and rename it to **.env**. The following information will explain the values contained within:

- **PUB_KEY_HEX**: This is the public key from the account that signed the message.
- **SIG_HEX**: This is the signature generated from signing the message.
- **MSG**: This is the message that the user signed.

## Execute

In order to sign the message, first clone the repository. Then, navigate to the directory containing the code and execute the following commands:
- `go build`
- `go run main.go`

If the signature is valid, the following message should be displayed **Signature is valid!**.
