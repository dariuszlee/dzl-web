This is the first post on creating an HTTP/S session automator to do some tasks on Kijiji (Canadian Craig's list).

### Motivation

Kijiji has a service to keep your post at the top of a queue when searching in categories. I don't want to pay so lets do this for free.

<br>

The newest post shows up at the top of the queue, so simply create a mechanism to check the ad status and reposting it if it is too old should do the trick. This can be done simply by recreating the HTTP POST request used to generate the post in the first place...

### Problem 1:

Posting an ad on kijiji isn't simply an HTTP post, it does some javascript input validation beforehand that I would like to capture. 

### Solution 1:

Using a headless browser to execute the javascript and see what comes back.
