package utils

import "github.com/foursking/btype"

// ToMapStringString 结构体转string ignore 是否忽略空值
func ToMapStringString(x interface{}, ignore bool) map[string]string {
	mp := btype.ToMapStringDeep(x)
	mmp := make(map[string]string, len(mp))
	for s, i := range mp {
		is := btype.ToString(i)
		if i != "" || !ignore {
			mmp[s] = is
		}
	}
	return mmp
}
