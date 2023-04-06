# Template to run Wasm WASI WAGI with Knative functions



# Prerequisites

Install [Knative functions](https://knative.dev/docs/functions/install-func/), make sure it is at least version 1.9.1

```
func version​
v1.9.1​
```

# Install Wasm func repository

```
func repositories add wasm https://github.com/aslom/func-wasm
```

# Run sample code

```
mkdir wt1 
cd wt1
func create -l tinygo -t wasm/http-wagi
```

Edit main.go file.

Deploy

```
func deploy -v --registry docker.io/aslom
```

# Local testing

Build:

```
func build -r docker.io/aslom
```

Start the code:

```
docker run --rm -p 8080:8080 --entrypoint web docker.io/aslom/wt1:latest
```

Quick test in web browser by opening http://localhost:8080/

# Context:

https://github.com/knative/func/discussions/1551


