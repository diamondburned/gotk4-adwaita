let systemPkgs = import <nixpkgs> {};

in {
	gotk4 ? (systemPkgs.fetchFromGitHub {
		owner = "diamondburned";
		repo  = "gotk4";
		rev   = "4f507c20f8b07f4a87f0152fbefdc9a380042b83";
		hash  = "sha256:0zijivbyjfbb2vda05vpvq268i7vx9bhzlbzzsa4zfzzr9427w66";
	}),

	pkgs ? (import "${gotk4}/.nix/pkgs.nix" {
		overlays = [ (import ./overlay.nix) ];
	}),
}:

let shell = import "${gotk4}/.nix/shell.nix" {
	inherit pkgs;
};

in shell.overrideAttrs (old: {
	buildInputs = old.buildInputs ++ (with pkgs; [
		libadwaita
	]);
})
