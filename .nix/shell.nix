let systemPkgs = import <nixpkgs> {};

in {
	gotk4 ? (systemPkgs.fetchFromGitHub {
		owner = "diamondburned";
		repo  = "gotk4";
		rev   = "318362ceedaf5281a5afbb8c65888e958c328ef9";
		hash  = "sha256:1f4sds78mnim36ymqsbqw07r6fi035ssj346krdnfqj75an4d5hz";
	}),

	pkgs ? (import "${gotk4}/.nix/pkgs.nix" {
		src = systemPkgs.fetchFromGitHub {
			owner = "NixOS";
			repo  = "nixpkgs";
			rev   = "3fdd780";
			hash  = "sha256:0df9v2snlk9ag7jnmxiv31pzhd0rqx2h3kzpsxpj07xns8k8dghz";
		};
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
