# go-config-binsize

Quick comparison of binary sizes with different config language dependencies. Includes HCL, JSON, TOML, YAML.

Build environment:

* `CGO_ENABLED = 0`
* `GOFLAGS = -trimpath -ldflags="-s -w"`

## GOOS=darwin

```
$ du -h **/*-darwin
4.0M    hcl/bin/hcl-darwin
1.7M    json/bin/json-darwin
5.4M    jsonnet/bin/jsonnet-darwin
1.7M    toml/bin/toml-darwin
2.0M    yaml/bin/yaml-darwin
```

## GOOS=linux

```
$ du -h **/*-linux
3.8M    hcl/bin/hcl-linux
1.6M    json/bin/json-linux
4.7M    jsonnet/bin/jsonnet-linux
1.6M    toml/bin/toml-linux
1.9M    yaml/bin/yaml-linux
```

## GOOS=windows

```
$ du -h **/*-windows
3.7M    hcl/bin/hcl-windows
1.6M    json/bin/json-windows
4.9M    jsonnet/bin/jsonnet-windows
1.6M    toml/bin/toml-windows
1.9M    yaml/bin/yaml-windows
```