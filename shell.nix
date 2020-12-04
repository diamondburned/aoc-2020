{ pkgs ? import <nixpkgs> {} }:

let	zig = pkgs.zig.overrideAttrs (old: {
		version = "0.7.0";

		buildInputs = with pkgs; [
			llvmPackages_11.clang-unwrapped
			llvmPackages_11.llvm
			llvmPackages_11.lld
			libxml2
			zlib
		];

		doCheck = false;

		src = pkgs.fetchFromGitHub {
			owner  = "ziglang";
			repo   = "zig";
			rev    = "0c90ccc297e18f1a40a8111606a419e99732931f";
			sha256 = "0gq8xjqr3n38i2adkv9vf936frac80wh72dvhh4m5s66yafmhphg";
		};
	});

	zls = pkgs.stdenv.mkDerivation rec {
		name = "zls";
		version = "0.1.0";
	
		src = pkgs.fetchFromGitHub {
			owner  = "zigtools";
			repo   = "zls";
			rev    = "bf4f653bf985dba3e3849ba133848c37721d0445";
			sha256 = "0kgz5675k6x1i0g524z0sjbm2bg38yycqcr75cg1bi28d5v8x1h3";
			fetchSubmodules = true;
		};
	
		buildInputs = [ zig ];
	
		preConfigure = ''
			export XDG_CACHE_HOME=$(mktemp -d)
		'';
		buildPhase = ''
			zig build -Drelease-fast
		'';
		installPhase = ''
			mkdir $out
			mv zig-cache/bin $out
		'';
	};

in pkgs.mkShell {
	buildInputs = [ zls zig ];
}
