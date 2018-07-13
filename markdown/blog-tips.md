This will be a quick post to show how to recreate the animated GIFs you see on a github README. All from the command line (plus a few utilities)!

<br>
## Tools

I will describe the tools I use which will show by example a simple workflow GIF Creation workflow.

<br>
The general process is:
1. Cast or capture screen
2. Test capture
3. Upload to media server (spoiler: I use imgur here)
4. Embed url in markdown (or where ever else you want)

### 1. Screen Casting Tool: byzanz-record
<br>
This tool is available on arch in the community repository. Super super simple to use.
Simply type in your console:

<br>
`byzanz-record test`

`ffmpeg -i test test.gif`

<br>
Voila you have a quick gif of your screen.

CAVEAT: Multiple monitors will be displayed if you have this setup. (I have i3 shortcuts to switch my monitor setup on the fly. I simply hit the single screen mode shortcut, record, then go back)

### 1a. Screen Capture Tool: flameshot (?)
<br>
Another community arch/pacman package. Its pretty clever and works how you would expect screen capture utilies to work. Here is a quick gif using byzanz to capture flameshot gui.

<br>
<img alt="Flameshot Usage" src="https://i.imgur.com/WdiJGMU.gif" style="width:100%" /img>
<br>
<br>

### 2. Video Viewer: mpv

Okay, this is a quick one. *mpv* is a cli video viewer that can show you the output of your gifs. Simply run like so:

`mpv test.gif`

Make sure the file looks good and now lets upload it.

<br>
### 3. Uploading To Media Server
<br>
My *go-to* image data storage utility is Imgur. There are lots of CLI utilities that can give you instantaneous uploads. Anyways, here is my setup...

<br>
<img alt="Imgur Upload" src="https://i.imgur.com/WdiJGMU.gif" style="width:100%" /img>
<br>


<br>
There are a lots of utilites to use from the command line to upload images/gifs to imgur. The utility I use here aptly named **imgur**. It is available on the Arch User Repositories.

### 4. Embedding Url in Markdown

<br>
Using imgur as a media server is simple. Once you have uploaded the image, grab the url of the image returned by the upload utility and post it in your markdown like so:

`!["Flameshot Usage"](/path/to/imgur)`
<br>
or
<br>
`<img alt="Flameshot Usage" src="https://i.imgur.com/WdiJGMU.gif" style="width:100%" /img>`

<br>
So there you have it, a quick and dirty *all from the commandline* way of performing Giffifying anything you please. I'm sure you have another way of doing this, let me know how you do it.

Next blogging tip post I think I will do CSS with Markdown, these posts look pretty bland (ATM!).

Go do something!

Darius
