package main

import (
	"sort"
	"strings"
)

/*
=== Поиск анаграмм по словарю ===

Напишите функцию поиска всех множеств анаграмм по словарю.
Например:
'пятак', 'пятка' и 'тяпка' - принадлежат одному множеству,
'листок', 'слиток' и 'столик' - другому.

Входные данные для функции: ссылка на массив - каждый элемент которого - слово на русском языке в кодировке utf8.
Выходные данные: Ссылка на мапу множеств анаграмм.
Ключ - первое встретившееся в словаре слово из множества
Значение - ссылка на массив, каждый элемент которого, слово из множества. Массив должен быть отсортирован по возрастанию.
Множества из одного элемента не должны попасть в результат.
Все слова должны быть приведены к нижнему регистру.
В результате каждое слово должно встречаться только один раз.

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {

}

func getAnagrams(words []string) map[string][]string{
	anagrams := map[string]string{

	}
	resultMap := map[string]map[string]struct{}{

	}

	result := map[string][]string{}

	for _, w := range words{
		w = strings.ToLower(w)

		runes := []rune(w)
		sort.SliceStable(runes, func(i, j int) bool {
			return runes[i] > runes[j]
		})
		anag := string(runes)
		first, ok := anagrams[anag]
		if !ok{
			anagrams[anag] = w
			first = w
		}
		_, ok = resultMap[first]
		if !ok{
			resultMap[first] = make(map[string]struct{})
		}
		resultMap[first][w] = struct{}{}
	}

	for k, v := range resultMap{
		if len(v) == 1{
			continue
		}
		strs := make([]string, 0, len(v))
		for k := range v{
			strs = append(strs, k)
		}
		sort.Strings(strs)
		result[k] = strs
	}
	
	return result
}