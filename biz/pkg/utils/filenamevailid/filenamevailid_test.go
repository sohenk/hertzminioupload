package filenamevailid

import "testing"

func TestFilenamevailid(t *testing.T) {
	t.Log(Filenamevailid("sdfsdf.xxx", ".gif,.jpeg,.png,.jpg"))
}
