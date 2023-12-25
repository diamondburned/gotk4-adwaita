{ pkgs ? import <nixpkgs> {} }:

let
	gotk4-nix = pkgs.fetchFromGitHub {
		owner  = "diamondburned";
		repo   = "gotk4-nix";
		rev    = "ad91dabf706946c4380d0a105f0937e4e8ffd75f";
		sha256 = "0rkw9k98qy7ifwypkh2fqhdn7y2qphy2f8xjisj0cyp5pjja62im";
	};

in import "${gotk4-nix}/shell.nix" {
	base = {
		pname = "gotk4-adwaita";
		version = "dev";

		buildInputs = pkgs: with pkgs; [
			gobject-introspection
			glib
			graphene
			gdk-pixbuf
			gtk4
			gtk3
			vulkan-headers
			libadwaita
		];
	};
	pkgs = import "${gotk4-nix}/pkgs.nix" {
		sourceNixpkgs = pkgs.fetchFromGitHub {
			owner = "NixOS";
			repo  = "nixpkgs";
			rev   = "0b3d618173114c64ab666f557504d6982665d328"; # nixos-23.11
			hash  = "sha256-q5j06XOsy0qHOarsYPfZYJPWbTbc8sryRxianlEPJN0=";
		};
		useFetched = true;
	};
}
