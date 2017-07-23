package matekasse

import "testing"

func TestInsertUser(t *testing.T) {
	connectDB()
	createUser(0)
	closeDB()
}

func BenchmarkInsertUser(b *testing.B) {
	b.StopTimer()
	connectDB()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		createUser(ID(i))
	}
	b.StopTimer()
	closeDB()
}
