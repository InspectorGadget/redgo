package structs

type RedisValue struct {
	Key   string `json:"key" binding:"required"`
	Value any    `json:"value" binding:"required"`
}

func (rv RedisValue) GetKey() string {
	return rv.Key
}

func (rv RedisValue) GetValue() any {
	return rv.Value
}
