package charset

import (
	"github.com/go-playground/assert/v2"
	"testing"
)

func TestDetectGBK(t *testing.T) {
	assert.Equal(t, true, DetectGBK([]byte("³£ÓÃÊýÑ§·ûºÅµÄ LaTeX ±íÊ¾·½·¨")))
	assert.Equal(t, true, DetectGBK([]byte("是GBK")))
	assert.Equal(t, false, DetectGBK([]byte("GBKではありません")))
	assert.Equal(t, true, DetectGBK([]byte("\xaa\xbb")))
	assert.Equal(t, true, DetectGBK([]byte("is_gbk")))
}
