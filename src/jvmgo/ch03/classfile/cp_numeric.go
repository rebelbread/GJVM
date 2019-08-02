package classfile

import "math"

type ConstantIntegerInfo struct {
	val int32
}
func (self *ConstantIntegerInfo) readInfo(read *ClassReader){
	val := read.readUint32()
	self.val = int32(val)
}
type ConstantFloatInfo struct {
	val float32
}
func (self *ConstantFloatInfo) readInfo(read *ClassReader){
	val := read.readUint32()
	self.val = math.Float32frombits(val)
}
type ConstantLongInfo struct {
	val int64
}
func (self *ConstantLongInfo) readInfo(read *ClassReader){
	val := read.readUint64()
	self.val = int64(val)
}
type ConstantDoubleInfo struct {
	val float64
}
func (self *ConstantDoubleInfo) readInfo(read *ClassReader){
	val := read.readUint64()
	self.val = math.Float64frombits(val)
}