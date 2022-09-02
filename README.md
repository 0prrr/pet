Yet another Pet (or I just call it Pet-Enhanced). Since I haven't changed any of the original code, I will not change the name of the tool.
   
Much appreciation to `knqyf263` who made Pet. [Original Pet repo](https://github.com/knqyf263/pet)

I personally use it with `glow`, a CLI markdown rendering tool to better format the output. Like this:

![image](https://user-images.githubusercontent.com/28176389/129502096-65807f4f-10cb-4db6-9bae-dc770283ef63.png)

I haven't tested all the markdown syntax, just `header`, `inline code block` and `table`. If markdown went wrong, please checkout the [glow doc](https://github.com/charmbracelet/glow).

Another thing, if you want to put angle brackets (`<>`) in your markdown text, remember to html encode them as `&lt;` and `&gt;` or use backtick to surround the angle bracket"`" when in a table.

Image support added. I personally upload image to imgur, and add in the snippet the keyword `img::` followed by the image link, eg:

```
img::https://i.imgur.com/V3vsK.png
```

And if you have `feh` installed, when you search for that piece of snippet, then `feh` will automatically display the image, which makes my wiki tools more powerful.
![2022-01-04 19_48_02-pet_README md at main · reyalpmi_pet](https://user-images.githubusercontent.com/28176389/148054694-b294b7a5-1517-4784-b0d2-0b89457ee9f1.png)

Ability to read markdown content from file is added on Jan, 12th, 2022. Now, for simple markdown snippets, you can fire up pet and input line by line. But for longer snippets with bunch of codes, just put the whole markdown in a file, andissue command

```
pet new /path/to/file
```

to add the content to pet.

Enjoy!
