package match

// match m a
func MatchArr(pattern string, data []string, limit ...int) []string {
	weight := make([]float64, 0, 5)
	res := make([]string, 0, 5)

	// start init
	res = append(res, "")
	weight = append(weight, 0)
	length := 0

	for _, content := range data {
		w := matchWeight(pattern, content)
		if w > 0 {
			// insert
			res = append(res, content)
			weight = append(weight, w)
			length++

			// swim
			j := length
			for {
				parent := j / 2
				if weight[parent] <= w {
					break
				}

				if parent == 0 {
					break
				}

				// switch
				weight[j], weight[parent] = weight[parent], weight[j]
				res[j], res[parent] = res[parent], res[j]
				j = parent

			}

		}

	}

	// fetch by order
	for i := 1; i <= length; i++ {
		// switch root
		weight[1], weight[length-i+1] = weight[length-i+1], weight[1]
		res[1], res[length-i+1] = res[length-i+1], res[1]

		j := 1
		for {
			// sink
			childLeft := j * 2
			childRight := childLeft + 1

			// two nodes
			if childRight <= length-i {
				if weight[childLeft] <= weight[childRight] && weight[j] > weight[childLeft] {
					// switch
					weight[childLeft], weight[j] = weight[j], weight[childLeft]
					res[childLeft], res[j] = res[j], res[childLeft]
					j = childLeft
					continue
				}

				if weight[childLeft] <= weight[childRight] && weight[j] > weight[childRight] {
					// switch
					weight[childRight], weight[j] = weight[j], weight[childRight]
					res[childRight], res[j] = res[j], res[childRight]
					j = childRight
					continue
				}

				break

			}

			// one node
			if childLeft <= length-i && weight[j] > weight[childLeft] {
				// switch
				weight[childLeft], weight[j] = weight[j], weight[childLeft]
				res[childLeft], res[j] = res[j], res[childLeft]
				j = childLeft

				continue
			}

			// none
			break
		}

	}

	return res[1:]
}

// ab  axxxxb
func matchWeight(pattern string, content string) (weight float64) {
	p := []byte(pattern)
	c := []byte(content)

	lenp := len(p)
	lenc := len(c)

	// parttern
	// content
	// (lenc - posb)/(lenc-posa) / lenp * lenp / lenc
	i := 0
	j := 0
	for {
		if i >= lenp || j >= lenc {
			break
		}

		if p[i] == c[j] {
			delta := float64(lenc-j) / float64(lenc-i) / float64(lenc)
			weight = weight + delta
			i++
			j++

		} else {
			j++
		}

	}

	if i < lenp {
		weight = 0
	}

	return
}
