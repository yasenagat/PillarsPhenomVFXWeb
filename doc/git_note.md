文件在仓库中存在，但是不想每次修改后就提交（这样的文件一般是与本机开发环境相关的文件），可以这样操作：
1. 打开：
.git/config

2. 增加：
[alias]
	ignore = update-index --assume-unchanged
	unignore = update-index --no-assume-unchanged
	ignored = !git ls-files -v | grep "^[[:lower:]]"
	
3. 运行：
git ignore filenam
这样，再次提交时这个文件的修改就不会被提交。
4. 如果需要再次提交这个文件，运行 :
git unignore filename