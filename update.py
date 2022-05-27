import os
import json

a = input("path to the json folder: ")

b = input("cid: ")

c = os.listdir(a)

for i in range(len(c)):
	if c[i].endswith("json"):
		d = a+"/"+c[i]
		e = open(d,"r")
		f = json.load(e)
		e.close()

		f["image"] = "ipfs://"+b+"/"+f["image"]
		#f["description"] = b
		g = open(d,"w")
		json.dump(f,g)
		g.close()
