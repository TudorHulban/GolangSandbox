## Requirements
pwd - command displays the current working directory and,<br/>
cd - changes the location of working directory.

- The default root directory is "/".
- If the cd parameter contains ".."(without quotes), that means to step one directory back.
- The absolute path of directory is separated by slashes "/"(without quotes).

Input
```
pwd
cd /home/csed
pwd
cd /lab/root/../dir
pwd
cd /home
pwd
cd lab
pwd
```

Output
```
/
/home/csed/
/lab/dir/
/home/
/home/lab/
```
