{ systemPkgs ? import <nixpkgs> {} }:

let unstable = import (systemPkgs.fetchFromGitHub {
	owner  = "NixOS";
	repo   = "nixpkgs";
	rev    = "fbfb79400a08bf754e32b4d4fc3f7d8f8055cf94";
	sha256 = "0pgyx1l1gj33g5i9kwjar7dc3sal2g14mhfljcajj8bqzzrbc3za";
}) {
	overlays = [
		(self: super: {
			go = super.go.overrideAttrs (old: {
				version = "1.17beta1";
				src = builtins.fetchurl {
					url    = "https://golang.org/dl/go1.17rc1.linux-arm64.tar.gz";
					sha256 = "sha256:0kps5kw9yymxawf57ps9xivqrkx2p60bpmkisahr8jl1rqkf963l";
				};
				doCheck = false;
			});
			libadwaita = super.libadwaita.overrideAttrs (old: {
				version = "1.0.0-alpha.2";
				src = super.fetchFromGitLab {
					domain = "gitlab.gnome.org";
					owner  = "GNOME";
					repo   = "libadwaita";
					rev    = "f5932ab4250c8e709958c6e75a1a4941a5f0f386";
					hash   = "sha256:1yvjdzs5ipmr4gi0l4k6dkqhl9b090kpjc3ll8bv1a6i7yfaf53s";
				};
				buildInputs = old.buildInputs ++ (with super; [
					fribidi
				]);
				nativeBuildInputs = old.nativeBuildInputs ++ (with super; [
					cmake
					gi-docgen
				]);
				doCheck = false;
				outputs = [ "out" "dev" ];
				mesonFlags = [ "-Dtests=false" "-Dgtk_doc=false" ];
			});
		})
	];
};

in unstable.mkShell {
	buildInputs = with unstable; [
		# gotk4
		gobjectIntrospection
		glib
		graphene
		gdk-pixbuf
		gnome3.gtk
		gtk4
		vulkan-headers

		# adwaita
		libadwaita
	];

	nativeBuildInputs =	with unstable; [
		pkgconfig
		go
	];

	CGO_ENABLED = "1";

	TMP    = "/tmp";
	TMPDIR = "/tmp";
}
