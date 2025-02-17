package prompts

import "fmt"

// GetPrompts generates a prompt message suitable for the assistant
func GetPrompts(tools string) string {
	// Use a multi-line string for better readability
	prompt := fmt.Sprintf(`
您是一位有用的助手，可以使用这些工具：%s
	根据用户的问题选择适当的工具。
	如果不需要工具，请直接回复。
	重要提示：当您需要使用工具时，您必须仅使用以下确切的 JSON 对象格式进行回复，不能使用其他任何内容：
	{
		"name": "name",
		"arguments": {
			"argument-name": "value"
		}
	}
	收到工具的响应后：
	1. 将原始数据转换为自然的对话响应
	2. 保持响应简洁但信息丰富
	3. 关注最相关的信息
	4. 使用用户问题中的适当上下文
	5. 避免简单重复原始数据
	请仅使用上面明确定义的工具。`, tools)
	return prompt
}
