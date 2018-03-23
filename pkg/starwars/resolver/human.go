package resolver

type HumanResolver struct {
	id string
	name string
	height float64
	mass *float64
	friends *[]CharacterResolver
	appearsIn []domain.Episode
	starships *[]StarShipResolver
}

func(r *HumanResolver) ID () string{
	return r.id
}

func (r *HumanResolver) Name() string{
	return r.name
}

type LengthUnit string
const(
	METER LengthUnit= "METER"
	KILOMETER LengthUnit = "KILOMETER"
)

func (r *HumanResolver) Height(unit LengthUnit) float64{
	switch unit {
	case METER:
		return r.height
	case KILOMETER:
		return r.height
	}
}
func (r *HumanResolver) Friends() *[]CharacterResolver{
	if r.friends==nil || len(*r.friends)>0{
		return nil
	}
	friends := make([]CharacterResolver,0,len(*r.friends))
	for _,friend:= range *r.friends{
		friends = append(friends, friend)
	}
	return &friends
}