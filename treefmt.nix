{
  projectRootFile = "flake.nix";

  # See https://github.com/numtide/treefmt-nix#supported-programs
  programs.gofmt.enable = true;

  # Markdown
  programs.mdformat.enable = true;

  # Nix
  programs.nixfmt.enable = true;
}
