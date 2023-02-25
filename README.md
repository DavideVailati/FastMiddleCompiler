# FastMiddleCompiler
A fast middle layer for all-purpose JSON minecraft bedrock development
It's use is exclusive to Escher projects, a branch to generalize it will be created
# Requirements
You must have Go 1.20 installed and able to compile source code, instructions and commands will be provided below.
# How to set up
First clone the repository in any folder 
```shell
git clone https://github.com/DavideVailati/FastMiddleCompiler.git
```
Then go in the directory
```shell
cd FastMiddleCompiler
```
Execute the setup bat file
```shell
setup.bat
```
Now grab an editor of your choice and edit the `rootPath.txt` file so that its contents match the directory of your 
Escher project. For example:
```shell
C:/echerProjects/exampleProject/
```
Then open the `input.fmc` file, here is where all of your code will be called by the compiler. Paste in this example
code after the first line, which will be the directory in which the file will be saved.

```
{
    version(),
    "minecraft:entity":{
        "description":{
            bpname(example)
        }
    }
}
```
So your whole `input.fmc` file should look something like this:
```
$bp/entities/test.json
{
    version(),
    "minecraft:entity":{
        "description":{
            bpname(example)
        }
    }
}
```
To compile FMC:
```shell
go build fmc
```
Then, double-click on the `fmc.exe` executable or run it through commands.
# How to use the language
The main file you'll want to edit is the `input.fmc` file, from now on referred as input.
Inside there you'll write JSON/FMC files, broken up by lines that start with $, as those are the lines
specifying the path to put the compiled JSON in. For example:
```
$example/path1
{
    //JSON1
}
$example/path2
{
    //JSON2
}
```
JSON1 will be compiled into example/path1, and JSON2 in example/path2.
In runtime, `bp` and `rp` will be replaced with the correct paths, depending on the contents of `rootPath.txt`.
For example `$bp/entities/example.json` will become `C:/escherProjects/exampleProject/0. Behaviour Pack-BP/entities/example.json`.