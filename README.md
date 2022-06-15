# Steganographic Image Thing

## What it does

* Stores an image inside the 2 least significant bits of another image.
* Allows you to decode an image stored like this.

## Why might you use it

* Save disk space by storing multiple pictures in one???
* Sneakily send people pictures of cats disguised as pictures of other cats???
* Enhance the 2 least significant bits of non steganographically encoded images???

## How to use it

* Specify a mode. Either `encode` or `decode` with the `-m` flag,
* Specify a source file with the `-s` flag,
* Specify an output file with `-o` flag,
* If you are using the `encode` mode be sure to specify the image to be encoded with the `-a` flag,

## Sample commands with this tool.

Encoding `cat.png` into `tree.png` and saving the result as `definitelyNotACat.png`:

```
steganographisaurus -s="tree.png" -a="cat.png" -o="definitelyNotACat.png" -m="encode"
```
azure
decoding `definitelyNotACat.png` and saving the result as `out.png`:

```
steganographisaurus -s="out.png" -o="dec.png" -m="decode"
```