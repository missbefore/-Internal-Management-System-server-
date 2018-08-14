package main

import (
	"sync"
	"io"
	"log"
	"os"
	"encoding/json"
	"errors"
	"net/rpc"
)

var saveQueueLength = 1000


type URLStore struct {
	urls  map[string]string
	mu    sync.RWMutex
	save  chan record
}

type  ProxyStore struct {
	urls *URLStore

	client *rpc.Client
}

type Store interface {
	Put(url, key *string) error
	Get(key, url *string) error
}


type record struct {
	Key, URL string
}

func NewURLStore(filename string) *URLStore {
	s := &URLStore{urls: make(map[string]string)}

	if filename != "" {
		s.save = make(chan record, saveQueueLength)

		if err := s.load(filename); err != nil {

			log.Println("Error loading URLStore: ", err)

		}

		go s.saveLoop(filename)
	}

	return s
}

func (s *URLStore) saveLoop(filename string) {
	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_APPEND ,0664)

	if err != nil {
		log.Println("URLStore:", err)
	}

	defer f.Close()

	e := json.NewEncoder(f)

	for {
		r := <-s.save

		if err := e.Encode(r); err != nil {
			log.Println("Error saving to URLstore:", err)
		}
	}
}

func (s *URLStore) Get(key, url *string) error {
	s.mu.RLock()
	defer s.mu.RUnlock()

	if u, ok := s.urls[*key]; ok{

		*url = u

		return nil
	}

	return errors.New("key not found")
}

func (s *URLStore) Set(key, url *string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, present := s.urls[*key]; present {

		return errors.New("key already exists")
	}
	s.urls[*key] = *url
	return nil
}

func (s *URLStore) Count() int {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return len(s.urls)
}


func (s *URLStore) load(filename string) error {

	f, err := os.Open(filename)

	if err != nil {
		log.Println("Error opening URLStore:", err)
		return err
	}

	defer f.Close()

	d := json.NewDecoder(f)

	if err == nil {
		var r record

		if err = d.Decode(&r); err == nil {
			s.Set(&r.Key, &r.URL)
		}
	}

	if err == io.EOF {
		return nil
	}

	log.Println("Error decoding URLStore:", err)
	return err
}

func (s *URLStore) Put(url, key *string) error {
	for {
		*key = genKey(s.Count())

		if err := s.Set(key, url); err == nil {
			break
		}
	}

	if s.save != nil {

		s.save <- record{*key, *url}
	}

	return nil
}

func NewProxyStore(addr string) *ProxyStore {
	client, err := rpc.DialHTTP("tcp", addr)

	if err != nil {

		log.Println("Error constructing ProxyStore: ", err)
	}

	return &ProxyStore{urls: NewURLStore(""), client:client}
}

func (s *ProxyStore) Get(key, url *string) error {
	if err := s.urls.Get(key, url); err == nil {
		return nil
	}

	if err := s.client.Call("Store.Get", key, url); err != nil {
		return err
	}

	s.urls.Set(key, url)

	return nil
}

func (s *ProxyStore) Put(url, key *string) error {

	if err := s.client.Call("Store.Post", url, key);err != nil {
		return err
	}

	s.urls.Set(key, url)

	return nil
}
