Yet another Pet (or I just call it Pet-Enhanced). Since I haven't changed any of the original code, I will not change the name of the tool.
   
Much appreciation to `knqyf263` who made Pet. [Original Pet repo](https://github.com/knqyf263/pet)

I personally use it with `glow`, a CLI markdown rendering tool to better format the output. Like this:

![image](https://user-images.githubusercontent.com/28176389/129502096-65807f4f-10cb-4db6-9bae-dc770283ef63.png)

I haven't tested all the markdown syntax, just `header`, `inline code block` and `table`. If markdown went wrong, please checkout the [glow doc](https://github.com/charmbracelet/glow).

Another thing, if you want to put angle brackets (`<>`) in your markdown text, remember to html encode them as `&lt;` and `&gt;` or use backtick to surround the angle bracket"`" when in a table.

# Install

Download the release and extract to whatever place in `PATH`, and issue the command:
```
pet
```

to use.

# Requirements

`Pet` depends on [`glow`](github.com/charmbracelet/glow) to handle markdown syntax. And `feh` to handle images. Install them first. The defualt search program is `fzf`.

```
sudo apt install feh
sudo apt install fzf
go install github.com/charmbracelet/glow@latest
```

# Usage

## Alias

Put the following scirpt in `.bashrc` or `.zshrc`.

```
pets () {
    cmd=`pet search`
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

Use `pets` to run the executable.

## Sinlge command

### Save Snippet

`n` is for `normal` mode.

![image](https://user-images.githubusercontent.com/28176389/208106824-b7ff653e-9642-411a-a2d0-aace7b278d14.png)

### Search for Snippet

Texts in `<>` are arguments, pet will prompt you for input.

![image](https://user-images.githubusercontent.com/28176389/208107108-fa81d9e9-2cdd-49d9-90cd-7295adf86af4.png)

![image](https://user-images.githubusercontent.com/28176389/208107357-ef61740c-b8ce-4268-8ee6-1679e882175c.png)

Then, you can paste the command for execution.

![image](https://user-images.githubusercontent.com/28176389/208108345-cd2e5a1c-f1e0-4e2e-b704-378350686e55.png)


## Markdown Mode

### Save Snippet

`m` for `markdown` mode. `eof` exits input and start description.

![image](https://user-images.githubusercontent.com/28176389/208107945-a7d8c09d-e306-445a-91a2-6d21b66f9254.png)

### Search for Snippet

![image](https://user-images.githubusercontent.com/28176389/208108100-0c75e346-725f-453d-b561-8c0be9a0839c.png)

![image](https://user-images.githubusercontent.com/28176389/208108146-75d42947-954a-4d16-a6e4-53da9d08a960.png)

## Markdown File

### Save Snippet

![image](https://user-images.githubusercontent.com/28176389/208108800-ccf45361-cfef-446d-9260-b08e0e036840.png)

![image](https://user-images.githubusercontent.com/28176389/208108916-22caa6d5-87b5-471f-a763-a9a6de720a30.png)

### Search for Snippet

![image](https://user-images.githubusercontent.com/28176389/208109016-1534737f-de61-4885-9419-ecbb7a1338d4.png)

![image](https://user-images.githubusercontent.com/28176389/208109049-fc8343c3-ebaa-4ac4-b4b0-cc3cd4119b68.png)

## Saving Image

![image](https://user-images.githubusercontent.com/28176389/208110386-d7d69d66-641e-4c68-a60e-ea2c4e11ffaa.png)

![image](https://user-images.githubusercontent.com/28176389/208111518-32c140e0-6354-4336-87e4-b03ba7556e58.png)


# Support for Images
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

# Supoort for Files

Ability to read markdown content from file is added on Jan, 12th, 2022. Now, for simple markdown snippets, you can fire up pet and input line by line. But for longer snippets with bunch of codes, just put the whole markdown in a file, and issue command

```
pet new /path/to/file
```

to add the content to pet.

Enjoy!
