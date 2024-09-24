// File:		session.go
// Created by:	Hoven
// Created on:	2024-09-24
//
// This file is part of the Example Project.
//
// (c) 2024 Example Corp. All rights reserved.

package pgin

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-contrib/sessions/memstore"
	redisStore "github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
	"github.com/go-puzzles/puzzles/plog"
	"github.com/gomodule/redigo/redis"
)

type StoreOption interface {
	GetStore() sessions.Store
}

type CookieStore struct {
	keyPairs [][]byte
}

func InitCookieStore(keyPairs ...[]byte) *CookieStore {
	return &CookieStore{keyPairs: keyPairs}
}

func (c *CookieStore) GetStore() sessions.Store {
	return cookie.NewStore(c.keyPairs...)
}

type RedisStore struct {
	pool     *redis.Pool
	keyPairs [][]byte
}

func InitRedisStore(pool *redis.Pool, keyPairs ...[]byte) *RedisStore {
	return &RedisStore{pool: pool, keyPairs: keyPairs}
}

func (r *RedisStore) GetStore() sessions.Store {
	store, err := redisStore.NewStoreWithPool(r.pool, r.keyPairs...)
	if err != nil {
		plog.Fatalf("new redis store failed: %v", err)
	}
	return store
}

type MemoryStore struct {
	keyPairs [][]byte
}

func InitMemStore(keyPairs ...[]byte) *MemoryStore {
	return &MemoryStore{keyPairs: keyPairs}
}

func (m *MemoryStore) GetStore() sessions.Store {
	return memstore.NewStore(m.keyPairs...)
}

func NewSession(key string, opt StoreOption) gin.HandlerFunc {
	store := opt.GetStore()
	return sessions.Sessions(key, store)
}

func GetSession(c *gin.Context) sessions.Session {
	return sessions.Default(c)
}

func GetSessionByKey(c *gin.Context, key string) sessions.Session {
	return sessions.DefaultMany(c, key)
}
