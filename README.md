# aterm2json

[![built with nix](https://builtwithnix.org/badge.svg)](https://builtwithnix.org)

A simple tool to convert between [nix](https://nixos.org) aterm format and JSON.


If you are using [Nix](https://nixos.org/), you can easily run this locally.

```sh
nix shell github:fzakaria/aterm2json
```

You can round trip this!

```console
> cat /nix/store/z3hhlxbckx4g3n9sw91nnvlkjvyw754p-myname.drv
Derive([("out","/nix/store/40s0qmrfb45vlh6610rk29ym318dswdr-myname","","")],[],[],"mysystem","mybuilder",[],[("builder","mybuilder"),("name","myname"),("out","/nix/store/40s0qmrfb45vlh6610rk29ym318dswdr-myname"),("system","mysystem")])

> aterm2json /nix/store/z3hhlxbckx4g3n9sw91nnvlkjvyw754p-myname.drv | jq | json2aterm -

Derive([("out","/nix/store/40s0qmrfb45vlh6610rk29ym318dswdr-myname","","")],[],[],"mysystem","mybuilder",[],[("builder","mybuilder"),("name","myname"),("out","/nix/store/40s0qmrfb45vlh6610rk29ym318dswdr-myname"),("system","mysystem")])
```