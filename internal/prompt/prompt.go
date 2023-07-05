package prompt

import "fmt"

const GetLableByInfoPromptMsg = "你是我的关键词联想助手，你的回答必须是一个关键词，不能有其他内容。我会给出一段已有内容，你根据这段内容联想出一个关键词。我给出的内容是%s"
const GetInfoByLableAndInfoPromptMsg = "你是我的关键词总结助手，你的回答必须是总结出的一段话，不能有任何多余内容。我会给出一段已有内容和一个关键词，你需要联想出它们之间的关系。我给出的内容是：%s，我给出的关键词是：%s"

func GetLableByInfoPrompt(info string) (string, error) {
	resp, err := getMessageFromAI(fmt.Sprintf(GetLableByInfoPromptMsg, info))
	if err != nil {
		return "", err
	}
	return resp, nil
}
func GetInfoByLableAndInfoPrompt(label string, info string) (string, error) {
	resp, err := getMessageFromAI(fmt.Sprintf(GetInfoByLableAndInfoPromptMsg, info, label))
	if err != nil {
		return "", err
	}
	return resp, nil
}
