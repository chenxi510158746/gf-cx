package v1

import "github.com/gogf/gf/v2/frame/g"

type OcrReq struct {
	g.Meta   `path:"/ocr/do" method:"post" tags:"AI识别" summary:"文字识别"`
	FilePath string `v:"required#参数错误" json:"file_path"`
}

type OcrRes struct {
	List []*OcrItem `json:"list" dc:"识别列表"`
}

type OcrItem struct {
	Text string `json:"text" dc:"识别结果文字"`
}
