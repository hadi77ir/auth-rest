package storage

import (
	"auth-rest/internal/app"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/storage/memory/v2"
)

type NewStorageFunc func(name string) fiber.Storage

type StorageKeyType int

const (
	StorageKey StorageKeyType = 0x11
)

func NewFromGlobals(globals *app.AppGlobals, name string) fiber.Storage {
	return app.Value[NewStorageFunc](globals, StorageKey)(name)
}

func Setup(globals *app.AppGlobals, storageFn NewStorageFunc) (NewStorageFunc, error) {
	if storageFn == nil {
		storageFn = func(name string) fiber.Storage {
			return memory.New()
		}
	}
	globals.Set(StorageKey, storageFn)
	return storageFn, nil
}
