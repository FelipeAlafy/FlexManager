# FlexManager



## Building - Using docker:
* First install docker.
```
sudo docker run -it fedora /bin/bash
```

```
yum install mingw64-gtk3 go glib2-devel
```

```
dnf install llvm clang make gtk3-devel mingw64-gcc-12.2.1-8.fc38.x86_64
```

```
bash -c "sed -i -e 's/-Wl,-luuid/-luuid/g' /usr/x86_64-w64-mingw32/sys-root/mingw/lib/pkgconfig/gdk-3.0.pc"
```
* Setup git in container to clone, this repo.

```
cd root
```

```
 mkdir go go/src go/bin go/pkg
```

```
cd go/src
```
* Get program from github, and CD to it.
* The next step will take about 8 minutes. 
```
PKG_CONFIG_PATH=/usr/x86_64-w64-mingw32/sys-root/mingw/lib/pkgconfig CGO_ENABLED=1 CC=x86_64-w64-mingw32-gcc GOOS=windows GOARCH=amd64 go install -v github.com/gotk3/gotk3/gtk
```

* Compiling to windows:
```
CGO_ENABLED=1 CC=x86_64-w64-mingw32-gcc GOOS=windows GOARCH=amd64 go build -ldflags -H=windowsgui
```
* Get gtk libs:
```
yes | cp -r /usr/x86_64-w64-mingw32/sys-root/mingw/*
```
* get id from current session image:
```
sudo docker ps -alq
```
* Copy your program from container to your local machine:
```
sudo docker cp <image-id>:/root/go/src/Union/union Documentos/union
```
* Remember to set the correct permissions of the folder to allow your user to open it and copy it. 

# Install:
* Install the [
GTK-for-Windows-Runtime-Environment-Installer ](https://github.com/tschoonj/GTK-for-Windows-Runtime-Environment-Installer) or just donwload from this direct [link](https://drive.google.com/file/d/1Gyi5yugTFvHv6NLX9WBJHsQ1f9PpMZZX/view?usp=sharing).
* And now, you're able to install Union in your windows machine.
* Just move Union folder to Documents or some other folder in your machine.
* If you want an icon in an executable file, just create a shortcut in windows and insert the project logo. And with that you can pin Union to the Start menu or the taskbar.
