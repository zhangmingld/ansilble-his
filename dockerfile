ARG
1.ARG will always be override by ENV
2.Build will failed if you use a non-exists arg in dockerfile
3.ARG are effcitve only after the its defination line
4.ARG has some pre-defined value , some proxy arg, please visit the webpage
VOLUME
1.Volume creates a mountpoint in container
2.User can not decide where the mountpoint will be mounted on the host , docker decided it
