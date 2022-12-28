package main

type Store interface {
  Get(key string) ([]byte, error)

  Put(key string, value []byte) error

  Delete(key string) error

  Close() error

  IsNotFoundError(err error) bool

  IsBadRequestError(err error) bool
}
