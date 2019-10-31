package jwtauthdomain

import "log"

type jwtRegisterDomainInterface interface {
	Create(string, string) error
	Delete(string) error
	Exist(string) (bool, error)
	Get(string) (*RegisterJwt, bool, error)
}

var (
	//JwtRegister jwt collect
	JwtRegister jwtRegisterDomainInterface

	collectJwts = map[string]string{
		"qZhsF2HfuWZEBghFa4nl2Kidyp22": "gFfq8ut-n-qxlkHiMCT6zusVvkNK2iCg0_xvEi1vBXK-KCWPG1dlyT5UUnuVlj3vHgFj3uERLsQo6fh8ReofN3KjNRomD8uLJlfykFJ6AbczzpxiXLqK8Y57qXsziF-bz78NKQG01RYxxeVgLnq_euPLkPkmOwEB5QayGjD3uTaHb0BphznsXdo4_7xjjW4GTfKxGt4EFq0WE88eEptBV83kWMUrO5jIsMt45vMrreaHOSnSAGDdge7vuEClL-GzcJdr_uaFCAJVEGYbmgV9VLkGu2XEgDJI0plrKLAqJRQBEKvzPtbaBuQVEDTqk-GOYPKrsUCPJryo43dHU4gUSQ",
	}
)

type jwtRegister struct{}

func init() {
	JwtRegister = &jwtRegister{}
}

func (s *jwtRegister) Create(uid string, jwtSign string) error {
	if _, ok := collectJwts[uid]; ok {
		log.Printf("input data exists")
	}
	collectJwts[uid] = jwtSign

	return nil
}

func (s *jwtRegister) Delete(uid string) error {
	if _, ok := collectJwts[uid]; ok {
		delete(collectJwts, uid)
	}
	return nil
}

func (s *jwtRegister) Exist(uid string) (bool, error) {
	if _, ok := collectJwts[uid]; ok {
		return true, nil
	}
	return false, nil
}

func (s *jwtRegister) Get(uid string) (*RegisterJwt, bool, error) {
	if v, ok := collectJwts[uid]; ok {
		return &RegisterJwt{
			Sign: v,
		}, true, nil
	}
	return nil, false, nil
}
