let systemPkgs = import <nixpkgs> {};

	nixpkgs = builtins.fetchTarball {
		name   = "gotk4-nixpkgs";
		url    = "https://github.com/NixOS/nixpkgs/archive/8d5200b4155a6fc505d8d79c6c845385b25efd77.tar.gz";
		sha256 = "1mphrs9d5pcxwfb3l74v8yvfi02qmn6i2ra4f8kbna1amiaa5c8h";
	};

in {
	gotk4 ? (systemPkgs.fetchFromGitHub {
		owner = "diamondburned";
		repo  = "gotk4";
		rev   = "318362ceedaf5281a5afbb8c65888e958c328ef9";
		hash  = "sha256:1f4sds78mnim36ymqsbqw07r6fi035ssj346krdnfqj75an4d5hz";
	}),

	pkgs ? (import "${gotk4}/.nix/pkgs.nix" {
		src = nixpkgs;
	}),
}:

pkgs.mkShell {
	name = "gotk4-adwaita-shell";

	buildInputs = with pkgs; [
		gobject-introspection
		glib
		graphene
		gdk-pixbuf
		gtk4
		gtk3
		vulkan-headers
		libadwaita
	];

	nativeBuildInputs = with pkgs; [
		pkgconfig
		go
	];

	CGO_ENABLED = "1";

	TMP    = "/tmp";
	TMPDIR = "/tmp";
}
