package concurrency

import (
	"fmt"
	"library_management/services"
	"time"
)

type ReservationWorker struct {
	Library  *services.Library
	WorkerID int
}

func NewReservationWorker(library *services.Library, workerID int) *ReservationWorker {
	return &ReservationWorker{
		Library:  library,
		WorkerID: workerID,
	}
}

func (rw *ReservationWorker) Run() {
	for request := range rw.Library.ReservationQueue {
		rw.processReservation(request)
	}
}

func (rw *ReservationWorker) processReservation(request services.ReservationRequest) {
	rw.Library.Mu.Lock()
	book, ok := rw.Library.Books[request.BookID]
	if !ok {
		rw.Library.Mu.Unlock()
		request.Result <- fmt.Errorf("Worker %d: book with ID %d not found", rw.WorkerID, request.BookID)
		return
	}

	if book.Status == "Reserved" || book.Status == "Borrowed" {
		rw.Library.Mu.Unlock()
		request.Result <- fmt.Errorf("Worker %d: book with ID %d is already reserved or borrowed", rw.WorkerID, request.BookID)
		return
	}

	book.Status = "Reserved"
	rw.Library.Books[request.BookID] = book
	rw.Library.Mu.Unlock()

	time.Sleep(1 * time.Second)

	fmt.Printf("Worker %d: Book %d reserved by member %d\n", rw.WorkerID, request.BookID, request.MemberID)

	time.AfterFunc(5*time.Second, func() {
		rw.Library.Mu.Lock()
		defer rw.Library.Mu.Unlock()
		book, ok := rw.Library.Books[request.BookID]
		if !ok {
			return
		}
		if book.Status == "Reserved" && request.Timestamp.Equal(time.Now().Add(-5*time.Second)) {
			book.Status = "Available"
			rw.Library.Books[request.BookID] = book
			fmt.Printf("Worker %d: Reservation for Book %d by member %d cancelled due to timeout\n", rw.WorkerID, request.BookID, request.MemberID)
		}
	})

	request.Result <- nil
}

func StartReservationWorkers(library *services.Library, numWorkers int) {
	for i := 0; i < numWorkers; i++ {
		worker := NewReservationWorker(library, i+1)
		go worker.Run()
	}
}
