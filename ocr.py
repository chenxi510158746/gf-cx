#!/data/software/anaconda3/envs/paddle_env/bin/python3
# -*- coding: UTF-8 -*-

import json
import sys

from paddleocr import PaddleOCR

img_path=sys.argv[1]
out_path=sys.argv[2]

# Paddleocr目前支持的多语言语种可以通过修改lang参数进行切换
# (lang="ch")例如`ch`, `en`, `fr`, `german`, `korean`, `japan`
ocr = PaddleOCR(det_model_dir='/data/software/model/ch_PP-OCRv4_det_infer/', 
                rec_model_dir='/data/software/model/ch_PP-OCRv4_rec_infer/', 
                cls_model_dir='/data/software/model/ch_ppocr_mobile_v2.0_cls_infer/',
                 use_angle_cls=True, use_gpu=False, lang="ch", show_log=False)
result = ocr.ocr(img_path, cls=True)
result_json = json.dumps(result, ensure_ascii=False)
print(result_json)
with open(out_path, 'w') as fo:
    fo.write(result_json)
