package main

import "log"

// メッセージ定義
type GetRequest struct{ Key string }
type GetResponse struct{ Value string }
type SetRequest struct{ Key, Value string }

// キャッシュアクターの定義
type CacheActor struct {
	cache map[string]string
}

func (state *CacheActor) Receive(context actor.Context) {
	switch msg := context.Message().(type) {
	case *GetRequest:
		value, ok := state.cache[msg.Key]
		if ok {
			context.Respond(&GetResponse{Value: value})
			// q: context.Respond()の挙動を教えてください。
			// a: context.Respond()は、メッセージを送信したアクターに対して返信を送信します。
		} else {
			context.Respond(&GetResponse{Value: ""})
		}
	case *SetRequest:
		state.cache[msg.Key] = msg.Value
	default:
		log.Println("unknown message", msg)
	}
}
