package escrow

import (
	"reflect"
)

type FreeCallUserStorage struct {
	delegate TypedAtomicStorage
}

func NewFreeCallUserStorage(atomicStorage AtomicStorage) *FreeCallUserStorage {
	return &FreeCallUserStorage{
		delegate: &TypedAtomicStorageImpl{
			atomicStorage: &PrefixedAtomicStorage{
				delegate:  atomicStorage,
				keyPrefix: "/free-call-user/storage",
			},
			keySerializer:     serialize,
			keyDeserializer:   deserialize,
			keyType:           reflect.TypeOf(FreeCallUserKey{}),
			valueSerializer:   serialize,
			valueDeserializer: deserialize,
			valueType:         reflect.TypeOf(FreeCallUserData{}),
		},
	}
}

func (storage *FreeCallUserStorage) Get(key *FreeCallUserKey) (state *FreeCallUserData, ok bool, err error) {
	value, ok, err := storage.delegate.Get(key)
	if err != nil || !ok {
		return nil, ok, err
	}
	return value.(*FreeCallUserData), ok, err
}

func (storage *FreeCallUserStorage) GetAll() (states []*FreeCallUserData, err error) {
	values, err := storage.delegate.GetAll()
	if err != nil {
		return
	}

	return values.([]*FreeCallUserData), nil
}

func (storage *FreeCallUserStorage) Put(key *FreeCallUserKey, state *FreeCallUserData) (err error) {
	return storage.delegate.Put(key, state)
}

func (storage *FreeCallUserStorage) PutIfAbsent(key *FreeCallUserKey, state *FreeCallUserData) (ok bool, err error) {
	return storage.delegate.PutIfAbsent(key, state)
}

func (storage *FreeCallUserStorage) CompareAndSwap(key *FreeCallUserKey, prevState *FreeCallUserData, newState *FreeCallUserData) (ok bool, err error) {
	return storage.delegate.CompareAndSwap(key, prevState, newState)
}
