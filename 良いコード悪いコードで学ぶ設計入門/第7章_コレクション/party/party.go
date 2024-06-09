package party

import (
	"errors"
	"fmt"
	"slices"
)

type FieldManager struct {
	Members []Member
}

type Member struct {
	ID      int
	Name    string
	IsAlive bool
}

const MaxMemberCount = 5

func (fm *FieldManager) AddMember(members []Member, newMember Member) error {
	if slices.Contains(members, newMember) {
		return fmt.Errorf("既にパーティに加わっています。")
	}
	if len(fm.Members) >= MaxMemberCount {
		return fmt.Errorf("これ以上メンバーを追加できません。")
	}
	fm.Members = append(fm.Members, newMember)
	return nil
}

func (fm *FieldManager) PartyIsAlive(members []Member) bool {
	return slices.ContainsFunc(members, func(member Member) bool {
		return member.IsAlive
	})
}

type SpecialEventManager struct{}

func (sem *SpecialEventManager) AddMember(members []Member, member Member) []Member {
	members = append(members, member)
	return members
}

type BattleManager struct{}

func (bm *BattleManager) MembersAerAlive(members []Member) bool {
	return slices.ContainsFunc(members, func(member Member) bool {
		return member.IsAlive
	})
}

// Party ファーストクラスコレクション
type Party struct {
	Members []Member
}

func NewParty(members []Member) *Party {
	return &Party{
		Members: members,
	}
}

// Add Addは新しいメンバーをパーティに加えるメソッド
// Goの設計哲学は、イミュータブルなデータ構造を使うことが推奨されている
func (p *Party) Add(newMember Member) (*Party, error) {
	if p.Exists(newMember) {
		return nil, errors.New("既にパーティに加わっています。")
	}
	if p.IsFull() {
		return nil, errors.New("これ以上メンバーを追加できません。")
	}
	newMembers := append([]Member{}, p.Members...)
	newMembers = append(newMembers, newMember)
	return NewParty(newMembers), nil
}

// GetMembers GetMembersはメンバーリストのコピーを返すメソッド
// メンバーリストを直接返すと、呼び出し元でメンバーリストを変更できてしまうため
func (p *Party) GetMembers() []Member {
	membersCopy := make([]Member, len(p.Members))
	copy(membersCopy, p.Members)
	return membersCopy
}

func (p *Party) IsAlive() bool {
	return slices.ContainsFunc(p.Members, func(member Member) bool {
		return member.IsAlive
	})
}

func (p *Party) Exists(member Member) bool {
	return slices.ContainsFunc(p.Members, func(m Member) bool {
		return m.ID == member.ID
	})
}

func (p *Party) IsFull() bool {
	return len(p.Members) >= MaxMemberCount
}
