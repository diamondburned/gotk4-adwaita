# gotk4-adwaita

[Package documentation](https://pkg.go.dev/github.com/diamondburned/gotk4-adwaita/pkg/adw)

Go generated bindings for [Libadwaita][adw].

## Update Bindings

Currently, these bindings are built against the version `1.2.0` of Libadwaita.

The easiest way to upgrade the bindings is to fork this repository and use Nix
to regenerate the bindings. Adjust the Nix specific files, [as needed][nix-overlay].

For completeness, instructions for Flatpak as well as a manual way are provided.

### Nix

For building with Nix, please refer to the [gotk4 documentation][gotk4-contributing].

### Flatpak

You should follow these instructions, if you plan to package your application
as a Flatpak.

This way, you can generate bindings for the version of Libadwaita,
provided by your chosen runtime or for the newest version of Adwaita.

To generate the bindings, do the following:

```sh
$ flatpak-builder --force-clean build-dir io.github.diamondburned.gotk4_adwaita.yml
$ flatpak-builder --run build-dir io.github.diamondburned.gotk4_adwaita.yml go generate
```

You might need to install the referenced SDKs and runtimes, first.
You might want to read the [Flatpak building introduction][flatpak].

### Manual

You should use the manual way, if you do want to use the version of Libadwaita,
your system provides.

If you do not want to install all the libraries onto your system, you can use
a tool like [Toolbox][toolbox].

Optionally enter your toolbox via `toolbox enter` and run the following command
afterwards.

```sh
$ go generate
```

Install the missing system-libraries, until the generation succeeds.

[adw]: https://gnome.pages.gitlab.gnome.org/libadwaita/doc/
[flatpak]: https://docs.flatpak.org/en/latest/building-introduction.html
[gotk4-contributing]: https://github.com/diamondburned/gotk4/blob/4/CONTRIBUTING.md
[nix-overlay]: https://github.com/diamondburned/gotk4-adwaita/issues/3#issuecomment-1304912311
[toolbox]: https://containertoolbx.org/
