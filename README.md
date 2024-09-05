Ebanx Coding Challenge


How to run 
$ go run main.go


GET /balance
checks balance for account. Must contain querystring i.e. '?id=XXXX'. returns 404 0 if account does not exist

e.g.
GET /balance?id=1234
output

200 20

POST /event
Handles Banking Operations for accounts

Event type descritions:
Deposit
Adds funds to account, must countain destination(accountId) and amount. Creates account if account does not exist.
e.g.

"type":"deposit", "destination":"100", "amount":10

output

201 {"destination": {"id":"100", "balance":10}}

Withdraw
Withdraws funds from account, must countain origin(accountId) and amount. returns 404 0 if account does not exist. returns 404 0 if account does not exist or funds are above balance.
e.g.
"type":"withdraw", "origin":"200", "amount":10

output 

201 {"origin": {"id":"100", "balance":15}}

Transfer
Transfers funds between accounts, must countain  origin(accountId), destination(accountId) and amount. returns 404 0 if origin account does not exist or funds are above balance. Creates destination account if it does not exist.
"type":"transfer", "origin":"100", "amount":15, "destination":"300"

output

201 {"origin": {"id":"100", "balance":0}, "destination": {"id":"300", "balance":15}}


POST /reset
clears existing entries

output
200 OK
