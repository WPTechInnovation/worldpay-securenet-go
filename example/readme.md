A sample program that tokenizes some payment card credentials then charges the payment card using the token.

API credentials are provided via command line flags. The `-proxy` flag can be passed if a HTTP proxy is to be used and should be of the format `scheme://host:port` for example `http://example.com:8080`

Usage:

`example -publickey <public-key> -securekey <secure-key> -securenetid <secure-net-id>`
