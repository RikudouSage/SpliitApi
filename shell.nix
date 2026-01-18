{ pkgs ? import <nixpkgs> {} }:
pkgs.mkShell {
    nativeBuildInputs = with pkgs.buildPackages;
    [
      pkgs.golangci-lint
      pkgs.revive
    ];
}
