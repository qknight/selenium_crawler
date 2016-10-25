{ pkgs ? import <nixpkgs>{} } :
let 
  stdenv = pkgs.stdenv;
  buildGoPackage = pkgs.buildGoPackage;
  fetchgit = pkgs.fetchgit;
  fetchhg = pkgs.fetchhg;
  fetchbzr = pkgs.fetchbzr; 
  fetchsvn = pkgs.fetchsvn;
in

buildGoPackage rec {
  name = "robin-${version}";
  version = "20161024-${stdenv.lib.strings.substring 0 7 rev}";
  rev = "6159f49025fd5500e5c2cf8ceeca4295e72c1de5";

  goPackagePath = "/home/joachim/Desktop/projects/fablab/robin";

  src = ./.;

  goDeps = ./deps.nix;

  # TODO: add metadata https://nixos.org/nixpkgs/manual/#sec-standard-meta-attributes
  meta = {
  };
}
