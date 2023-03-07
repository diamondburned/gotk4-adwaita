{ pkgs ? import <nixpkgs> {} }:

let gotk4-nix = pkgs.fetchFromGitHub {
		owner = "diamondburned";
		repo  = "gotk4-nix";
		rev   = "c37b9a15d8117b5af60a0206d96a09ebf8e7d1ef";
		hash  = "sha256:05j49fbsqjvw7pqyqpc2g4w2q78n3vkgbbgkkfw2gbckbw7rydh5";
	};

	nixpkgs = pkgs.fetchFromGitHub {
		owner = "NixOS";
		repo  = "nixpkgs";
		rev   = "0dfdc09dba3ea6e873c5dbc81dc88577cab97979";
		hash  = "sha256:1j56ls1vp3syx9d0py60hlnjbrkg0hm3kv2q2qf8k7m4xwr5ckx7";
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
		sourceNixpkgs = nixpkgs;
		useFetched = true;
	};
}
