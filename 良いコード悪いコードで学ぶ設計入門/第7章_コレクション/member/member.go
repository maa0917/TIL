package member

import "slices"

// ループ処理中の条件分岐ネストを減らすために、早期continueや早期breakを使う

type StateType string

const (
	Poison StateType = "poison"
	Dead   StateType = "dead"
)

type Member struct {
	Name                string
	HitPoint            int
	States              []StateType
	AttackPower         int
	TeamAttackSucceeded bool
}

func (m *Member) ContainsState(state StateType) bool {
	return slices.Contains(m.States, state)
}

func (m *Member) AddState(state StateType) {
	m.States = append(m.States, state)
}

func (m *Member) RemoveState(state StateType) {
	for i, s := range m.States {
		if s == state {
			m.States = append(m.States[:i], m.States[i+1:]...)
			return
		}
	}

	// 以下のように書くこともできる
	// index := slices.IndexFunc(m.States, func(s StateType) bool {
	// 	return s == state
	// })
	// if index != -1 {
	// 	m.States = slices.Delete(m.States, index, index+1)
	// }
}

func (m *Member) HasTeamAttackSucceeded() bool {
	return m.TeamAttackSucceeded
}

// 早期continueを使うことで、if文のネストを減らすことができる
func poison() {
	members := []Member{
		{Name: "勇者", HitPoint: 10},
		{Name: "戦士", HitPoint: 20},
		{Name: "僧侶", HitPoint: 30},
		{Name: "魔法使い", HitPoint: 40},
	}
	for _, member := range members {
		if member.HitPoint == 0 {
			continue
		}
		if !member.ContainsState(Poison) {
			continue
		}

		member.HitPoint -= 10

		if member.HitPoint > 0 {
			continue
		}
		member.HitPoint = 0
		member.AddState(Dead)
		member.RemoveState(Poison)
	}
}

// 早期breakを使うことで、if文のネストを減らすことができる
func totalDamage() {
	members := []Member{
		{Name: "勇者", AttackPower: 10},
		{Name: "戦士", AttackPower: 20},
		{Name: "僧侶", AttackPower: 30},
		{Name: "魔法使い", AttackPower: 40},
	}
	totalDamage := 0
	for _, member := range members {
		if !member.HasTeamAttackSucceeded() {
			break
		}
		damage := float64(member.AttackPower) * 1.1
		if damage < 30 {
			break
		}
		totalDamage += int(damage)
	}
}
