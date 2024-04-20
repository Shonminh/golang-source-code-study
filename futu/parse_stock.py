# type T struct {
# 	Code    int    `json:"code"`
# 	Message string `json:"message"`
# 	Data    struct {
# 		List []struct {
# 			K  int     `json:"k"`
# 			O  string  `json:"o"`
# 			C  string  `json:"c"`
# 			H  string  `json:"h"`
# 			L  string  `json:"l"`
# 			V  int     `json:"v"`
# 			T  float64 `json:"t"`
# 			R  float64 `json:"r"`
# 			Lc float64 `json:"lc"`
# 			Cp string  `json:"cp"`
# 		} `json:"list"`
# 	} `json:"data"`
# }
import json
import time

if __name__ == '__main__':
    f = open("a.json", 'r')
    load = json.load(f)
    print(load)
    li = load.get('data').get('list')
    print(li)
    start = 1683518400
    sum = 0
    n = 0
    for l in li:
        t = int(l.get('k'))
        # x = time.strftime('%y-%m-%d', t.)
        day = time.strftime("%Y-%m-%d %H:%M:%S", time.localtime(t))
        if t > start:
            # print(l, day)
            n += 1
            sum += float(l.get('o'))

    print(sum / n)

