Yet another Pet (or I just call it Pet-Enhanced). Since I haven't changed any of the original code, I will not change the name of the tool.
   
Much appreciation to `knqyf263` who made Pet. [Original Pet repo](https://github.com/knqyf263/pet)

It's a little command snippet management tool which can be used as a quick reference to concepts and commands in case you forgot. Pretty handy.

I added markdown support to the tool and personally use it with `glow`, a CLI markdown rendering tool to better format the output. Like this:

![image](https://user-images.githubusercontent.com/28176389/129502096-65807f4f-10cb-4db6-9bae-dc770283ef63.png)

I haven't tested all the markdown syntax, just `header`, `inline code block` and `table`. If markdown went wrong, please checkout the [glow doc](https://github.com/charmbracelet/glow).

NOTE: Please refer to known issues for caveats.

# Dependencies

`Pet` depends on [`glow`](github.com/charmbracelet/glow) to handle markdown syntax. And `feh` to handle images. Install them first. The default search program is `fzf`.

```
sudo apt install feh
sudo apt install fzf
go install github.com/charmbracelet/glow@latest
```

# Install

You can either build pet from source, or download the pre-compiled binary from release.

## Build from Source

1. Clone the repo to `GOPATH` on your system;
2. Issue command `go build`;
3. Move `pet` to `/usr/local/bin/pet`;
3. Profit!

## Download Pre-compiled Binary

Download the pre-compiled binary and extract to `/usr/local/bin/`;

# Usage

## Shell Function

Put the following scirpt in whatever shell's `.rc` file.

```txt
pets () {
    # proxychains can be remove according to the image service you are using
    # refer to image support in the below section
    cmd=`proxychains -q /usr/local/bin/pet search`
    if [[ "$cmd" == "#"* ]]
    then
        echo "$cmd" | glow -s dark - | less -r
        echo "$cmd" | glow -s dark -
    else
        echo "$cmd"
        echo "$cmd" | tr -d '\n' | clip.exe
    fi
}
```

Run `source xxx.rc`, then fire up `pets` to run the executable. Replace `clip.exe` (for WSL) to `xclip -sel c` for Linux.

## Use for Sinlge command

### Save Snippet

`n` is for `normal` mode.

![image](https://user-images.githubusercontent.com/28176389/208106824-b7ff653e-9642-411a-a2d0-aace7b278d14.png)

### Search for Snippet

Issue command `pets`, enter keywords to search for your snippet, locate the target snippet, then press `Enter`.

![image](https://user-images.githubusercontent.com/28176389/208107108-fa81d9e9-2cdd-49d9-90cd-7295adf86af4.png)

Note that texts in `<>` are arguments, pet will prompt you for input. Fill in each argument, then press `Enter`.

![image](https://user-images.githubusercontent.com/28176389/208107357-ef61740c-b8ce-4268-8ee6-1679e882175c.png)

Now, the command is in your clipboard, you can paste the command for execution.

![image](https://user-images.githubusercontent.com/28176389/208108345-cd2e5a1c-f1e0-4e2e-b704-378350686e55.png)

## Use for Markdown Mode (Input Line by Line)

### Save Snippet

`m` for `markdown` mode. `eof` exits input and start description.

![image](https://user-images.githubusercontent.com/28176389/208107945-a7d8c09d-e306-445a-91a2-6d21b66f9254.png)

### Search for Snippet

Issue command `pets`, enter keywords to search for your snippet, locate the target snippet, then press `Enter`.

![image](https://user-images.githubusercontent.com/28176389/208108100-0c75e346-725f-453d-b561-8c0be9a0839c.png)

The snippet will be displayed by `glow` in terminal.

![image](https://user-images.githubusercontent.com/28176389/208108146-75d42947-954a-4d16-a6e4-53da9d08a960.png)

## Use for Markdown File

### Save Snippet

Put your markdown in a file, save it.

![image](https://user-images.githubusercontent.com/28176389/208108800-ccf45361-cfef-446d-9260-b08e0e036840.png)

Use `pet new /path/to/file` to read and add Markdown content from the file.

![image](https://user-images.githubusercontent.com/28176389/208108916-22caa6d5-87b5-471f-a763-a9a6de720a30.png)

### Search for Snippet

Issue command `pets`, enter keywords to search for your snippet, locate the target snippet, then press `Enter`.

![image](https://user-images.githubusercontent.com/28176389/208109016-1534737f-de61-4885-9419-ecbb7a1338d4.png)

The snippet will be displayed by `glow` in terminal.

![image](https://user-images.githubusercontent.com/28176389/208109049-fc8343c3-ebaa-4ac4-b4b0-cc3cd4119b68.png)

## Use for Markdown File with Image

Upload your image to your image service, then copy the URL to the image, save it in pet using the following format:

`img::https://example.com/5Zhdgesz.png`

just like in the screenshot below.

![image](https://user-images.githubusercontent.com/28176389/208110386-d7d69d66-641e-4c68-a60e-ea2c4e11ffaa.png)

Then save the markdown file with image.

![image](https://user-images.githubusercontent.com/28176389/208111518-32c140e0-6354-4336-87e4-b03ba7556e58.png)

You can add more than one image to a file, one on each line:

```txt
img::https://exmaple.com/img1.png
img::https://exmaple.com/img2.png
```

# Relese Notes

## Latest (v0.9)

### Support for Images

Image support added. I personally upload image to imgur, and add in the snippet the keyword `img::` followed by the image link, eg:

```
img::https://i.imgur.com/V3vsK.png
```

* REQUIREMENTS: Install `feh`

```
sudo apt install feh
```

And if you have `feh` installed, when you search for that piece of snippet, then `feh` will automatically display the image, which makes my wiki tools more powerful.
![2022-01-04 19_48_02-pet_README md at main · reyalpmi_pet](https://user-images.githubusercontent.com/28176389/148054694-b294b7a5-1517-4784-b0d2-0b89457ee9f1.png)

### Support for Saving Markdown as File

Ability to read markdown content from file is added on Jan, 12th, 2022. Now, for simple markdown snippets, you can fire up pet and input line by line. But for longer snippets with bunch of codes, just put the whole markdown in a file, and issue command

```
pet new /path/to/file
```

to add the content to pet.

Enjoy!

# Known Issues

1. Keep Markdown as simple as possible, otherwise `glow` will not display them correctly. The following markdowns have been tested:

   ```txt
   Header 1: # Title
   Inline code: `command`
   Horizontal Sep: ----------
   Code block: ``` code here ```
   Table:
   |Col1|Col2|
   |----|----|
   |val1|val2|
   ```

   Example Markdown:

   ````txt
   # This Is the Title

   -----------

   Steps to download file from a URL in command line.

   1. Open cmd;
   2. Issue the following command:

   ```
   powershell wget https://example.com/fancy.png -o fancy.png
   ```
   ````

   Remember, markdown mode snippets always starts with `#` as first character.

2. When not in Markdown mode (snippets that starts with #), angle brackets (`<` and `>`) will be interpreted as arguments needed to command, so pet will prompt you for input. This is not cool for saving PHP oneliner. So, either save PHP oneliner as markdown, or just LEAVE OUT the `>` at the end, and add that after pasting the command.

3. When prompted for argument, try not to paste in very long text, the text will get messed up somehow. You can try for yourself. This is not going to be fixed in a while, so, bear with me.

# Bug Reports

Feel free to open issue.

# License

MIT

# Credit

Again, thanks to Teppei Fukuda (knqyf263) who made Pet Snippet manager. [Original Pet repo](https://github.com/knqyf263/pet)
