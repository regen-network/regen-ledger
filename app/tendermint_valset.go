package app

import (
	tmcrypto "github.com/tendermint/tendermint/crypto"
	tmjson "github.com/tendermint/tendermint/libs/json"
)

const mainnetValidatorSet = `{
	"pub_keys":[
		{"type":"tendermint/PubKeyEd25519","value":"dmRWr9awPjLmuwFoCr9bWmRhYLmva2EuVWxnngL8B8M="},
		{"type":"tendermint/PubKeyEd25519","value":"DC2csM0Nb+dbZfdPMKZXqQOKpifkr4NaDF7B7fcgaRg="},
		{"type":"tendermint/PubKeyEd25519","value":"oceaCqELoCYevpgttA9XXbAHx7D86Traubeg93sBa4o="},
		{"type":"tendermint/PubKeyEd25519","value":"dNO95TeRS9oyZgluL+wETqZE6IVZY/vypyTF7RDbYMg="},
		{"type":"tendermint/PubKeyEd25519","value":"8sCaQ9aButHI/pOltYdMxO8O32uzmI5nvL/EI/Zf7Y0="},
		{"type":"tendermint/PubKeyEd25519","value":"QgF008QqXTdoaI/SaoGirg0FIE3NL1aa3qUlHvVc++g="},
		{"type":"tendermint/PubKeyEd25519","value":"1S+ACbTEumVNXH2/UoRZ488Lnbt3NY/0yAM5scWNGkk="},
		{"type":"tendermint/PubKeyEd25519","value":"51scl1HwS6oniJs79RDD3KRZTYtbRAycd4jJZMD/dsU="},
		{"type":"tendermint/PubKeyEd25519","value":"zRQaHoEC2aP+6/XH3cPuTvapuUeYhWr/y4gsTBsCsvE="},
		{"type":"tendermint/PubKeyEd25519","value":"4ZO9Oj2esqGYQA6KiND6MYa8X/TCq3Xc4zFi1wytSHE="},
		{"type":"tendermint/PubKeyEd25519","value":"y3Oto59qVrxP5YxacgxRjKxRrCLhchSNrnf9nGd2/FU="},
		{"type":"tendermint/PubKeyEd25519","value":"GdzHU8uMEOeuTi8N7IMZ/D3j/qeJr/8Ax7Rwq+pcHMc="},
		{"type":"tendermint/PubKeyEd25519","value":"UapA+x2nsBedTt+akcrg802GWtVv+xg45jFjDFf+wnI="},
		{"type":"tendermint/PubKeyEd25519","value":"trwfTyf0/ZQrEmaEQ0HIQ/ATYpAKGD7QZnbsrkPklHM="},
		{"type":"tendermint/PubKeyEd25519","value":"d4dYN4hkh12tHyjGHe7fQqoIpf5biexmMfQIQdxEjjA="},
		{"type":"tendermint/PubKeyEd25519","value":"P7nV2cdGK+GuugajXNZSsk1kHhgXimLosSVNvOCKkf0="},
		{"type":"tendermint/PubKeyEd25519","value":"wUfvfu/8iimS6ujoBKFlAY2y43wuiVX2L4SRfu2jUak="},
		{"type":"tendermint/PubKeyEd25519","value":"u9Stwssek8x7wtqru3ULntBj8ExaK1fxz9Qwv/+YC9w="},
		{"type":"tendermint/PubKeyEd25519","value":"Xm9K1+YyWxzF7fphSqAYZFTLeBVkGVY2LUiw+3u/PO8="},
		{"type":"tendermint/PubKeyEd25519","value":"NB4b0PhWjB3a+HfTfem8ni8Dl+cPVvHLFFOuDQm6O38="},
		{"type":"tendermint/PubKeyEd25519","value":"mDI5ZqbX7/9YYO/zfzu85zF04ak5MXnERujZy5w7jXk="},
		{"type":"tendermint/PubKeyEd25519","value":"gDWGKPWna8KyMQQoHe5v+RhHtTpE9uIDFwuSsE4LeLw="},
		{"type":"tendermint/PubKeyEd25519","value":"h6CtQjnxCEFvkHb67WJSwRTmvx9Gxpho+DI65AlpkfU="},
		{"type":"tendermint/PubKeyEd25519","value":"nE3MD28jzjqro3lQkcidGrADk6177iIH6oun2Q4zGCY="},
		{"type":"tendermint/PubKeyEd25519","value":"NsXQczt2U63VIrPmLP6hcDMcnjvux9PFVOPcJq5PIJM="},
		{"type":"tendermint/PubKeyEd25519","value":"i/9JBBXJIM1nBVvUflwHZzxr8SxjZ6JAMyrr4gxk4l8="},
		{"type":"tendermint/PubKeyEd25519","value":"54mnWakAy8Hp5CE98dFO8Djs394D2zGeQOxtIiZeFyo="},
		{"type":"tendermint/PubKeyEd25519","value":"c0DXximsIdIfWyTPaHoWUT598u9vwphrwk+SK5Xs0FQ="},
		{"type":"tendermint/PubKeyEd25519","value":"Joz8y7LSpSAaNlyWA7eOzSqZelSI/qMlHvqYxJQ0MJk="},
		{"type":"tendermint/PubKeyEd25519","value":"o+rsN3srSoUiHhRI44lCSDMniSDDnV61f6pl2Jp4j0w="},
		{"type":"tendermint/PubKeyEd25519","value":"h9QUiBrtWNGSLNM/2hb9aQDou8w0oFRQ6SOX1fipBdo="},
		{"type":"tendermint/PubKeyEd25519","value":"hLfP4kC2jh00RkVqT2jaMgyE2kw3UiuXEmggkcnW44I="},
		{"type":"tendermint/PubKeyEd25519","value":"Cg1mhAkAne/RhrEPQ8ecFPSg4dF3nuU0IlinoSoWw3w="},
		{"type":"tendermint/PubKeyEd25519","value":"ikrvSs0ECtOyUMgRhWY0UWEsrCCCLOCZb+adVFIO9dI="},
		{"type":"tendermint/PubKeyEd25519","value":"Zg7IUTroLWX77DNJfSRql+rVIgKTSgWOV+YGVylYwOQ="},
		{"type":"tendermint/PubKeyEd25519","value":"dbBRl6KnM3SRmlRuzudG4N6WFK9YbEyC93SowtB3cZs="},
		{"type":"tendermint/PubKeyEd25519","value":"bfZkRX1n74T08pt0UsWTkMAMAYig0FQlnSbg4A24oOY="},
		{"type":"tendermint/PubKeyEd25519","value":"KTQ507Dvk2+IPYAn0Q5B3yc1+eCOlMi5AZvm9cSFbzo="},
		{"type":"tendermint/PubKeyEd25519","value":"rLiATS7pCxvV/3cE6aId/rGp820go83W5YxdHVhDTWw="},
		{"type":"tendermint/PubKeyEd25519","value":"0DQYnKQYmSSYYSW4jdxhDVLufsGUn8uXdLmEO8Rl/vc="},
		{"type":"tendermint/PubKeyEd25519","value":"/z2cdjMmwsm1CxHMui5CXetppGbYZC62GIC2wtIpB00="},
		{"type":"tendermint/PubKeyEd25519","value":"kq9DoBE0XpRQyJGPJaIZfasyEIVncys29RUfpsYEV6A="},
		{"type":"tendermint/PubKeyEd25519","value":"3tg4r0RlcNLEu6oHEB4Yz0cAs90pcavSXtHyXs+VR+Q="},
		{"type":"tendermint/PubKeyEd25519","value":"M1HTjAbKbbD7RirqeCxtrlfbcZrbsxpVLlrzt2Zvgz8="},
		{"type":"tendermint/PubKeyEd25519","value":"rNCqe3r4D3OhqzkXR7WYsTKGv0ML9uDa/0E8PDgs8RM="},
		{"type":"tendermint/PubKeyEd25519","value":"L6eM0THwZKt17HoKUAZw/6XKZpQGCiHJb6QjUxyG8f8="},
		{"type":"tendermint/PubKeyEd25519","value":"raAybBPv2WfkJw9a+m9C0Nj/k46T+C49VfSPF3MoRto="},
		{"type":"tendermint/PubKeyEd25519","value":"YkMQfckILjJkTrDbeJItcebMgp/F1z4xhZfsFLKh7mw="},
		{"type":"tendermint/PubKeyEd25519","value":"HCr40jINXYDb64I3Lj5xIWIFPrN2B+zkjteX8O1FxB8="},
		{"type":"tendermint/PubKeyEd25519","value":"XYW73g4iZOxuy2/2CcdC5G285JS7TNLzb4bBUKKtEfg="},
		{"type":"tendermint/PubKeyEd25519","value":"Jzs7yEqITdGHshBh6OReFxzx1M/opErIiCWJkArEeEk="},
		{"type":"tendermint/PubKeyEd25519","value":"iaGs5a5UFTkh8hsuDURnuRppSWekRMQvq3bCE+fCO64="},
		{"type":"tendermint/PubKeyEd25519","value":"E3CCBjqvutXKQwRvtHdqnlcJ/MH5Vtin0Nfscre1MYI="},
		{"type":"tendermint/PubKeyEd25519","value":"fNMHX1YUd2PQrZovxfINY06x/YNRZLpchHng/SjcDYc="},
		{"type":"tendermint/PubKeyEd25519","value":"OzSbiS3I3TRCi9//0vY0qaccxS7e0i+B0aGXA0ml2lQ="},
		{"type":"tendermint/PubKeyEd25519","value":"wRjCi50CmKUBK44XXbXuBaxfrlcfsg0gqxLQiOV1yqc="},
		{"type":"tendermint/PubKeyEd25519","value":"3117JVUfZDZz29xe9KyLcnzZsosMi67nPCAeGeyPG4Q="},
		{"type":"tendermint/PubKeyEd25519","value":"nUlsZa2Sxc7uK2xDWLMB4ie0WhBDN/P/AgEN28P0mrY="},
		{"type":"tendermint/PubKeyEd25519","value":"Qlmx+IbOKWoQHtYbVtsFJTY1wfn1ZHzzH/kZhIMGZPk="},
		{"type":"tendermint/PubKeyEd25519","value":"sdUB/fUOO7Eu7n/sYlYBHcFuEVu5uQfrosGRSR8CB9I="},
		{"type":"tendermint/PubKeyEd25519","value":"b+2wn68d90woJ8cfoQpRXCJ9bMdfv3eGIz7ZeFafXNE="},
		{"type":"tendermint/PubKeyEd25519","value":"8VXXhh5hQ62TmkAoeB0RFpbjcBenSDWzEgx+p0ND8Ps="},
		{"type":"tendermint/PubKeyEd25519","value":"gCpeXCFwjVlx0itOc5gCypG2Nq4k1lCQZc8krGCoyOs="},
		{"type":"tendermint/PubKeyEd25519","value":"L8ohTG7fW50WlpX8gsV0iKwo/nZ0x8K8eKPutt2zI6w="},
		{"type":"tendermint/PubKeyEd25519","value":"WxPVrg4uL0TB/QKBCaWOceY37zgvZzdeXCa79aCpX1A="},
		{"type":"tendermint/PubKeyEd25519","value":"WsQZW6j2m5OfHiF4Z+A5VSWmIgXD8TCN+u0cLENGhsk="},
		{"type":"tendermint/PubKeyEd25519","value":"u5T+9Z/3WffP2aWQNrvlYugh+FRNqy9clfFBgQJq4Lc="},
		{"type":"tendermint/PubKeyEd25519","value":"7FPRQ78CxHfncvwMPSRolvUki1XOpNx8GOETIzKbsA4="},
		{"type":"tendermint/PubKeyEd25519","value":"hqGqqsfoHst5rkaUbQ1EJvIwa+mMkPO/YVLKLQiiJJU="},
		{"type":"tendermint/PubKeyEd25519","value":"Qio/sUXCoQ44F2EUhW2woRlEjGEVqsErUV2wBHAAyfY="},
		{"type":"tendermint/PubKeyEd25519","value":"9SJKKAZuEXMsBHffx5Hb/3quT1R+1XrppNCzoFuFMSk="},
		{"type":"tendermint/PubKeyEd25519","value":"1mKn3yFKq1OFpK7vNgT5TfL16sAun9E/6KJmIt7qiUI="},
		{"type":"tendermint/PubKeyEd25519","value":"/HaF0v0BxPf11/syDxNOwsA+qDNiXNz17/gSTQCqosg="},
		{"type":"tendermint/PubKeyEd25519","value":"88SLxwK1oMmMUdWyGYHjovQWEsrlxAxRjsjzhtRwMHk="},
		{"type":"tendermint/PubKeyEd25519","value":"6mWcJcqb2h2PT6s6rlqQ+UlEbQ09+eZikwmYxMGKn5E="}
	]
}`

