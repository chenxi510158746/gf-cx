#!/data/software/anaconda3/envs/paddle_env/bin/python3
# -*- coding: UTF-8 -*-

import json
import sys

from paddleocr import PaddleOCR

img_path=sys.argv[1]
out_path=sys.argv[2]
lan=sys.argv[3]
if lan=="korean":
    det_model="/data/software/model/korean/Multilingual_PP-OCRv3_det_infer/"
    rec_model="/data/software/model/korean/korean_PP-OCRv3_rec_infer/"
    cls_model="/data/software/model/korean/ch_ppocr_mobile_v2.0_cls_infer/"
    lan="korean"
elif lan=="japan":
    det_model="/data/software/model/japan/Multilingual_PP-OCRv3_det_infer/"
    rec_model="/data/software/model/japan/japan_PP-OCRv3_rec_infer/"
    cls_model="/data/software/model/japan/ch_ppocr_mobile_v2.0_cls_infer/"
    lan="japan"
else:
    det_model="/data/software/model/ch_PP-OCRv4_det_infer/"
    rec_model="/data/software/model/ch_PP-OCRv4_rec_infer/"
    cls_model="/data/software/model/ch_ppocr_mobile_v2.0_cls_infer/"
    lan="ch"

# Paddleocr目前支持的多语言语种可以通过修改lang参数进行切换
# (lang="ch")例如`ch`, `en`, `fr`, `german`, `korean`, `japan`
ocr = PaddleOCR(det_model_dir=det_model, rec_model_dir=rec_model, cls_model_dir=cls_model, use_angle_cls=True, use_gpu=False, lang=lan, show_log=False)
result = ocr.ocr(img_path, cls=True)
result_json = json.dumps(result, ensure_ascii=False)
print(result_json)
with open(out_path, 'w') as fo:
    fo.write(result_json)
