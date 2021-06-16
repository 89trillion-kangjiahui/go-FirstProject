package entity

type Soldier struct {
	Id                       string `json:"id"`
	Name                     string `json:"Name"`
	Note                     string `json:"note"`
	UnlockArena              string `json:"UnlockArena"`
	ArmyType                 string `json:"ArmyType"`
	Race                     string `json:"Race"`
	Rarity                   string `json:"Rarity"`
	Quality                  string `json:"Quality"`
	AttackType               string `json:"AttackType"`
	Prefab                   string `json:"Prefab"`
	PrefabFX                 string `json:"PrefabFX"`
	PrefabEnemyFX            string `json:"PrefabEnemyFX"`
	AttackPattern            string `json:"AttackPattern"`
	Radius                   string `json:"Radius"`
	AvoidancePriority        string `json:"AvoidancePriority"`
	MoveType                 string `json:"MoveType"`
	AtkTargets               string `json:"AtkTargets"`
	MaxHp                    string `json:"MaxHp"`
	Atk                      string `json:"Atk"`
	SplashAtk                string `json:"SplashAtk"`
	AtkRange                 string `json:"AtkRange"`
	ViewRange                string `json:"ViewRange"`
	Width                    string `json:"Width"`
	Length                   string `json:"Length"`
	Def                      string `json:"Def"`
	Heal                     string `json:"Heal"`
	ShootSpeed               string `json:"ShootSpeed"`
	MoveSpeed                string `json:"MoveSpeed"`
	FirstAtkCount            string `json:"FirstAtkCount"`
	AttackBuildingScaleArray string `json:"AttackBuildingScaleArray"`
	ArmyCount                string `json:"ArmyCount"`
	Skill                    string `json:"Skill"`
	BeatBackDistance         string `json:"BeatBackDistance"`
	RecycleMoney             string `json:"RecycleMoney"`
	Desc                     string `json:"Desc"`
	CombatPoints             string `json:"CombatPoints"`
	KillGold                 string `json:"KillGold"`
}

