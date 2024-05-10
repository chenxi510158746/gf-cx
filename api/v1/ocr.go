package v1

import (
	"gf-cx/internal/model"

	"github.com/gogf/gf/v2/frame/g"
)

type OcrReq struct {
	g.Meta   `path:"/ocr/do" method:"post" tags:"AI识别" summary:"文字识别"`
	FilePath string `v:"required#参数错误" json:"file_path"`
}

type OcrRes struct {
	List []*model.OcrItem `json:"list" dc:"识别列表"`
}
