# Convertimg
The CLI application resize images and convert to other formats.

## Build/Install
### Requirements:
- Go^1.20

Execute command:
```go build main.go```
So, you may have a executable program named *main* (Linux/Mac) or *main.exe* (Windows).
Rename the file to *convertimg* (Linux/Mac) or *convertimg.exe* (Windows)

And so, you have the app builded, but if you want execute the command globaly on your enviroment, you must configure the directory of executable on enviroment variable called `PATH`.
[How to set a command on PATH Windows](https://www.computerhope.com/issues/ch000549.htm)
[How to set a command on PATH Linux](https://phoenixnap.com/kb/linux-add-to-path)

## How to use
#### Converting a image
For convert any image from any format you use command:
<code>convertimg \<filepath> [-f JPEG|PNG|BMP|GIF|TIFF]</code>
- **filepath**: the argument `filepath` is the full path to image who you want convert.
- **-f**: the flag `-f` is to especificate the output format. Accept values are JPEG, PNG, BMP, GIF and TIFF. *The default is JPEG.*

Other parameters are:
- **-o**: the flag accept a output filepath to new image generated.
- **-w**: define the `width` in pixels to output image.
- **-h**: define the `height` in pixels to output image.
	*Note: if specified only one of arguments `-h` e `-w` the other dimension out resized mantaning aspect.*

**Example:**
```convertimg /home/image.png -f JPEG -o /home/new-image.jpg -w 300 -h 280```

#### Return the base64 encoded image
You can use the flag `--base64` to get a base64 encoded image, like it:
```convertimg /home/image.jpeg --base64```

Output:
```data:image/jpeg;base64,RXhhbXBsZSBiYXNlNjQ=```

***Note**: this flag can't be used in convert mode.*
