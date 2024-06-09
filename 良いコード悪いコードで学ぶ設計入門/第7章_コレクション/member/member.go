package member

import "slices"

// 早期continueを使うことで、if文のネストを減らすことができる

type StateType string

const (
	Poison StateType = "poison"
	Dead   StateType = "dead"
)

type Member struct {
	Name     string
	HitPoint int
	States   []StateType
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

func main() {
	members := []Member{
		{Name: "勇者", HitPoint: 10},
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
