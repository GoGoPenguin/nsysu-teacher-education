package redis

import (
	"errors"

	"github.com/gomodule/redigo/redis"
)

// Close redis connection
func (r *Dao) Close() {
	r.conn.Close()
}

// Exists returns if key exists
func (r *Dao) Exists(key string) (bool, error) {
	return redis.Bool(r.conn.Do("EXISTS", key))
}

// Expire set a timeout on key
func (r *Dao) Expire(key string, ttl int) error {
	_, err := r.conn.Do("EXPIRE", key, ttl)

	return err
}

// SetEx sets a key-value pair with expire time
func (r *Dao) SetEx(key string, ttl int, value interface{}) error {
	_, err := r.conn.Do("SETEX", key, ttl, value)

	return err
}

// Del removes the specified keys. A key is ignored if it does not exist.
// Delete multiple keys should use []string for parameter {key}
func (r *Dao) Del(key interface{}) (err error) {
	switch key.(type) {
	case []string:
		_, err = r.conn.Do("DEL", redis.Args{}.AddFlat(key)...)
	default:
		_, err = r.conn.Do("DEL", key)
	}
	return
}

// Get gets value of given key
func (r *Dao) Get(key string) (string, error) {
	return redis.String(r.conn.Do("GET", key))
}

// MGet Returns the values of all specified keys. For every key that
// does not hold a string value or does not exist, the special value nil is returned.
// Because of this, the operation never fails.
func (r *Dao) MGet(key []string) ([]string, error) {
	var result []string

	value, err := r.conn.Do("MGET", redis.Args{}.AddFlat(key)...)

	if err != nil {
		return nil, err
	}

	switch value := value.(type) {
	case []interface{}:
		result = make([]string, len(value))
		for i := 0; i < len(value); i++ {
			switch v := value[i].(type) {
			case []uint8:
				result[i] = string(v)
			}
		}
	default:
		return nil, errors.New("unsupported type from redis MGET command")
	}

	return result, nil
}

// HSet sets field in the hash stored at key to value
func (r *Dao) HSet(key string, field string, value interface{}) error {
	_, err := r.conn.Do("HSET", key, field, value)
	return err
}

// HGet gets value of a specific field of key
func (r *Dao) HGet(key string, field string) (string, error) {
	return redis.String(r.conn.Do("HGET", key, field))
}

// HDel Removes the specified fields from the hash stored at key
func (r *Dao) HDel(key string, fieldValue interface{}) (err error) {
	switch fieldValue.(type) {
	case []string, []int, []int32, []int64, []float32, []float64, []interface{}:
		_, err = r.conn.Do("HDEL", redis.Args{}.Add(key).AddFlat(fieldValue)...)
	default:
		_, err = r.conn.Do("HDEL", key, fieldValue)
	}
	return
}

// HMSet sets multiple fields for a specific key
func (r *Dao) HMSet(key string, fieldValue interface{}) error {
	_, err := r.conn.Do("HMSET", redis.Args{}.Add(key).AddFlat(fieldValue)...)
	return err
}

//HMGet returns the values associated with the specified fields in the hash stored at key.
func (r *Dao) HMGet(key string, fields interface{}) ([]string, error) {
	var result []string

	value, err := r.conn.Do("HMGET", redis.Args{}.Add(key).AddFlat(fields)...)

	if err != nil {
		return nil, err
	}

	switch value := value.(type) {
	case []interface{}:
		result = make([]string, len(value))
		for i := 0; i < len(value); i++ {
			result[i] = string(value[i].([]uint8))
		}
	default:
		return nil, err
	}

	return result, nil
}

// HGetAll get all fields for a given key
func (r *Dao) HGetAll(key string) (map[string]string, error) {
	var result map[string]string

	value, err := r.conn.Do("HGETALL", key)

	if err != nil {
		return nil, err
	}

	switch value := value.(type) {
	case []interface{}:
		result = make(map[string]string, len(value))
		for i := 0; i < len(value); i += 2 {
			result[string(value[i].([]uint8))] = string(value[i+1].([]uint8))
		}
	}

	return result, nil
}

// SAdd add member to a set called key.  If key does not exist, a new set is created.
// param 'member' support array of string, int, int32, int64, float32, float64, interface{}
func (r *Dao) SAdd(key string, member interface{}) (err error) {
	switch member.(type) {
	case []string, []int, []int32, []int64, []float32, []float64, []interface{}:
		_, err = r.conn.Do("SADD", redis.Args{}.Add(key).AddFlat(member)...)
	default:
		_, err = r.conn.Do("SADD", key, member)
	}

	return
}

// SRem removes the specified members from the set stored at key.
// param 'member' support array of string, int, int32, int64, float32, float64, interface{}
func (r *Dao) SRem(key string, member interface{}) (err error) {
	switch member.(type) {
	case []string, []int, []int32, []int64, []float32, []float64, []interface{}:
		_, err = r.conn.Do("SREM", redis.Args{}.Add(key).AddFlat(member)...)
	default:
		_, err = r.conn.Do("SREM", key, member)
	}

	return
}

// SIsMember Returns if member is a member of the set stored at key.
func (r *Dao) SIsMember(key string, member interface{}) (bool, error) {
	exist, err := r.conn.Do("SISMEMBER", key, member)

	if exist.(int64) == 1 {
		return true, nil
	}
	return false, err
}

// SMembers get all members from the set called key
func (r *Dao) SMembers(key string) ([]string, error) {
	var result []string

	value, err := r.conn.Do("SMEMBERS", key)

	if err != nil {
		return nil, err
	}

	switch value := value.(type) {
	case []interface{}:
		result = make([]string, len(value))
		for i := 0; i < len(value); i++ {
			result[i] = string(value[i].([]uint8))
		}
	}

	return result, nil
}
