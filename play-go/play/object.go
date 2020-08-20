/*
包 <package name> 文档注释。

参考 https://github.com/LCSyy/go-play

名称：
	包引入后，包名将作为包内容的访问符。因此选择一个短小，易书写的包名十分重要。
	一般包名都为小写的一个单词。包名一般和其包所在目录同名。

	接口名称。如果是只有一个方法的接口，其名称则为方法名加-er，或者适当修改后的
	名称形式。如：Reader, Writer, CloseNotifier等。

分号：
	Go语言采用分号来表示语句的终止。但分号并不在源码中使用，而是由语法分析器通过
	一些规则添加到语句末尾。添加分号的规则是：
		如果一行末尾的符号能够有效结束该行语句，则添加分号。
*/
package play


// Object 结构是play框架的核心类型。
type Object struct {
	ObjectName string
	status     uint // 对象当前状态
}

// NewObject Object构造函数。
func NewObject(name string) Object {
	return Object{
		ObjectName: name,
		status:     Active,
	}
}

// SetObjectName 设置对象的名称。
func (self *Object) setObjectName(name string){
	self.ObjectName = name
}

// Status 用于获取对象当前状态。
func (self Object) Status() uint {
	return self.status
}
