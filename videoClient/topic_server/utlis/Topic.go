package utlis

import (
	"math"
	"sort"
	"time"
	"topic_server/handler/model"
)

// HotTopicAlgorithm 热度算法控制器
type HotTopicAlgorithm struct {
	decayFactor float64   // 时间衰减因子
	baseTime    time.Time // 基准时间(当前时间)
}

func NewHotTopicAlgorithm() *HotTopicAlgorithm {
	return &HotTopicAlgorithm{
		decayFactor: 0.03, // 经验值
		baseTime:    time.Now().UTC(),
	}
}

// CalculateHotScore 计算单个话题热度分
func (h *HotTopicAlgorithm) CalculateHotScore(topic model.VideoTopic) float64 {
	// 1. 计算时间衰减 (指数衰减模型)

	hours := h.baseTime.Sub(topic.CreatedAt).Hours()
	timeDecay := math.Exp(-h.decayFactor * hours)

	count := SumCount(topic.Id)

	// 2. 使用对数缩放视频参与数量
	logScore := math.Log1p(float64(count[topic.Id]))

	// 3. 计算最终热度分
	return logScore * timeDecay
}

// GetHotTopics 获取热门话题列表 (带分页)
func (h *HotTopicAlgorithm) GetHotTopics(topics []model.VideoTopic, page, pageSize int) []model.VideoTopic {
	// 更新基准时间为当前时间
	h.baseTime = time.Now().UTC()

	// 计算所有话题的实时热度分
	for i := range topics {
		topics[i].HotScore = h.CalculateHotScore(topics[i])
	}

	// 按热度降序排序
	sort.Slice(topics, func(i, j int) bool {
		return topics[i].HotScore > topics[j].HotScore
	})

	// 分页处理
	start := (page - 1) * pageSize
	if start >= len(topics) {
		return []model.VideoTopic{}
	}
	end := start + pageSize
	if end > len(topics) {
		end = len(topics)
	}

	return topics[start:end]
}

func SumCount(id int32) map[int32]int {
	var videoDate model.VideoDataTopic

	var Sum map[int32]int
	var count int

	dateTopic, err := videoDate.SumVideoDateTopic(id)
	if err != nil {
		panic(err)
	}
	for _, dataTopic := range dateTopic {
		if dataTopic.TopicId == id {
			count++
			Sum = map[int32]int{
				id: count,
			}
		}
	}
	return Sum
}
