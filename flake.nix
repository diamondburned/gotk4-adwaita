{
  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs?ref=nixos-unstable";
    nixpkgs-gotk4.url = "github:NixOS/nixpkgs?ref=nixos-24.05";
    gotk4-nix.url = "github:diamondburned/gotk4-nix/main";
    flake-utils.url = "github:numtide/flake-utils";
  };

  outputs =
    {
      self,
      nixpkgs,
      nixpkgs-gotk4,
      gotk4-nix,
      flake-utils,
    }:

    flake-utils.lib.eachDefaultSystem (
      system:
      let
        pkgs = nixpkgs.legacyPackages.${system};
        pkgs-gotk4 = import nixpkgs-gotk4 {
          inherit system;
          overlays = [ gotk4-nix.overlays.patchelf ];
        };
      in
      {
        devShells.default = gotk4-nix.lib.mkShell {
          base.pname = "gotk4-adwaita";
          pkgs = pkgs-gotk4;

          buildInputs = with pkgs-gotk4; [
            gobject-introspection
            glib
            graphene
            gdk-pixbuf
            gtk4
            gtk3
            vulkan-headers
            libadwaita
          ];

          nativeBuildInputs = [ pkgs.go ];

          packages = with pkgs; [ self.formatter.${system} ];
        };

        formatter = pkgs.nixfmt-rfc-style;
      }
    );
}
