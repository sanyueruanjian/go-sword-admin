package utils

type (
	Mode string
)

const (
	ModeDev  Mode = "dev"     //开发模式
	ModeTest Mode = "test"    //测试模式
	ModeProd Mode = "prod"    //生产模式
)
