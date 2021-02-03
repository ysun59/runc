### prerequisites
#### install golang
#### set environment variable $GOPATH to the location you want
#### create log file log.txt
`sudo touch /tmp/myrunc/log.txt`
`sudo chmod 777 /tmp/myrunc/log.txt`

### clone, compille, and install
#### go to $GOPATH/src/github.com/opencontainers/, if not exist, create it
#### under $GOPATH/src/github.com/opencontainers, clone this repo
#### go $GOPATH/src/github.com/opencontainers/runc, run the following commands
`make clean`
`make BUILDTAGS='seccomp apparmor'`
`sudo make install`
#### now, the customized runc executable is installed at /usr/local/sbin/runc

### modify docker config
#### modify docker config file /etc/docker/daemon.json, if not exists, create it
#### add the following configuration
`
{
  "default-runtime": "custom",
  "runtimes": {
    "custom": {
      "path": "/usr/local/sbin/runc"
    }
  } 
}
`
#### save this file and restart the docker daemon using the following command
`sudo systemctl restart docerk.server`

#### Good luck~