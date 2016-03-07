[![GoDoc](https://godoc.org/github.com/garlsecurity/securepassctl?status.svg)](https://godoc.org/github.com/garlsecurity/securepassctl)
[![Travis-CI 
Status](https://api.travis-ci.org/garlsecurity/securepassctl.png?branch=master)](http://travis-ci.org/#!/garlsecurity/securepassctl)
[![Coverage](http://gocover.io/_badge/github.com/garlsecurity/securepassctl)](http://gocover.io/github.com/garlsecurity/securepassctl)

# securepassctl

Go (golang) port of the SecurePass tool

# Usage
```console
$ spctl -h
Usage: spctl [global options] command [command options] [arguments...]
Manage distributed identities.
  
  --debug, -D	enable debug output
  --help, -h	show help
  --version, -v	print the version
  
Commands:
    ping		ping a SecurePass's remote endpoint
    app			manage applications
    config		configure SecurePass
    group-member	test group membership
    logs		display SecurePass logs
    radius		manage RADIUS information
    realm		manage realm settings
    user		manage users
    help, h		Shows a list of commands or help for one command
    

spctl home page: <https://github.com/garlsecurity/securepassctl>
SecurePass online help: <http://www.secure-pass.net/integration-guides-examples/>
Report bugs to <https://github.com/garlsecurity/securepassctl/issues>
```
