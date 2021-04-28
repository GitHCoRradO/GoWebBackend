package request

type HeadForm struct {
	Rate int `header:"Rate"`
	Domain string `header:"Domain"`
}
