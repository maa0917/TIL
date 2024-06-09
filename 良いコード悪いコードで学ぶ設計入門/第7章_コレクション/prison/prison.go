package prison

// わざわざ自前でコレクション処理を実装する必要はなく、 標準ライブラリに用意された便利なメソッドを利用すべきということを述べている
// Goの標準ライブラリには、JavaのStream APIに相当するものがない

import (
	"fmt"
	"slices"
)

type Item struct {
	Name string
}

// anyMatch は、items の中に predicate が true を返す要素があるかどうかを返す
// Java の Stream#anyMatch に相当する
func anyMatch(items []Item, predicate func(Item) bool) bool {
	for _, item := range items {
		if predicate(item) {
			return true
		}
	}
	return false
}

func main() {
	items := []Item{
		{Name: "アイテム1"},
		{Name: "牢屋の鍵"},
		{Name: "アイテム2"},
	}

	// 牢屋の鍵を持っているかどうかを調べる
	hasPrisonKey := anyMatch(items, func(item Item) bool {
		return item.Name == "牢屋の鍵"
	})
	fmt.Println("牢屋の鍵を持っているか:", hasPrisonKey)

	// 牢屋の鍵を持っているかどうかを調べる（slices.Contains を使う）
	itemNames := make([]string, len(items))
	for i, item := range items {
		itemNames[i] = item.Name
	}
	hasPrisonKey = slices.Contains(itemNames, "牢屋の鍵")
	fmt.Println("牢屋の鍵を持っているか:", hasPrisonKey)
}
