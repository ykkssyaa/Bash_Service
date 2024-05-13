package consts

import "time"

const (
	CtxTimeout     = 5 * time.Minute
	ReadOutputTime = 1 * time.Second
)

/*
CmdPullSize Первоначальный размер хранилища.
Задаётся для уменьшения количества эвакуаций мапы в начале работы сервиса.
При увеличении нагрузок можно изменить значение в дальнейшем
*/
const CmdPullSize = 10
