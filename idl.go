package idl

// Information description language
// 信息描述语言

const (
	// 信息的结束符号
	INFO_SUFFIX = byte(':')
	// 属性的起始符号
	ATTR_PREFIX = byte('-')
)

// 信息
type Information struct {
	// 名称
	Name string
	// 属性集
	Attrs AttributeGroup
}

// 属性
type Attribute struct {

	// 名称
	name string
	// 值
	value string

}

// 获取属性名
// 可能为空
func (a *Attribute) GetName() string{
	return a.name
}

// 获取属性的值
func (a *Attribute) GetValue() string{
	return a.value
}

// 两个属性是否相等
func (a *Attribute) Equals(attr *Attribute) bool{
	if a.name == attr.name && a.value == attr.value{
		return true
	}
	return false
}

// 属性组
type AttributeGroup []*Attribute

// 添加一个属性
func (g *AttributeGroup) Add(attr *Attribute){
	t := *g
	t = append(t, attr)
	*g = t
}

// 移除
func (g *AttributeGroup) Remove(attr *Attribute){
	t := *g

	if (len(t) == 0){
		return
	}


	item := -1
	for i, v := range(t){
		if v.Equals(attr) {
			item = i
			break
		}
	}

	if item != -1 {
		t = append(t[:item], t[item + 1 :]...)
		*g = t
	}
}
