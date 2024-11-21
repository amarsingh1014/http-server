package main

// import (
// 	"log"
// 	"syscall"
// )

// const (
// 	maxEvents = 32
// )

// type Epoll struct {
// 	fd int
// 	events []syscall.EpollEvent
// }

// func createEpoll() (*Epoll, error) {
// 	fd, err := syscall.EpollCreate(0)
// 	if err != nil {
// 		log.Fatalf("Error creating epoll: %v", err)
// 	}

// 	return &Epoll{fd: fd,
// 		events: make([]syscall.EpollEvent, maxEvents),
// 	}, nil
// }

// func (e *Epoll) addFD(fd int) error {
// 	event := syscall.EpollEvent{
// 		Events: syscall.EPOLLIN ,
// 		Fd: int32(fd),
// 	}

// 	return syscall.EpollCtl(e.fd, syscall.EPOLL_CTL_ADD, fd, &event)
// }

// func (e *Epoll) deleteFD(fd int) error {
// 	return syscall.EpollCtl(e.fd, syscall.EPOLL_CTL_DEL, fd, nil)
// }

// func (e *Epoll) wait() ([]syscall.EpollEvent , error) {
// 	n, err := syscall.EpollWait(e.fd, e.events, -1)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return e.events[:n], nil
// }