package main

import (
	"fmt"
	"strings"
)

// 6253. 回环句
func isCircularSentence(sentence string) bool {
	arr := strings.Split(sentence, " ")
	n := len(arr)
	for k, curNode := range arr {
		nextNode := arr[(k+1)%n]
		if curNode[len(curNode)-1] != nextNode[0] {
			return false
		}
	}

	return true
}

/**
6254. 划分技能点相等的团队
给你一个正整数数组 skill ，数组长度为 偶数 n ，其中 skill[i] 表示第 i 个玩家的技能点。将所有玩家分成 n / 2 个 2 人团队，使每一个团队的技能点之和 相等 。
团队的 化学反应 等于团队中玩家的技能点 乘积 。
返回所有团队的 化学反应 之和，如果无法使每个团队的技能点之和相等，则返回 -1 。
**/
func dividePlayers(skill []int) int64 {
	var sum int
	var ji int
	n := len(skill)
	numM := make(map[int][]int, n)
	passK := make(map[int]int, n)

	var pair [][2]int

	for k, i := range skill {
		sum += int(i)
		numM[i] = append(numM[i], k)
	}

	fmt.Println(numM)

	avg := int(sum / (n / 2))
	for k, v := range skill {
		fmt.Println("num", v, numM[v], numM)
		_, ex := passK[k]
		// 如果已经被选走了 就直接到下一个了。
		if ex {
			continue
		}

		need := avg - v
		if get(need, numM, passK) {
			if !get(v, numM, passK) {
				return -1
			}

			p := [2]int{v, need}
			pair = append(pair, p)
		} else {
			return -1
		}
	}
	fmt.Println(pair)
	for _, pv := range pair {
		ji += pv[0] * pv[1]
	}
	return int64(ji)
}
func get(need int, numM map[int][]int, passK map[int]int) bool {
	val, ok := numM[need]
	// 如果找到 搭伙的
	if ok {
		// passK[k] = 1
		passK[val[0]] = 1

		if len(val) >= 1 {
			val = val[1:]
			numM[need] = val
		}

		// 如果没有了 就直接删除
		if len(val) == 0 {
			delete(numM, need)
		}
	} else {
		// fmt.Println(k, v, passK, pair, numM)
		return false
	}
	return true
}
func main() {
	speech := tts.Speech{Folder: "audio", Language: "en"}
	speech.Speak("Flying to the moon")

}
