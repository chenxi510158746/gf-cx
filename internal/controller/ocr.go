package controller

import (
	"context"
	"encoding/json"
	v1 "gf-cx/api/v1"
	"os"
	"os/exec"
)

var Ocr = cOcr{}

type cOcr struct {
}

func (c *cOcr) Do(ctx context.Context, req *v1.OcrReq) (res *v1.OcrRes, err error) {
	outFile := "/data/wwwroot/gf-cx/temp/result.txt"
	cmd := exec.Command("/data/software/anaconda3/envs/paddle_env/bin/python3", "/data/wwwroot/gf-cx/ocr.py", req.FilePath, outFile)
	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}
	defer func() {
		os.Remove(outFile)
	}()
	var d [][][][]any
	err = json.Unmarshal(output, &d)
	if err != nil {
		return nil, err
	}
	var ocrItems []*v1.OcrItem
	if len(d) > 0 {
		for _, d1 := range d {
			if len(d1) > 0 {
				for _, d2 := range d1 {
					text := d2[1][0].(string)
					ocrItems = append(ocrItems, &v1.OcrItem{Text: text})
				}
			}
		}
	}
	return &v1.OcrRes{List: ocrItems}, nil
}
