package cache

import (
	"reflect"
	"testing"
)

func TestNewCache(t *testing.T) {
	testCache := NewCache()
	if reflect.TypeOf(testCache) != reflect.TypeOf(&Cache{}) {
		t.Errorf("NewCache method does not return Cache instance true")
	}
}

func TestGetDB(t *testing.T) {
	testCache := NewCache()
	testDB := testCache.db
	t.Log(reflect.TypeOf(testDB).Key())
	if !reflect.DeepEqual(testDB, testCache.db) {
		t.Errorf("TestGetDB method does not DeepEqual to db of the cache")
	}
}

func TestGet(t *testing.T) {
	testCache := NewCache()
	testKey := "test_key"
	testValue := "test_value"
	testCache.Set(testKey, testValue)
	returnedValue, exist := testCache.Get(testKey)
	if !exist {
		t.Errorf("Get does not return existence correctly")
	}
	if returnedValue != testValue {
		t.Errorf("Get does not return the correct value")
	}
}

func TestSet(t *testing.T) {
	testCache := NewCache()
	testKey := "test_key"
	testValue := "test_value"
	testCache.Set(testKey, testValue)
	returnedValue, exist := testCache.Get(testKey)
	if !exist {
		t.Errorf("Set does not set")
	}
	if returnedValue != testValue {
		t.Errorf("Set does not set the correct value")
	}
	newTestValue := "test_value2"
	testCache.Set(testKey, newTestValue)

	returnedValue, _ = testCache.Get(testKey)
	if returnedValue != newTestValue {
		t.Errorf("Set does not override the existing key")
	}
}

func TestDelete(t *testing.T) {
	testCache := NewCache()
	testKey := "test_key"
	testValue := "test_value"
	testCache.Set(testKey, testValue)
	testCache.Delete(testKey)
	returnedValue, exist := testCache.Get(testKey)
	if exist {
		t.Errorf("Delete does not change the existence")
	}
	if returnedValue == testValue {
		t.Errorf("Delete does not delete")
	}

}

func TestFlush(t *testing.T) {
	testCache := NewCache()
	testKey := "test_key"
	testValue := "test_value"
	testCache.Set(testKey, testValue)
	testCache.Flush()
	if len(testCache.db) != 0 {
		t.Errorf("Flush does not remove all the stored data")
	}
}
