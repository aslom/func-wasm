# func-wasm

# Add yoour Docker prefix to trusted builders to allow func build to work

For details and context see:
https://github.com/knative/func/discussions/1551#discussioncomment-5220469

Buuild your won func

```
cd knative/func
make
```

###

```
func repositories add wasm https://github.com/aslom/func-wasm
mkdir wt1 && cd $_
func create -l tinygo -t wasm/http-wagi
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
