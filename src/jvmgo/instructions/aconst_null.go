package instructions

type aconst_null struct {

}

func (self *aconst_null) fetchOperands(bcr *BytecodeReader) {
    // no operands
}

func (self *aconst_null) execute() {
    // todo
}