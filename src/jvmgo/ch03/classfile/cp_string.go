package classfile

type ConstantStringInfo struct {
	cp          ConstantPool
	stringIndex uint16
}

func (self *ConstantStringInfo) readInfo(reader *ClassReader){
	self.stringIndex = reader.readUint16()
}
func (self *ConstantStringInfo)String(){
	self.cp.getUtf8(self.stringIndex)
}

