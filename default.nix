{
  version,
  buildGoApplication,
  lib,
}: let
  fs = lib.fileset;
in
  buildGoApplication {
    pname = "aterm2json";
    inherit version;
    pwd = ./.;
    src = fs.toSource {
      root = ./.;
      fileset =
        fs.intersection
        (fs.gitTracked ./.)
        (fs.unions [
          ./cmd
          ./go.mod
          ./go.sum
        ]);
    };

    modules = ./gomod2nix.toml;

    meta = {
      description = "A tool to convert between aterm and json.";
      homepage = "https://github.com/fzakaria/aterm2json";
      mainProgram = "aterm2json";
    };
  }
