package tempconv0

type Celsius float64    // 摄氏温度
type Fahrenheit float64 // 华氏温度

const (
	AbsoluteZeroC Celsius = -273.15 // 绝对零度
	FreezingC     Celsius = 0       // 结冰点温度
	BoilingC      Celsius = 100     // 沸水温度
)

// CToF 我们在这个包声明了两种类型：Celsius和Fahrenheit分别对应不同的温度单位。
// 它们虽然有着相同的底层类型float64，但是它们是不同的数据类型，
// 因此它们不可以被相互比较或混在一个表达式运算。刻意区分类型，
// 可以避免一些像无意中使用不同单位的温度混合计算导致的错误；
// 因此需要一个类似Celsius(t)或Fahrenheit(t)形式的显式转型操作才能将float64转为对应的类型。
// Celsius(t)和Fahrenheit(t)是类型转换操作，它们并不是函数调用。类型转换不会改变值本身，但是会使它们的语义发生变化。
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }
