/**
 * @Author: lixiang
 * @Date: 2025/3/11 21:07
 * @Description: code.go
 */

package result

type Codes struct {
	Message map[uint]string
	Success uint
	Failed  uint
}

var ApiCode = &Codes{
	Success: 200,
	Failed:  500,
}

func init() {
	ApiCode.Message = map[uint]string{
		ApiCode.Success: "success",
		ApiCode.Failed:  "failed",
	}
}
func (c *Codes) getMsg(code uint) string {
	message := c.Message[code]
	return message
}
