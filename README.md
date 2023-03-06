# Template to run Wasm WASI WAGI with Knative func


Context:
https://github.com/knative/func/discussions/1551#discussioncomment-5220469

# If you are going to use your own build packs

Change everything from docker.io/aslom to your Docker prefix

Use your own docker prefix and rebuiild https://github.com/aslom/knative-wasm-buildpacks

Modify files inside to use Docker prefix

# Add Docker prefix to trusted builders to allow func build to work

For details and context see:
https://github.com/knative/func/discussions/1551#discussioncomment-5220469

Issue explaining it:
https://github.com/knative/func/issues/1599

Clone knative/func and modify ` pkg/builders/buildpacks/builder.go` to add "docker.io/aslom/" to trusted builders prefixes.


Build your own func

```
cd knative/func
make
```

# Run sample code

```
func repositories add wasm https://github.com/aslom/func-wasm
mkdir wt1 && cd $_
func create -l tinygo -t wasm/http-wagi
```

Edit main.go

Build:

```
$YOUR_VERSION_OF_KNATIVE_FUNC/func build -r docker.io/aslom
```

Quick test localhost:8080/

```
docker run --rm -p 8080:8080 --entrypoint web docker.io/aslom/wt1:latest
```

Deploy

```
$YOUR_VERSION_OF_KNATIVE_FUNC/func deploy -v --registry docker.io/aslom
```
