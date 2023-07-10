package common

import (
	"cqserver/golibs/logger"
	"math/rand"
	"time"
)

/**************************************************************************/
//-----------------------------权重随机-----------------------------------*/
/**************************************************************************/

/**
*随机一个权重道具
*items [][]int  [[道具Id,道具数量，道具权重],[道具Id,道具数量，道具权重]]
*return 道具Id,道具数量
 */
func RandWeightBySlice2(randData [][]int) (int, int) {

	if len(randData) == 0 {
		return -1, 0
	}
	sum := 0
	for _, v := range randData {
		if len(v) != 3 {
			logger.Error("随机道具错误，道具配置错误：%v", randData)
			return -1, 0
		}
		sum += v[2]
	}

	randNum := rand.Intn(sum)
	logger.Debug("sun:%v  randNum:%v", sum, randNum)
	count := 0
	for _, v := range randData {
		count += v[2]
		if randNum < count {
			return v[0], v[1]
		}
	}
	return -1, 0
}

/**
*随机权重
*randData []int [索引]权重
*return 索引
 */
func RandWeightByIntSlice(randData []int) int {

	sum := 0
	for _, v := range randData {
		sum += v
	}

	randNum := rand.Intn(sum)
	count := 0
	for k, v := range randData {
		count += v
		if randNum < count {
			return k
		}
	}
	logger.Error("ErrLog 随机权重出错，没有随机到任何东西：%v", randData)
	return -1
}

/**
*随机权重
*randData map[int]int{索引:权重，索引：权重}
*return 索引
 */
func RandWeightByMap(randData map[int]int) int {

	sum := 0
	for _, v := range randData {
		sum += v
	}
	if sum <= 0 {
		return -1
	}
	randNum := rand.Intn(sum)
	count := 0
	for k, v := range randData {
		count += v
		if randNum < count {
			return k
		}
	}
	logger.Error("ErrLog 随机权重出错，没有随机到任何东西：%v", randData)
	return -1
}

/**
*随机权重
*randData map[int]int{索引:权重，索引：权重}
*return 索引
 */
func RandMultiWeightByMap(randData map[int]int, num int) []int {

	sum := 0
	for _, v := range randData {
		sum += v
	}
	if num > len(randData) {
		num = len(randData)
	}

	rands := make([]int, num)
	for i := 0; i < num; i++ {
		logger.Info("sum:%v，randData:%v,以随机：%v", sum, randData, rands)
		randNum := rand.Intn(sum)
		count := 0
		for k, v := range randData {
			count += v
			if randNum < count {
				rands[i] = k
				delete(randData, k)
				sum -= v
				break
			}
		}
	}

	logger.Error("ErrLog 随机权重出错，没有随机到任何东西：%v", randData)
	return rands
}

/**************************************************************************/
//-----------------------------概率随机-----------------------------------*/
/**************************************************************************/

/**
*随机道具 每个道具概率获取
*items [][]int  [[道具Id,道具数量，道具概率],[道具Id,道具数量，道具概率]]
*return 道具Id,道具数量
 */
func RandRateBySlice2(randData [][]int) map[int]int {

	if len(randData) == 0 {
		return make(map[int]int)
	}
	rand.Seed(time.Now().UnixNano())
	reMap := make(map[int]int)
	for _, v := range randData {
		if len(v) != 3 {
			logger.Error("随机道具错误，道具配置错误：%v", randData)
			return make(map[int]int)
		}
		if rand.Intn(10000) < v[2] {
			reMap[v[0]] += v[1]
		}
	}
	return reMap
}

/**
 *  @Description: 万分比随机
 *  @param rate
 *  @return bool
 */
func RandByTenShousand(rate int) bool {

	randNum := rand.Intn(10000)
	if randNum <= rate {
		return true
	}
	return false
}

/**
*随机一个权重道具
*items [][]int  [[道具Id,道具数量，道具权重,道具索引(对应配置表),type],[道具Id,道具数量，道具权重,道具索引(对应配置表),type]]
*return 道具Id,道具数量,道具索引,type
 */
func RandWeightBySlice3(randData [][]int) (int, int, int, int) {

	if len(randData) == 0 {
		return -1, 0, -1, -1
	}
	sum := 0
	for _, v := range randData {
		if len(v) != 5 {
			logger.Error("随机道具错误，道具配置错误：%v", randData)
			return -1, 0, -1, -1
		}
		sum += v[2]
	}

	randNum := rand.Intn(sum)
	count := 0
	for _, v := range randData {
		count += v[2]
		if randNum < count {
			return v[0], v[1], v[3], v[4]
		}
	}
	return -1, 0, -1, -1
}
