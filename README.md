# nagome-mac-say
Reading aloud [Nagome](https://github.com/diginatu/nagome) plugin for Mac.

## Usage
```console
# Install nagome-mac-say
$ go get -u github.com/y0za/nagome-mac-say

# Create plugin directory
$ mkdir -p ~/.config/Nagome/plugin/nagome-mac-say

# Create symbolic link
$ ln -s $GOPATH/bin/nagome-mac-say ~/.config/Nagome/plugin/nagome-mac-say/

# Copy config files (and change those as needed)
$ cp $GOPATH/src/github.com/y0za/nagome-mac-say/*.yml ~/.config/Nagome/plugin/nagome-mac-say/
```

## License
MIT License
