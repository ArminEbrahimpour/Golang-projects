### caching proxy 

A lightweight HTTP proxy server with caching capabilities. This proxy intercepts HTTP requests, checks if the response is cached, and serves the cached response if available. If the response is not cached, it forwards the request to the origin server, caches the response, and serves it to the client.
Installation

  Prerequisites:
```
  Go 1.20 or higher installed on your machine.
```
  Clone the Repository:

```bash
  git clone https://github.com/your-username/caching-proxy.git
  cd caching-proxy
```
  Build the Project:

```bash
  go build -o caching-proxy
```
  Run the Proxy:

```bash
  ./caching-proxy --origin http://example.com
```
  Replace `http://example.com` with the URL of the origin server you want to proxy.

  
