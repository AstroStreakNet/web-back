{ pkgs ? import <nixpkgs> {} }:

pkgs.mkShell {
  buildInputs = [
    pkgs.go
    pkgs.postgresql
    pkgs.git 
  ];

  # Set up the GOPATH environment variable
  shellHook = ''
    # ensure go modules are enabled
    export GO111MODULE=on
    go mod tidy
  '';
}
