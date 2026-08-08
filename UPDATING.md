# Updating gotk4-adwaita

This guide details the procedure for updating `gotk4-adwaita` to newer Libadwaita and Nixpkgs versions.

_Assumes you're inside `nix develop` already._

1. **Update Flake Inputs & Go Dependencies**
   - Bump `github:NixOS/nixpkgs`'s ref in `flake.nix` (e.g., `github:NixOS/nixpkgs?ref=nixos-25.05` -> `github:NixOS/nixpkgs?ref=nixos-26.05`).
   - Run `nix flake update && go get -u` to update the Nix flake inputs and Go module dependencies.

2. **Run Code Generation**
   - Run `go generate` inside the Nix development shell:
     ```sh
     go generate
     ```
   - Resolve any codegen errors or enum member collisions by adding preprocessors or filters in `generator.go`.

3. **Validate Package Compilation**
   - Build all generated packages in `pkg/`:
     ```sh
     (cd pkg && go build -v ./...)
     ```
   - If CGo compilation errors occur due to generated code, update generator preprocessors or postprocessors in `generator.go`, then re-run `go generate` and `go build`.

4. **Run Tests**
   - Validate tests pass across the repository:
     ```sh
     go test ./...
     cd pkg && go test ./...
     ```

5. **Document Changes in `CHANGELOG.md`**
   - Create or update `CHANGELOG.md` with a release section for the target version (e.g., `## [1.7.0] - YYYY-MM-DD`).
   - Follow `gotk4`'s changelog format with the following sections:
     - **Version Upgrades**: Record Libadwaita version bump, Go toolchain version, Nixpkgs channel/commit, and Go dependency updates.
     - **GIR Code Generation & Generator Fixes**: Document preprocessor/postprocessor rules or generator adjustments in `generator.go`.
     - **New APIs Overview**: Summarize new Libadwaita widgets, classes, methods, and constants in `pkg/adw`, preserving exact Go casing.

6. **Branch Management**
   - Identify the Libadwaita minor version from `pkg-config --modversion libadwaita-1` (e.g., `1.7.0` -> minor version `7`).
   - Create or update branch `adw-1.<minor>` (e.g., `adw-1.7`).
