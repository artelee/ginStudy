package api

func GetWelcomeMessage(id string) string {
	print(`GetWelcomeMessage ============`)
	if id == `해리포터` {
		return `살아남은 아이의 접속을 환영합니다.`
	} else {
		return `머글은 돌아가세요.`
	}
}
