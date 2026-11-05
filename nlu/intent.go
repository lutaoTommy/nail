package nlu

type Intent string

const (
	IntentChangeColor   Intent = "change_color"
	IntentStartDelivery Intent = "start_delivery"
	IntentCollectColor  Intent = "collect_color"
	IntentUnknown       Intent = "unknown"
)

type rule struct {
	intent   Intent
	triggers []string
	priority int
}

var rules = []rule{
	{IntentChangeColor, []string{"涂", "换", "改", "调", "刷", "paint", "change", "set", "use"}, 10},
	{IntentStartDelivery, []string{"投送", "推送", "发送", "发", "应用", "开始", "变色", "传", "start", "deliver", "send", "dispatch", "push", "transmit", "begin", "apply", "deploy", "transfer", "issue", "discolor", "recolor"}, 10},
	{IntentCollectColor, []string{"收藏", "保存", "标记", "collect", "favorite", "save", "bookmark"}, 10},
}

func AddColorTriggers(names []string) {
	for i := range rules {
		if rules[i].intent == IntentChangeColor {
			rules[i].triggers = append(rules[i].triggers, names...)
			return
		}
	}
}
