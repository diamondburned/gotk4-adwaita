self: super: {
	libadwaita = super.libadwaita.overrideAttrs (old: {
		version = "1.0.0-alpha.3";

		src = super.fetchFromGitLab {
			domain = "gitlab.gnome.org";
			owner  = "GNOME";
			repo   = "libadwaita";
			rev    = "40c19ab2591763a482ebc79c82f1da32eea3bab6";
			hash   = "sha256:1bfxsq6sm0xpp5z4q2h2qvwkkbga2ryr0ny1wvp6ccarr2dzch70";
		};

		doCheck = false;
	});
}
