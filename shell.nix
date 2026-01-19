{ pkgs ? import <nixpkgs> {} }:
let
  cc386   = pkgs.pkgsCross.gnu32.stdenv.cc;
  ccArmv7 = pkgs.pkgsCross.armv7l-hf-multiplatform.stdenv.cc;
  ccArm64 = pkgs.pkgsCross.aarch64-multiplatform.stdenv.cc;
in
pkgs.mkShell {
    nativeBuildInputs = with pkgs.buildPackages;
    [
      pkgs.golangci-lint
      pkgs.revive

      cc386
      ccArmv7
      ccArm64
    ];

    shellHook = ''
      # These env vars are what your Makefile expects
      export CC_386=${cc386.targetPrefix}cc
      export CC_ARMV7=${ccArmv7.targetPrefix}cc
      export CC_ARM64=${ccArm64.targetPrefix}cc

      echo "CC_386   = $CC_386"
      echo "CC_ARMV7 = $CC_ARMV7"
      echo "CC_ARM64 = $CC_ARM64"
    '';
}
