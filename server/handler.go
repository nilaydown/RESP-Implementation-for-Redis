package server

import (
	"errors"
	"strconv"
	"strings"
	"time"

	"github.com/nilayrajderkar/redis-implementation/resp"
)

// DefaultStore is the global key-value store used by the command handler.
var DefaultStore = NewStore()

// HandleRequest takes a serialized RESP string, deserializes it,
// processes the command, and returns a serialized response.
func HandleRequest(input string) string {
	deserialized, err := resp.Deserialize(input)
	if err != nil {
		return resp.Serialize(err)
	}

	array, ok := deserialized.(*[]interface{})
	if !ok {
		return resp.Serialize(errors.New("invalid command format"))
	}

	if len(*array) == 0 {
		return resp.Serialize(errors.New("empty command"))
	}

	command, ok := (*array)[0].(*string)
	if !ok {
		return resp.Serialize(errors.New("command must be a string"))
	}

	args := (*array)[1:]

	switch strings.ToUpper(*command) {
	case "PING":
		return resp.Serialize("PONG")

	case "ECHO":
		if len(args) < 1 {
			return resp.Serialize(errors.New("ECHO requires an argument"))
		}
		if str, ok := args[0].(*string); ok {
			return resp.Serialize(*str)
		}
		return resp.Serialize(errors.New("ECHO argument must be a string"))

	case "SET":
		return handleSet(args)

	case "GET":
		return handleGet(args)

	case "DEL":
		return handleDel(args)

	case "EXISTS":
		return handleExists(args)

	case "TTL":
		return handleTTL(args)

	case "INCR":
		return handleIncr(args, 1)

	case "DECR":
		return handleIncr(args, -1)

	case "KEYS":
		return handleKeys(args)

	default:
		return resp.Serialize(errors.New("Unknown command '" + *command + "'"))
	}
}

// handleSet implements SET key value [EX seconds].
func handleSet(args []interface{}) string {
	if len(args) < 2 {
		return resp.Serialize(errors.New("SET requires at least 2 arguments: key value"))
	}

	key, ok := args[0].(*string)
	if !ok {
		return resp.Serialize(errors.New("SET key must be a string"))
	}
	value, ok := args[1].(*string)
	if !ok {
		return resp.Serialize(errors.New("SET value must be a string"))
	}

	var ttl time.Duration

	// Parse optional flags: EX seconds
	for i := 2; i < len(args); i++ {
		flag, ok := args[i].(*string)
		if !ok {
			continue
		}
		switch strings.ToUpper(*flag) {
		case "EX":
			if i+1 >= len(args) {
				return resp.Serialize(errors.New("EX requires a numeric argument"))
			}
			secStr, ok := args[i+1].(*string)
			if !ok {
				return resp.Serialize(errors.New("EX value must be a string-encoded integer"))
			}
			sec, err := strconv.Atoi(*secStr)
			if err != nil || sec <= 0 {
				return resp.Serialize(errors.New("EX value must be a positive integer"))
			}
			ttl = time.Duration(sec) * time.Second
			i++ // skip the numeric arg
		}
	}

	DefaultStore.Set(*key, *value, ttl)
	return resp.Serialize("OK")
}

func handleGet(args []interface{}) string {
	if len(args) < 1 {
		return resp.Serialize(errors.New("GET requires a key argument"))
	}
	key, ok := args[0].(*string)
	if !ok {
		return resp.Serialize(errors.New("GET key must be a string"))
	}

	val, found := DefaultStore.Get(*key)
	if !found {
		// RESP null bulk string
		return "$-1\r\n"
	}
	return resp.Serialize(val)
}

func handleDel(args []interface{}) string {
	if len(args) < 1 {
		return resp.Serialize(errors.New("DEL requires at least one key"))
	}

	keys := make([]string, 0, len(args))
	for _, a := range args {
		if s, ok := a.(*string); ok {
			keys = append(keys, *s)
		}
	}
	deleted := DefaultStore.Del(keys...)
	return resp.Serialize(deleted)
}

func handleExists(args []interface{}) string {
	if len(args) < 1 {
		return resp.Serialize(errors.New("EXISTS requires at least one key"))
	}

	keys := make([]string, 0, len(args))
	for _, a := range args {
		if s, ok := a.(*string); ok {
			keys = append(keys, *s)
		}
	}
	count := DefaultStore.Exists(keys...)
	return resp.Serialize(count)
}

func handleIncr(args []interface{}, delta int) string {
	if len(args) < 1 {
		return resp.Serialize(errors.New("INCR/DECR requires a key argument"))
	}
	key, ok := args[0].(*string)
	if !ok {
		return resp.Serialize(errors.New("INCR/DECR key must be a string"))
	}

	val, err := DefaultStore.Incr(*key, delta)
	if err != nil {
		return resp.Serialize(err)
	}
	return resp.Serialize(val)
}

func handleTTL(args []interface{}) string {
	if len(args) < 1 {
		return resp.Serialize(errors.New("TTL requires a key argument"))
	}
	key, ok := args[0].(*string)
	if !ok {
		return resp.Serialize(errors.New("TTL key must be a string"))
	}
	return resp.Serialize(DefaultStore.TTL(*key))
}

func handleKeys(args []interface{}) string {
	pattern := "*"
	if len(args) > 0 {
		if s, ok := args[0].(*string); ok {
			pattern = *s
		}
	}
	keys := DefaultStore.Keys(pattern)
	result := make([]interface{}, len(keys))
	for i, k := range keys {
		result[i] = k
	}
	return resp.Serialize(result)
}