const redwoodValidatorSet = `{
	"pub_keys":[
		{"type": "tendermint/PubKeyEd25519","value": "vG+k+Nqk+tQoQgIsxelgJ5lwB289TUXOwVa/B8HMpcI="},
		{"type": "tendermint/PubKeyEd25519","value": "nUzWRiQMZEGTXgQF3jRfHqJMrboyFvGBsZE+uNtY21o="},
		{"type": "tendermint/PubKeyEd25519","value": "l3OqDLF58qqT+OLy2noStzatGnip1BOAmQG6tCdKDYk="},
		{"type": "tendermint/PubKeyEd25519","value": "5tSvpZNDQTHzuVDEDXDmvuYlwbNJT7bjn4sDtw+MMII="},
		{"type": "tendermint/PubKeyEd25519","value": "HrfR/uhvLiE2px33N11v2WHbNY2dWZOU1Yr+XC02pms="},
		{"type": "tendermint/PubKeyEd25519","value": "XilNKNUiLrN2hR/IwwLAY4P7w26rOkP7DwYyS4kCeMU="},
		{"type": "tendermint/PubKeyEd25519","value": "FDlx1MBigiI53bMpa3rCvqEgSU9zZoc3CH7Ty5pVJfg="}
	]	
}`

type PubKeys struct {
	Pubkeys []tmcrypto.PubKey `json:"pub_keys"`
}

func GetRedwoodValidatorSetPubkeys() (*PubKeys, error) {
	pks := &PubKeys{}
	if err := tmjson.Unmarshal([]byte(redwoodValidatorSet), pks); err != nil {
		return nil, err
	}

	return pks, nil
}

func GetMainnetValidatorSetPubkeys() (*PubKeys, error) {
	pks := &PubKeys{}
	if err := tmjson.Unmarshal([]byte(mainnetValidatorSet), pks); err != nil {
		return nil, err
	}

	return pks, nil
}
