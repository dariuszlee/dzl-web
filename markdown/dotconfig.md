A way to manage "dot" files between machines is an easy way to make you comfortable on any environment you find yourself on. Also... It makes you feel a little better about all the time you end up wasting on the things.

<br>
### Why would I need this?
<br>

If you don't know why you would need this, you probably shouldn't be reading this. A dotfile config is a hidden file, usually somewhere in the home directory of a linux/mac/cygwin machine, that is used by specific applications to... store config variables obviously.

Some of these config files can be very elaborative depending on the application that uses them, so re-using these files between machines has obvious benifits.

<br>
### Basic Solution

1. All config files are included under a single directory node (preferably your home directory)
2. A .gitignore file at the base of this directory/repo

The basic solution for my setup requires git installed. If you don't have access to git, you can stop here. Also, go ahead and make yourself an account on Github, Gitlab, or whatever Git hosting solution you want.

Go ahead and run

`cd <directory in step 1>`

`git init`

Add the .gitignore by running

`touch .gitignore`

<br>
### The .gitignore

The bulk of the solution is in the Gitignore file. It starts with ignoring all files. This is achieved by the following line at the start of your .gitignore.

`*`

This is followed by *not* ignoring all the files you require. Some examples include:

<code>
!.zshrc 
!.vimrc 
!.tmux.conf 
</code>

The *!* before the file name signals to not ignore these files.

### Last Step

`git add . & git commit .` then `git push` to your remote repository and voila, you have a globally accessible config setup.

<br>
### Further Considerations
There are a few things I would like to add to make this even more configurable.

#### 1. Machine specific configurations

My current, quick and easy solution to this running the command `whoami` and checking the output in an `if` statement.

This doesn't work for dotfiles that have no way of accessing the command line or can't do conditional switching.

#### 2. Password Management

I use GnuPG and pass to encrypt/decrypt passwords . I haven't looked into a distributed way to handle passwords but encorporating this into a setup would make life easier. As for right now, I need to manually copy all my passwords between computers.

<br>
### La Conclussion

This is one of the easier ways to achieve a solution like this. If you have any improvements, of course, let me know below.

<br>
Sincerely,
Darius


