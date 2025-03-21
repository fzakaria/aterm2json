{
  description = "A tool to convert between aterm and json.";

  inputs = {
    nixpkgs.url = "github:nixos/nixpkgs/nixos-24.11";
    systems.url = "github:nix-systems/default";
    gomod2nix = {
      url = "github:nix-community/gomod2nix";
      inputs.nixpkgs.follows = "nixpkgs";
    };
  };

  outputs = {
    self,
    systems,
    nixpkgs,
    ...
  } @ inputs: let
    eachSystem = f:
      nixpkgs.lib.genAttrs (import systems) (system:
        f {
          pkgs = import nixpkgs {
            overlays = [
              (final: _prev: {
                aterm2json = final.callPackage ./default.nix {
                  # https://github.com/nix-community/nixos-facter/blob/906098c600609d95a475449272d59b68bda2ef83/nix/packages/nixos-facter/default.nix#L18
                  # there's no good way of tying in the version to a git tag or branch
                  # so for simplicity's sake we set the version as the commit revision hash
                  # we remove the `-dirty` suffix to avoid a lot of unnecessary rebuilds in local dev
                  version = final.lib.removeSuffix "-dirty" (self.shortRev or self.dirtyShortRev);
                };
              })
              (inputs.gomod2nix.overlays.default)
            ];
            inherit system;
          };
          inherit system;
        });
  in {
    formatter = eachSystem ({pkgs, ...}: pkgs.alejandra);

    packages = eachSystem ({pkgs, ...}: {
      default = pkgs.aterm2json;
    });

    devShells = eachSystem ({pkgs, ...}:
      with pkgs; {
        default = mkShell {
          inputsFrom = [
            aterm2json
          ];

          packages = [
            (mkGoEnv {pwd = ./.;})
            gomod2nix
          ];
        };
      });
  };
}
