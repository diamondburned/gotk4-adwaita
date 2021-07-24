{ systemPkgs ? import <nixpkgs> {} }:

let gotk4_src = systemPkgs.fetchFromGitHub {
		owner = "diamondburned";
		repo  = "gotk4";
		rev   = "5628639f0e5d43f53a16b0416fbc80e22c839d75";
		hash  = "sha256:0nq6iinmksfndfnk9i4xnb7y0vrxzmg1lki4l2dlia6pn1r5xw3y";
	};

	gotk4 = import "${gotk4_src}/shell.nix" {
		inherit systemPkgs;
	};

	libadwaita = let pkgs = gotk4.pkgs;
	in pkgs.libadwaita.overrideAttrs (old: {
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
	});

in gotk4.overrideAttrs (old: {
	buildInputs = old.buildInputs ++ [ libadwaita ];
}) // {
	# Expose libadwaita for external use.
	inherit libadwaita;
	inherit (gotk4) pkgs;
}
