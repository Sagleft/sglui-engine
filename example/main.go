package main

import (
	"fmt"
	"log"

	freemiumengine "github.com/Sagleft/freemium-engine"
)

const appPublic = "AQENI3HF3HLWCMZUXIJL5FHDHK6Q5H4CO75JYWKD5S56RCDX3CKWZ2AJBHCFTHM5YIWQMBVS6XVEHHZFIBZUL7CKLYFS45EHL22RD7DA6NACT2F7JFNOLTPD6Z6QA53KQL6HNGCX7YE7LTKX3CHPUZS7GT4A===="

const appSecret = "FD7YCAYBAEFXA22DN5XHIYLJNZSXEAP7QIAACAQBANIHKYQBBIAACAKEAH7YIAAAAAFP7AYFAEBP7BQAAAAP7GP7QIAWCBAI2RWOLWOXMEZTJOQSX2KOGOV5B2PYE572TRMUH3F35CEHPWEVNTUASCOELGOZ3QRNAYDLF5PKIOPSKQDTIX6EUXQLFZ2IOXVVCH6GB42AFHUL6SK24XG6H5T5AB3WVAX4O2MFP7QJ6XGVPWEO7JTF6NHYAEYQFJBWJADULNDH5SYO4BNW6IFW7RKZ2ZKS6MG5QVRXECNJNBFLYROZCDEAHCX2GKREISXLHKFV5E7O24AA===="

func main() {
	err := freemiumengine.Validate(appPublic)

	fmt.Println(err)
}

func create() {
	if err := freemiumengine.Create(appSecret); err != nil {
		log.Fatalln(err)
	}

	fmt.Println("done")
}