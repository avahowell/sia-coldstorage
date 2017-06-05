# sia-coldstorage-json

`sia-coldstorage-json` is a utility that generates a seed and a collection of addresses for a Sia wallet, without the blockchain. This is useful for creating 'cold wallets'. This utility is forked from `sia-coldstorage` (https://github.com/johnathanhowell/sia-coldstorage) for the purpose or returning the seed and addresses as a JSON object that can be used to create a paper wallet web utlity for Siacoin.

## USAGE

Download the latest release for your platform, then navigate to the binary in your terminal and run `./sia-coldstorage-json`.
This will generate your seed and addresses, returning it to you in JSON format. Ideally, you would run this on a
system that was very secure, i.e. an airgapped LiveCD or through a service like Amazon Lambda to generate through an online web app.

## NOTE

Addresses in JSON output need to have `trim` run on them. Will fix this bug once I have a chance to run through the `Go` documentation.

## DONATE

If you find this utility useful and would like to buy us some coffee or beer you can send our coin our way:

- *BTC*:  `1NQXNGmAwWmkNEoqCSXBPYmxcbgkQbXNCM`
- *SC*: `adbc875c136fa6a39d37c146c151280b759d7a5462b4f8b734711d00b6168bb2f0742b3cf18b`
- *ETH*: `0x9289E1B3a846061B14E81FD521a3ac910EC908A7`

## LICENSE

The MIT License (MIT)
