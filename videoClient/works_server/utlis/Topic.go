package utlis

import (
	"fmt"
	"math"
	"sort"
	"time"
	"works_server/handler/model"
)

// Topic 表示话题数据结构
//type Topic struct {
//	ID         int       // 话题ID
//	Title      string    // 话题标题
//	CreateTime time.Time // 创建时间
//	VideoCount int       // 参与视频数
//	HotScore   float64   // 实时热度分
//}

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

	var videoDate model.VideoDataTopic

	dateTopic, err := videoDate.SumVideoDateTopic(topic.Id)
	if err != nil {
		panic(err)
	}

	var sum int

	for _, dataTopic := range dateTopic {
		if dataTopic.TopicId == topic.Id {
			sum++
		}
	}

	fmt.Println(sum)

	// 2. 使用对数缩放视频参与数量
	logScore := math.Log1p(float64(sum))

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

// 示例使用
func main() {
	// 初始化算法
	algo := NewHotTopicAlgorithm()

	var topics []model.VideoTopic

	// 获取热门话题 (每页10条)
	hotTopics := algo.GetHotTopics(topics, 1, 10)

	// 输出结果
	for _, t := range hotTopics {
		println(t.Id, t.Title, t.HotScore)
	}
}
