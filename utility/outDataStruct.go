package utility

// 前端请求返回自定义状态代码和信息
type FeedbackMessage struct {
	FeedbackCode int
	FeedbackText string
	Data         interface{}
}

type MaterialsOut struct {
	MaterialCode string
	MaterialName string
	MaterialType string
	MaterialPath string
	Length       string
}

type MaterialOut struct {
	Material
	Size   string
	Length string
}
