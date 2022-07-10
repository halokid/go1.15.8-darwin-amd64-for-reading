package gotest

// 检测单行输出
func ExampleSayHello() {
  SayHello()
  // OutPut: Hello World
}
// 检测多行输出
func ExampleSayGoodbye() {
  SayGoodbye()
  // OutPut:
  // Hello,
  // goodbye
}
// 检测乱序输出
func ExamplePrintNames() {
  PrintNames()
  // Unordered output:
  // Jim
  // Bob
  // Tom
  // Suex
}
