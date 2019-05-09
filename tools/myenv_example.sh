#cp myenv_example.sh ~/


export PATH=~/work/github/mingw/mingw/bin/:$PATH
#需要windows格式路径的gopath
export GOPATH=`cygpath.exe -p ~ -a -w`\\go

export TEMP=/tmp
export TMP=/tmp

