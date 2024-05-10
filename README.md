## Google Chrome cache attack

This proof of concept project is an example for the Google Chrome cache attack. It can inject custom code into web pages by modifying the browser cache files. The demo app doesn't require root privileges to run. The browser won't notice that the web page (retrieved from the cache) was infected, for HTTPS content it'll display that the connection is secure and the certificate is valid. Also it's possible to set a different cache expiration time (1 day, month or even 1 year). 

`Disclaimer: This tool is only intended for security research. Users are responsible for all legal and related liabilities resulting from the use of this tool. The original author does not assume any legal responsibility.`


The demo app will inject [this html file](/test-files/test.html) for all page urls listed in [main.go](/main.go):

https://github.com/dneprDroid/chrome-poc-demo/assets/13742733/d1029f26-0bb9-463e-893b-82f106d4825e

Tested browser version: 124.0.6367.62 and older. 

Supported platforms:

- macOS
- Linux 

**NOTE**: For Windows, Google Chrome uses a different cache format. It's possible to run this attack on this platform, but it requires to implement **the block file caching**. 

### Build 

Build the project:

```bash 
go mod download

CGO_ENABLED=1 go build -o chrome-poc

```

Using Docker:

```bash 
# build a Linux executable
docker build --platform=linux/amd64 -t chrome-poc-linux:latest -f Dockerfile.linux .

```

```bash 
# cross-compile a macOS executable via osxcross
docker build --platform=linux/amd64 -t chrome-poc-osxcross:latest -f Dockerfile.osxcross . 

```
