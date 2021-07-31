self: super: {
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
}
