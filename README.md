# presentation

My presentations for myself using [present](https://pkg.go.dev/golang.org/x/tools/present).

## Setup

Go and `present` are pinned in `shell.nix`:

```sh
nix-shell
```

## Serving

```sh
make serve
```

Open <http://127.0.0.1:3999>, not `localhost`: the Run button on code slides is rejected from any other origin. 

