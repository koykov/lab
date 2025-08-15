package main

import (
	"fmt"
	"math"
	"math/rand"
	"sort"
)

// OverlappingPermutationsTest проверяет частоту всех возможных перестановок.
func OverlappingPermutationsTest(data []float64, k int) (bool, map[string]int) {
	if len(data) < k {
		return false, nil
	}

	// Генерируем все возможные перестановки для k элементов
	perms := generatePermutations(k)
	permCounts := make(map[string]int)

	// Проходим по всем перекрывающимся группам
	for i := 0; i <= len(data)-k; i++ {
		group := data[i : i+k]
		permKey := getPermutationKey(group, perms)
		permCounts[permKey]++
	}

	// Проверяем, равномерно ли распределены перестановки (χ²-тест)
	expected := float64(len(data)-k+1) / float64(len(perms))
	chi2 := 0.0
	for _, count := range permCounts {
		chi2 += math.Pow(float64(count)-expected, 2) / expected
	}

	// Критическое значение χ² для (n!-1) степеней свободы (α=0.05)
	criticalValue := getChi2CriticalValue(len(perms) - 1)
	isUniform := chi2 < criticalValue

	return isUniform, permCounts
}

// generatePermutations возвращает все возможные перестановки для k элементов.
func generatePermutations(k int) [][]float64 {
	indices := make([]int, k)
	for i := 0; i < k; i++ {
		indices[i] = i
	}

	var perms [][]float64
	for _, p := range permutations(indices) {
		perm := make([]float64, k)
		for i, idx := range p {
			perm[i] = float64(idx)
		}
		perms = append(perms, perm)
	}
	return perms
}

// permutations возвращает все перестановки массива (рекурсивно).
func permutations(arr []int) [][]int {
	var helper func([]int, int)
	res := [][]int{}

	helper = func(arr []int, n int) {
		if n == 1 {
			tmp := make([]int, len(arr))
			copy(tmp, arr)
			res = append(res, tmp)
		} else {
			for i := 0; i < n; i++ {
				helper(arr, n-1)
				if n%2 == 1 {
					arr[0], arr[n-1] = arr[n-1], arr[0]
				} else {
					arr[i], arr[n-1] = arr[n-1], arr[i]
				}
			}
		}
	}
	helper(arr, len(arr))
	return res
}

// getPermutationKey возвращает ключ перестановки.
func getPermutationKey(group []float64, perms [][]float64) string {
	sorted := make([]float64, len(group))
	copy(sorted, group)
	sort.Float64s(sorted)

	for i, perm := range perms {
		if equalSlices(group, perm) {
			return fmt.Sprintf("perm_%d", i)
		}
	}
	return "unknown"
}

// equalSlices проверяет, равны ли два слайса.
func equalSlices(a, b []float64) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

// getChi2CriticalValue возвращает критическое значение χ².
func getChi2CriticalValue(df int) float64 {
	// Для α=0.05 и df=5 (k=3 → 3!-1=5) ≈ 11.07
	chi2Table := map[int]float64{
		1:  3.84,  // k=2 (2!-1=1)
		5:  11.07, // k=3 (3!-1=5)
		23: 35.17, // k=4 (4!-1=23)
	}
	return chi2Table[df]
}

func main() {
	// Генерируем тестовые данные (можно заменить на свои)
	data := make([]float64, 10000)
	for i := range data {
		data[i] = rand.Float64() // Замените на свой ГСЧ
	}

	// Запускаем тест для троек (k=3)
	k := 3
	isUniform, counts := OverlappingPermutationsTest(data, k)

	fmt.Printf("Тест для k=%d:\n", k)
	fmt.Printf("Равномерное распределение? %v\n", isUniform)
	fmt.Printf("Количество каждой перестановки: %v\n", counts)
}
