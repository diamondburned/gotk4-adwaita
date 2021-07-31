let systemPkgs = import <nixpkgs> {};

in {
	gotk4 ? (systemPkgs.fetchFromGitHub {
		owner = "diamondburned";
		repo  = "gotk4";
		rev   = "4f507c20f8b07f4a87f0152fbefdc9a380042b83";
		hash  = "sha256:0zijivbyjfbb2vda05vpvq268i7vx9bhzlbzzsa4zfzzr9427w66";
	}),

	pkgs ? (import "${gotk4}/.nix/pkgs.nix" {}),

	libadwaita ? (pkgs.libadwaita.overrideAttrs (old: {
		version = "1.0.0-alpha.2";

		src = pkgs.fetchFromGitLab {
			domain = "gitlab.gnome.org";
			owner  = "GNOME";
			repo   = "libadwaita";
			rev    = "f5932ab4250c8e709958c6e75a1a4941a5f0f386";
			hash   = "sha256:1yvjdzs5ipmr4gi0l4k6dkqhl9b090kpjc3ll8bv1a6i7yfaf53s";
		};

		buildInputs = old.buildInputs ++ (with pkgs; [
			fribidi
		]);

		nativeBuildInputs = old.nativeBuildInputs ++ (with pkgs; [
			cmake
			gi-docgen
		]);

		doCheck = false;
		outputs = [ "out" "dev" ];
		mesonFlags = [ "-Dtests=false" "-Dgtk_doc=false" ];
	})),
}:

let shell = import "${gotk4}/.nix/shell.nix" {};

in shell.overrideAttrs (old: {
	inherit libadwaita;
	buildInputs = old.buildInputs ++ (with pkgs; [ libadwaita ]);
})
