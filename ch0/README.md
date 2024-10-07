# ch0 ：第0份编译
使用 github.com/llir/llvm/ 写死生成中间代码再通过clang编译成程序
```c
int main() {
    return 2;
}
```
使用llvm-go生成中间代码[Basic Introduction](https://llir.github.io/document/user-guide/basic/)。