# Changelog

All notable changes to `gotk4-adwaita` will be documented in this file.

## [1.9.2] - 2026-08-08

### Version Upgrades

- Bumped Libadwaita from `1.7.0` to `1.9.2`.
- Updated Go toolchain to `1.24.3`.
- Updated Nixpkgs channel ref to `nixos-26.05`.
- Upgraded `github.com/diamondburned/gotk4` dependency to `v0.4.0`.

### GIR Code Generation & Generator Fixes

- Re-generated Go bindings against Libadwaita 1.9 GIR definitions using `girgen`.
- Retained custom preprocessor rule `types.RenameCallable("Adw-1.Breakpoint.add_setter", "add_setter_direct")` and postprocessor `betterBreakpointAddSetter` for `Breakpoint.AddSetter`.

### New APIs Overview

#### Libadwaita 1.8

- `ShortcutLabel` (`AdwShortcutLabel`): Display formatted keyboard shortcuts (`NewShortcutLabel`, `Accelerator`, `SetAccelerator`, `DisabledText`, `SetDisabledText`).
- `ShortcutsDialog` (`AdwShortcutsDialog`): Dialog for keyboard shortcuts (`NewShortcutsDialog`, `AddSection`, `RemoveSection`, `SearchEnabled`, `SetSearchEnabled`).
- `ShortcutsItem` (`AdwShortcutsItem`): Item representing a keyboard shortcut inside a section (`NewShortcutsItem`, `Accelerator`, `SetAccelerator`, `Title`, `SetTitle`).
- `ShortcutsSection` (`AdwShortcutsSection`): Section grouping shortcut items (`NewShortcutsSection`, `Add`, `Remove`, `Title`, `SetTitle`).

#### Libadwaita 1.9

- `Sidebar` (`AdwSidebar`): Sidebar navigation container (`NewSidebar`, `AddSection`, `RemoveSection`, `SelectedItem`, `SetSelectedItem`).
- `SidebarItem` (`AdwSidebarItem`): Individual item in a sidebar (`NewSidebarItem`, `Title`, `SetTitle`, `IconName`, `SetIconName`).
- `SidebarSection` (`AdwSidebarSection`): Section grouping items in a sidebar (`NewSidebarSection`, `Add`, `Remove`, `Title`, `SetTitle`).
- `ViewSwitcherSidebar` (`AdwViewSwitcherSidebar`): Sidebar integrated with a view switcher stack (`NewViewSwitcherSidebar`, `Stack`, `SetStack`).
- `NoneAnimationTarget` (`AdwNoneAnimationTarget`): Dummy animation target (`NewNoneAnimationTarget`).
- `AboutDialog`: Added `AppdataResourcePath` method (`AppdataResourcePath`).
