# Template to run Wasm WASI WAGI with Knative func


Context:
https://github.com/knative/func/discussions/1551#discussioncomment-5220469

# Prerequisites

Install (Knative func)[https://knative.dev/docs/functions/install-func/], make sure it is at least version 1.9.1

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
mkdir wt1 && cd $_
func create -l tinygo -t wasm/http-wagi
```

Edit main.go file.

Build:

```
func build -r docker.io/aslom
```

Quick test localhost:8080/

```
docker run --rm -p 8080:8080 --entrypoint web docker.io/aslom/wt1:latest
```

Deploy

```
func deploy -v --registry docker.io/aslom
```
