package config_model

type Upload struct {
	OssType     string `yaml:"OssType"`     //指定类型
	Path        string `yaml:"Path"`        //本地文件访问路径
	StorePath   string `yaml:"StorePath"`   //本地文件存储路径
	MdPath      string `yaml:"MdPath"`      //Markdown 访问路径
	MdStorePath string `yaml:"MdStorePath"` //Markdown 存储路径
}
