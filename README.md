# 中文 README （For English service, please scroll down）

命令行代码片段管理工具。感谢 `knqyf263` 制作的原 Pet 工具. [原工具仓库](https://github.com/knqyf263/pet)

增加了 Markdown 功能，图片功能。

整个程序只有一个 snippet.toml 文件，一个 config.toml 文件（默认生成在 /home/$USER/.config/pet/）需要保存（直接开一个 github 私有仓库同步即可）。

保存命令或者 Markdown 文字，增加关键字；需要使用的时候，通过关键字搜索（fzf）即可。

目的在于提高效率，无需离开命令行，配合 tmux，无缝获取需要的命令或者信息。

# 依赖

Pet 需要依赖 `glow` 来在命令行显示 Markdown。需要 `feh` 来显示图片，并需要 `fzf` 做模糊搜索。

```
sudo apt install fzf feh
go install github.com/charmbracelet/glow@latest
```

详情见安装章节。

# 安装

从源码编译，或者下载编译好的文件。

## 编译

### Linux / Windows WSL

1. 安装依赖
   ```txt
   sudo apt install fzf feh proxychains4
   go install github.com/charmbracelet/glow@latest
   ```

2. 克隆仓库到 GOPATH；
   ```txt
   git clone https://github.com/0prrr/pet.git
   ```

3. 编译
   ```txt
   cd pet
   
   go build
   ```
   
4. 将 `pet` 二进制移到 `/usr/local/bin/`
   ```txt
   mv pet /usr/local/bin
   ```
   
5. 查看下文使用方式（增加 Alias），粘贴适用的代码到 shell `.rc` 文件;

### macOS

1. 安装依赖
   ```txt
   brew install fzf feh xorg-server xinit
   
   go install github.com/charmbracelet/glow@latest
   
   // Optional
   brew install proxychains-ng
   ```
   
2. 克隆仓库到 GOPATH；
   ```txt
   git clone https://github.com/0prrr/pet.git
   ```
   
3. 编译
   ```txt
   cd pet
   
   go build
   ```
   
4. 将 `pet` 二进制移到 `/usr/local/bin/`
   ```txt
   mv pet /usr/local/bin
   ```

5. 查看下文使用方式（增加 Alias），粘贴适用的代码到 shell `.rc` 文件;


## 下载预编译的二进制文件

下载相应平台最新版本（v0.9），解压到 `/usr/local/bin/` 即可。

# 使用方式

## 增加 Alias

Linux, 粘贴下面的代码到 shell 的 `.rc` 文件。

```txt
pets () {
    # proxychains 是因为我使用 Imgur 保存我的图片
    # 更多关于图片的支持看下文的图片支持章节
    cmd=`proxychains -q /usr/local/bin/pet search`
    if [[ "$cmd" == "#"* ]]
    then
        echo "$cmd" | glow -s dark - | less -r
        echo "$cmd" | glow -s dark -
    else
        echo "$cmd"
        echo "$cmd" | tr -d '\n' | xclip -sel c
    fi
}
```

Windows WSL, 粘贴下面的代码到 shell 的 `.rc` 文件。

```txt
pets () {
    # proxychains 是因为我使用 Imgur 保存我的图片
    # 更多关于图片的支持看下文的图片支持章节
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

macOS, 粘贴下面的代码到 shell 的 `.rc` 文件。

```txt
export DISPLAY=:0.0

if ! pgrep -x "xquartz" >/dev/null; then
    xquartz &
fi

pets () {
    # proxychains can be removed according to the image service you are using
    # refer to image support in the below section
    cmd=`proxychains -q /usr/local/bin/pet search`
    if [[ "$cmd" == "#"* ]]
    then
        echo "$cmd" | glow -s dark - | less -r
        echo "$cmd" | glow -s dark -
    else
        echo "$cmd"
        echo "$cmd" | tr -d '\n' | pbcopy
    fi
}
```

运行 `source xxx.rc`, 然后使用 `pets` 运行。

## 单行命令

### 保存单行命令

`n` 代表 `normal` 模式.

![image](https://user-images.githubusercontent.com/28176389/208106824-b7ff653e-9642-411a-a2d0-aace7b278d14.png)

### 搜索命令

运行 `pets`, 输入关键字，回车。

![image](https://user-images.githubusercontent.com/28176389/208107108-fa81d9e9-2cdd-49d9-90cd-7295adf86af4.png)

 注意 `<>` 是参数, pet 会要求你填入，填入参数后回车。

![image](https://user-images.githubusercontent.com/28176389/208107357-ef61740c-b8ce-4268-8ee6-1679e882175c.png)

命令显示在命令行，并已经复制到剪切板，粘贴运行即可。

![image](https://user-images.githubusercontent.com/28176389/208108345-cd2e5a1c-f1e0-4e2e-b704-378350686e55.png)

## Markdown 单行模式

### 保存代码片段

`m` 代表 `markdown` 模式. `eof` 代表输入结束。

![image](https://user-images.githubusercontent.com/28176389/208107945-a7d8c09d-e306-445a-91a2-6d21b66f9254.png)

### 搜索代码片段

运行 `pets`, 输入关键字，回车。

![image](https://user-images.githubusercontent.com/28176389/208108100-0c75e346-725f-453d-b561-8c0be9a0839c.png)

`glow` 会将代码片段显示在命令行。

![image](https://user-images.githubusercontent.com/28176389/208108146-75d42947-954a-4d16-a6e4-53da9d08a960.png)

## Markdown 文件

### 保存代码片段

在文件中输入 Markdown 内容，保存。

![image](https://user-images.githubusercontent.com/28176389/208108800-ccf45361-cfef-446d-9260-b08e0e036840.png)

使用 `pet new /path/to/file` 保存文件中的内容。优点是随时修改。

![image](https://user-images.githubusercontent.com/28176389/208108916-22caa6d5-87b5-471f-a763-a9a6de720a30.png)

### 搜索代码片段

运行 `pets`, 输入关键字，回车。

![image](https://user-images.githubusercontent.com/28176389/208109016-1534737f-de61-4885-9419-ecbb7a1338d4.png)

`glow` 会将代码片段显示在命令行

![image](https://user-images.githubusercontent.com/28176389/208109049-fc8343c3-ebaa-4ac4-b4b0-cc3cd4119b68.png)

## 带图片的 Markdown 文件

截图软件都带上传插件。上传截图到图片服务（如 Imgur），然后以如下格式保存图片链接：

`img::https://example.com/5Zhdgesz.png`

如下图

![image](https://user-images.githubusercontent.com/28176389/208110386-d7d69d66-641e-4c68-a60e-ea2c4e11ffaa.png)

然后保存该 Markdown 文件。

![image](https://user-images.githubusercontent.com/28176389/208111518-32c140e0-6354-4336-87e4-b03ba7556e58.png)

可以添加任意数量的图片，一张一行即可。

```txt
img::https://exmaple.com/img1.png
img::https://exmaple.com/img2.png
```

# 已知问题

1. 使用简单的 Markdown（如下例子），否则 `glow` 可能会显示不正常。以下 Markdown 语法已经过测试。

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

   Markdown 示例：

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

   Markdown 代码片段都以 `#` 开头（第一行第一个字符）。

2. 非 Markdown 模式下（单行命令模式），尖括 (`<` and `>`) 会被解析成参数，pet 会弹出输入框。这对于 PHP 单行命令并不友好。目前的解决方式是，保存 PHP 单行命令成 Markdown，或者在保存单行命令的时候，省略右括号（`>`），在使用的时候添上。

3. 在参数输入框页面，避免粘贴过长的参数。整个库有些问题，过长的文字格式会出错。建议在粘贴命令之后自行添加。

4. 由于支持 Markdown，所以在 Markdown 模式下，pet 不能自作主张转义所有的反斜杠，否则会导致该有的缩进无法正常显示，代码格式混乱。所以，反斜杠需要用户自行转义。比如路径：`C:\Windows\System32`, 转译成：`C:\\Windows\\System32`; 代码也一样: `printf("Value is: %d\n", val);`, 转译成: `printf("Value is: %d\\n", val);`. VIM 中, 可以使用整个命令做全局替换 `%s/\\/\\\\/g`. 其他编辑器，只需将 `\` 替换成 `\\` 即可。（！！！注意，`normal` 模式下，不需要转义！！！）

5. 由于 feh 依赖 X11，所以在 mac 必须使用 xquartz 来解决 feh 的运行问题。目前还没有计划使用其他的方式代替 feh。Again, bear with me。

# Bug 上报

欢迎 issue。

# 协议

MIT

# 致谢

感谢 Teppei Fukuda (knqyf263) 制作的 Pet Snippet manager. [原仓库](https://github.com/knqyf263/pet)

---------------------------------------------------------------------------------------------------------------------------------------------------
---------------------------------------------------------------------------------------------------------------------------------------------------

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
sudo apt install fzf feh
go install github.com/charmbracelet/glow@latest
```

# Install

You can either build pet from source, or download the pre-compiled binary from release.

## Build from Source

### Linux / Windows WSL

1. Install dependencies
   ```txt
   sudo apt install fzf feh
   
   go install github.com/charmbracelet/glow@latest
   ```

2. Clone the repo to `GOPATH` on your system
   ```txt
   git clone https://github.com/0prrr/pet.git
   ```

3. Compile
   ```txt
   go build
   ```

4. Move `pet` to `/usr/local/bin/pet`
   ```txt
   mv pet /usr/local/bin
   ```

5. Refer to Shell Function section next and add appropriate code to shell's `.rc` file;

### macOS

1. Install dependencies
   ```txt
   brew install fzf feh xorg-server xinit
   
   go install github.com/charmbracelet/glow@latest
   
   // Optional
   brew install proxychains-ng
   ```

2. Clone the repo to `GOPATH` on your system
   ```txt
   git clone https://github.com/0prrr/pet.git
   ```
   
3. Compile
   ```txt
   cd pet
   
   go build
   ```
   
4. Move `pet` to `/usr/local/bin/`
   ```txt
   mv pet /usr/local/bin
   ```
   
5. Refer to Shell Function section next and add appropriate code to shell's `.rc` file;

## Download Pre-compiled Binary

Download the pre-compiled binary and extract to `/usr/local/bin/`;

# Usage

## Shell Function

For Linux, put the following scirpt in whatever shell's `.rc` file.

```txt
pets () {
    # proxychains 是因为我使用 Imgur 保存我的图片
    # 更多关于图片的支持看下文的图片支持章节
    cmd=`proxychains -q /usr/local/bin/pet search`
    if [[ "$cmd" == "#"* ]]
    then
        echo "$cmd" | glow -s dark - | less -r
        echo "$cmd" | glow -s dark -
    else
        echo "$cmd"
        echo "$cmd" | tr -d '\n' | xclip -sel c
    fi
}
```

For Windows WSL, put the following scirpt in whatever shell's `.rc` file.

```txt
pets () {
    # proxychains can be removed according to the image service you are using
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

For macOS, put the following scirpt in whatever shell's `.rc` file.

```txt
export DISPLAY=:0.0

if ! pgrep -x "xquartz" >/dev/null; then
    xquartz &
fi

pets () {
    # proxychains can be removed according to the image service you are using
    # refer to image support in the below section
    cmd=`proxychains -q /usr/local/bin/pet search`
    if [[ "$cmd" == "#"* ]]
    then
        echo "$cmd" | glow -s dark - | less -r
        echo "$cmd" | glow -s dark -
    else
        echo "$cmd"
        echo "$cmd" | tr -d '\n' | pbcopy
    fi
}
```

Run `source xxx.rc`, then fire up `pets` to run the executable.

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

4. Because of Markdown support, pet cannot decide to escape all backslashes there in the command/snippets, because that will mess up with code formats. PLEASE escape backslash (double them) manually. Like if you're saving path: `C:\Windows\System32`, escape them as: `C:\\Windows\\System32`; for code, same: `printf("Value is: %d\n", val);`, escape as: `printf("Value is: %d\\n", val);`. If you're using vim, `%s/\\/\\\\/g` will do. Other text editors, just replace `\` with `\\`.（!!! Note that no escaping should be done under `normal` mode !!!）

5. feh (I'm old) replies on X11 to function. So on macOS, it must be solved with xquartz. feh won't be replaced in anticipated future. Again, bear with me.

# Bug Reports

Feel free to open issue.

# License

MIT

# Credit

Again, thanks to Teppei Fukuda (knqyf263) who made Pet Snippet manager. [Original Pet repo](https://github.com/knqyf263/pet)
