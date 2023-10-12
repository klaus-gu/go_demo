package main

// 泛型参数
type Slice[T int | float32 | float64] []T

// 泛型结构体
type MyStruct[T int | float32] struct {
	Name string
	Age  T
}

// 泛型接口
type MyInterface[T int | float32] interface {
	Print(T)
}

type MyInterfaceImpl struct {
}

func (m MyInterfaceImpl) Print(t int) {
	//TODO implement me
	panic("implement me")
}

type CommonInterface interface {
	~string | ~[]rune

	Read() string

	Write(content string)
}

type BasicInterface interface {
	Read() (words string)

	Write(content string)
}

type MyBasicInterface struct {
}

func (m MyBasicInterface) Read() string {
	//TODO implement me
	panic("implement me")
}

func (m MyBasicInterface) Write(content string) {
	//TODO implement me
	panic("implement me")
}

type MyCommonInterface struct {
}

func (m MyCommonInterface) Read() string {
	//TODO implement me
	panic("implement me")
}

func (m MyCommonInterface) Write(content string) {
	//TODO implement me
	panic("implement me")
}
