{
  description = "Alunya";

  inputs.nixpkgs.url = "github:nixos/nixpkgs/nixos-unstable";

  outputs =
    {
      self,
      nixpkgs,
    }:
    let
      system = "x86_64-linux";
      pkgs = nixpkgs.legacyPackages.${system};
    in
    {
      devShells.${system}.default = pkgs.mkShell {
        packages = with pkgs; [
          go
          gopls
          gofumpt
          golangci-lint

          vscode-langservers-extracted
          prettier

          mbake
          checkmake
          gnumake

          pkg-config
          libx11.dev
          libxcursor
          libxi
          libxinerama
          libxrandr
          libxxf86vm
          libxkbcommon
          wayland
        ];
      };
    };
}
