
One Image Server embeds a single image at compile time and serves that image at runtime, no matter what you ask for.

The image can be anything you want, assuming it's legal for you to use.

Examples

* Test image during development/qa
* A maintenance image when your servers are down
* Your significant other that you need to see periodically and never can find their picture
* An emergency no-smoking image for when you need one
* A reminder of your favorite beer
* A picture of cats
* A picture of dogs herding cats

Build
=====
Start with this dependency.
```
go get -u github.com/jteeuwen/go-bindata/...
```

Then copy your image file to 'resources/onlyimage.jpg'
The name 'resources/onlyimage.jpg' is required but it can be whatever you want as long as it's a bunch of bytes a browser would recognize as image.

```
go-bindata resources/
go build
```

Alternative Cross Compilation for Linux
```
GOOS=linux GOARCH=amd64 go build -o one-image-server-linux64
```

Run
===
Copy just the binary to wherever and run it.
```
   scp ./one-image-server <wherever>
   ssh <wherever>
   ./one-image-server
```

Once it's up, point your browser to "http://<server>:9666/"
and enjoy, over and over.

Rationale
=========
The original use case for One Image Server was to serve a test image regardless of the request path
while I was writing some image handling code.
I didn't care about the image as much as getting something valid from the requested path and a 200 response code.


Development Notes
=================
jteeuwen go-bindata disappeared at one point.
Alternative: go get -u github.com/a-urth/go-bindata/...
