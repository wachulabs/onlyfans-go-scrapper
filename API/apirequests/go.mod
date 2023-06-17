module github.com/mwaurawakati/onlyfans/API/apirequests

go 1.18

replace github.com/mwaurawakati/onlyfans/API/auth => .././auth

replace github.com/mwaurawakati/onlyfans/API/endpointlinks => .././endpointlinks

require (
	github.com/mwaurawakati/onlyfans/API/auth v0.0.0-00010101000000-000000000000
	github.com/mwaurawakati/onlyfans/API/endpointlinks v0.0.0-00010101000000-000000000000
)

require github.com/pkg/errors v0.9.1
