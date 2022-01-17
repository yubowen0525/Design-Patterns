package factory

type IRuleConfigParser interface {
	Parse(data []byte)
}

type jsonRuleConfigParser struct{}

func (J jsonRuleConfigParser) Parse(data []byte) {
	panic("implement me ")
}

type yamlRuleConfigParser struct{}

func (Y yamlRuleConfigParser) Parse(data []byte) {
	panic("implement me ")
}

type IRuleConfigParserFactory interface {
	CreateParser() IRuleConfigParser
}

type YamlRuleConfigParserFactory struct{}

func (Y YamlRuleConfigParserFactory) CreateParser() IRuleConfigParser {
	return yamlRuleConfigParser{}
}

type JsonRuleConfigParserFactory struct {
}

func (j JsonRuleConfigParserFactory) CreateParser() IRuleConfigParser {
	return jsonRuleConfigParser{}
}

// NewIRuleConfigParserFactory 用一个简单工厂封装工厂方法
func NewIRuleConfigParserFactory(t string) IRuleConfigParserFactory {
	switch t {
	case "json":
		return JsonRuleConfigParserFactory{}
	case "yaml":
		return YamlRuleConfigParserFactory{}
	}
	return nil
}

func test() {
	parse := NewIRuleConfigParserFactory("json")
	parse.CreateParser().Parse([]byte{1, 2})
}
