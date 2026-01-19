{
  description = "Dev shell for CGO cross-compilation (386, armv7, arm64)";

  inputs = {
    nixpkgs.url = "nixpkgs";
  };

  outputs = { self, nixpkgs }:
    let
      system = "x86_64-linux";
      pkgs = import nixpkgs { inherit system; };

      cc386   = pkgs.pkgsCross.gnu32.stdenv.cc;
      ccArmv7 = pkgs.pkgsCross.armv7l-hf-multiplatform.stdenv.cc;
      ccArm64 = pkgs.pkgsCross.aarch64-multiplatform.stdenv.cc;
    in {
      devShells.${system}.default = pkgs.mkShell {
        buildInputs = [
          pkgs.go
          pkgs.golangci-lint
          pkgs.revive

          cc386
          ccArmv7
          ccArm64
        ];

        shellHook = ''
          unset GOROOT

          export CC_386=${cc386.targetPrefix}cc
          export CC_ARMV7=${ccArmv7.targetPrefix}cc
          export CC_ARM64=${ccArm64.targetPrefix}cc

          echo "Using cross compilers:"
          echo "  CC_386   = $CC_386"
          echo "  CC_ARMV7 = $CC_ARMV7"
          echo "  CC_ARM64 = $CC_ARM64"

          echo ""

          echo "Using go config:"
          echo "  GOROOT   = $(go env GOROOT)"
          echo "  GOCACHE  = $(go env GOCACHE)"
          echo "  GOPATH   = $(go env GOPATH)"
        '';
      };
    };
}
