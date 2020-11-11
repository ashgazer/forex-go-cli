# forex-go-cli
First Command line app written in go that calls API and returns foreign exchange rates

To build from source please run the following command this will create the command line tool forex-cli 

`go build -o build/forex-cli `



How to run once built

See conversion rate

`./forex-cli -b GBP -d TRY`

See conversation rate plus amount that is exchanged (-c)
`./forex-cli -b GBP -d TRY -c 97`


# Flags

* -d Three letter destination currency code
* -b Three letter base currency you wish to convert from
* -c [optional] The value of money you wish to convert

if you have go installed you can run by going into the root folder.

`go run *.go -b GBP -d TRY -c 97`
