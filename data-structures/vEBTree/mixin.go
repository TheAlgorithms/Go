package vEBTree

type rsVEBTreeUInt32Mixin struct {
	rsVEBTreeMixin
}

func (m *rsVEBTreeUInt32Mixin) Less(lgu int, k1, k2 interface{}) bool {
	return (k1.(uint32) & ((1 << uint32(lgu)) - 1)) < (k2.(uint32) & ((1 << uint32(lgu)) - 1))
}

func (m *rsVEBTreeUInt32Mixin) High(lgu int, key interface{}) interface{} {
	return (key.(uint32) & ((1 << uint32(lgu)) - 1)) >> uint32(lgu>>1)
}

func (m *rsVEBTreeUInt32Mixin) Low(lgu int, key interface{}) interface{} {
	return key.(uint32) & ((1 << uint32(lgu>>1)) - 1)
}

func (m *rsVEBTreeUInt32Mixin) Index(lgu int, high, low interface{}) interface{} {
	return (high.(uint32) << uint32(lgu>>1)) | m.Low(lgu, low).(uint32)
}
