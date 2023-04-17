# FlexManager



## Building - Using docker:
* First install docker.
* ```sudo docker run -it fedora /bin/bash```
* `yum install mingw64-gtk3 go glib2-devel`
* `dnf install llvm clang make gtk3-devel`
* `bash -c "sed -i -e 's/-Wl,-luuid/-luuid/g' /usr/x86_64-w64-mingw32/sys-root/mingw/lib/pkgconfig/gdk-3.0.pc"`
* Setup git in container to clone, this repo.
* `cd root`
```
 mkdir go go/src go/bin go/pkg
```
* `cd go/src`
* Get program from github, and CD to it.
* `PKG_CONFIG_PATH=/usr/x86_64-w64-mingw32/sys-root/mingw/lib/pkgconfig CGO_ENABLED=1 CC=x86_64-w64-mingw32-gcc GOOS=windows GOARCH=amd64 go install -v github.com/gotk3/gotk3/gtk` #This will take about 8 minutes. 
* `CGO_ENABLED=1 CC=x86_64-w64-mingw32-gcc GOOS=windows GOARCH=amd64 go build -ldflags -H=windowsgui` #Compile
* `yes | cp -r /usr/x86_64-w64-mingw32/sys-root/mingw/*` . #Get gtk libs
* `sudo docker ps -alq` #get id from current session image
* `sudo docker cp <image-id>:/root/go/src/Union/union Documentos/union`

# Install:
* Install the [
GTK-for-Windows-Runtime-Environment-Installer ](https://github.com/tschoonj/GTK-for-Windows-Runtime-Environment-Installer) or just donwload from this direct [link](https://drive.google.com/file/d/1Gyi5yugTFvHv6NLX9WBJHsQ1f9PpMZZX/view?usp=sharing).
* And now, you're able to install Union in your windows machine.
* Just move Union folder to Documents or some other folder in your machine.
* If you want an icon in an executable file, just create a shortcut in windows and insert the project logo. And with that you can pin Union to the Start menu or the taskbar.
