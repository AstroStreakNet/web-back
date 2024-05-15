{ pkgs ? import <nixpkgs> {} }:

pkgs.mkShell {
    buildInputs = [
        pkgs.go
        pkgs.postgresql
        pkgs.git 
    ];

    # Set up fresh environment
    shellHook = ''
        # remove old go.mod and go.sum
        # ducttape fix to ensure interpreter stuff is visible to the syscall
        # since nix is not the primary dev env for the project
        rm go.mod go.sum
        go mod init webback
        go mod tidy
        go run ./main.go
    '';
}
