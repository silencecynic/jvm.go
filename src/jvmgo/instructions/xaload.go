package instructions

import "jvmgo/rtda"

//type Ref *rtda.Obj

// Load reference from array
type aaload struct {NoOperandsInstruction}
func (self *aaload) execute(thread *rtda.Thread) {
    stack, arrRef, index := popArrAndIndex(thread)
    refArr := arrRef.Fields().([]*rtda.Obj)
    ref := refArr[checkArrIndex(index, len(refArr))]
    stack.PushRef(ref)
}

// Load byte or boolean from array 
type baload struct {NoOperandsInstruction}
func (self *baload) execute(thread *rtda.Thread) {
    stack, arrRef, index := popArrAndIndex(thread)
    byteArr := arrRef.Fields().([]int8)
    val := byteArr[checkArrIndex(index, len(byteArr))]
    stack.PushInt(int32(val))
}

// Load char from array 
type caload struct {NoOperandsInstruction}
func (self *caload) execute(thread *rtda.Thread) {
    stack, arrRef, index := popArrAndIndex(thread)
    charArr := arrRef.Fields().([]uint16)
    val := charArr[checkArrIndex(index, len(charArr))]
    stack.PushInt(int32(val))
}

// Load double from array 
type daload struct {NoOperandsInstruction}
func (self *daload) execute(thread *rtda.Thread) {
    stack, arrRef, index := popArrAndIndex(thread)
    doubleArr := arrRef.Fields().([]float64)
    val := doubleArr[checkArrIndex(index, len(doubleArr))]
    stack.PushDouble(val)
}

// Load float from array 
type faload struct {NoOperandsInstruction}
func (self *faload) execute(thread *rtda.Thread) {
    stack, arrRef, index := popArrAndIndex(thread)
    floatArr := arrRef.Fields().([]float32)
    val := floatArr[checkArrIndex(index, len(floatArr))]
    stack.PushFloat(val)
}

// Load int from array 
type iaload struct {NoOperandsInstruction}
func (self *iaload) execute(thread *rtda.Thread) {
    stack, arrRef, index := popArrAndIndex(thread)
    intArr := arrRef.Fields().([]int32)
    val := intArr[checkArrIndex(index, len(intArr))]
    stack.PushInt(val)
}

// Load long from array 
type laload struct {NoOperandsInstruction}
func (self *laload) execute(thread *rtda.Thread) {
    stack, arrRef, index := popArrAndIndex(thread)
    longArr := arrRef.Fields().([]int64)
    val := longArr[checkArrIndex(index, len(longArr))]
    stack.PushLong(val)
}

// Load short from array 
type saload struct {NoOperandsInstruction}
func (self *saload) execute(thread *rtda.Thread) {
    stack, arrRef, index := popArrAndIndex(thread)
    shortArr := arrRef.Fields().([]int16)
    val := shortArr[checkArrIndex(index, len(shortArr))]
    stack.PushInt(int32(val))
}

func popArrAndIndex(thread *rtda.Thread) (*rtda.OperandStack, *rtda.Obj, int) {
    stack := thread.CurrentFrame().OperandStack()
    arrRef := stack.PopRef()
    index := int(stack.PopInt())
    if arrRef == nil {
        // todo
        panic("NullPointerException")
    }
    return stack, arrRef, index
}
